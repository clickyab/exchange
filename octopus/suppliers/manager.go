package suppliers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/suppliers/internal/models"
	"clickyab.com/exchange/octopus/suppliers/internal/renderer"
	"clickyab.com/exchange/octopus/suppliers/internal/restful"
	"github.com/clickyab/services/mysql"

	"github.com/sirupsen/logrus"
)

var (
	sm *supplierManager
)

type supplierManager struct {
	suppliers map[string]models.Supplier
	lock      *sync.RWMutex
}

func restRendererFactory(sup exchange.Supplier, in string) exchange.Renderer {
	switch in {
	case "rest":
		// TODO : /api is hardcoded
		// TODO : restRenderFactory should have no arg
		return renderer.NewRenderer()
	default:
		logrus.Panicf("supplier with key %s not found", in)
	}
	return nil
}

func (sm *supplierManager) loadSuppliers() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	m := models.NewManager()
	sm.suppliers = m.GetSuppliers(restRendererFactory)
}

func (sm *supplierManager) Initialize() {
	sm.loadSuppliers()
	reloadChan := make(chan os.Signal)
	signal.Notify(reloadChan, syscall.SIGHUP)
	go func() {
		for i := range reloadChan {
			logrus.Infof("Reloading supplier config, due to signal %s", i)
			sm.loadSuppliers()
		}
	}()
}

// GetSupplierByKey return a single supplier by its id
func GetSupplierByKey(key string) (exchange.Supplier, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	if s, ok := sm.suppliers[key]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("supplier with key %s not found", key)
}

// GetSupplierByName return a single supplier by its name
func GetSupplierByName(name string) (exchange.Supplier, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	for i := range sm.suppliers {
		if sm.suppliers[i].Name() == name {
			return sm.suppliers[i], nil
		}
	}

	return nil, fmt.Errorf("supplier with name %s not found", name)
}

// GetBidRequest try to get an bid request from a http request
func GetBidRequest(key string, r *http.Request) (exchange.BidRequest, error) {
	sup, err := GetSupplierByKey(key)
	if err != nil {
		return nil, err
	}

	// Make sure the profit margin is added to the request
	switch sup.Type() {
	case "rest":
		return restful.GetBidRequest(sup, r)
	default:
		logrus.Panicf("Not a supported type: %s", sup.Type())
		return nil, fmt.Errorf("not supported type: %s", sup.Type())
	}
}

func RenderBidRequestRtbToRest(bq exchange.BidRequest) restful.BidRequest {
	imp := []restful.Imp{}
	for i := range bq.Imp() {
		imp = append(imp, restful.Imp{
			IID:   bq.Imp()[i].ID(),
			IType: bq.Imp()[i].Type(),
			IBanner: &restful.Banner{
				IID:   bq.Imp()[i].Banner().ID(),
				IW:    bq.Imp()[i].Banner().Width(),
				IH:    bq.Imp()[i].Banner().Height(),
				IAttr: bq.Imp()[i].Attributes(),
			},
		})
	}
	res := restful.BidRequest{
		IImp: imp,
		IID:  bq.ID(),
		IDevice: restful.Device{
			ICID:        bq.Device().CID(),
			ILAC:        bq.Device().LAC(),
			IMNC:        bq.Device().MNC(),
			IMCC:        bq.Device().MCC(),
			ILang:       bq.Device().Language(),
			IConnType:   bq.Device().ConnType(),
			ICarrier:    bq.Device().Carrier(),
			IOs:         bq.Device().OS(),
			IModel:      bq.Device().Model(),
			IMake:       bq.Device().Make(),
			IDeviceType: bq.Device().DeviceType(),
			IIP:         bq.Device().IP(),
			IUA:         bq.Device().UserAgent(),
			IGeo: restful.Geo{
				ILatLon: exchange.LatLon{
					Lon:   bq.Device().Geo().LatLon().Lon,
					Lat:   bq.Device().Geo().LatLon().Lat,
					Valid: bq.Device().Geo().LatLon().Valid,
				},
				ICountry: exchange.Country{
					Valid: bq.Device().Geo().Country().Valid,
					Name:  bq.Device().Geo().Country().Name,
					ISO:   bq.Device().Geo().Country().ISO,
				},
				IRegion: exchange.Region{
					Valid: bq.Device().Geo().Region().Valid,
					Name:  bq.Device().Geo().Region().Name,
					ISO:   bq.Device().Geo().Region().ISO,
				},
				IIsp: exchange.ISP{
					Valid: bq.Device().Geo().ISP().Valid,
					Name:  bq.Device().Geo().ISP().Name,
				},
			},
		},
		IUser: restful.User{
			IID: bq.User().ID(),
		},
		ITMax:  bq.TMax(),
		ITest:  bq.Test(),
		IAttr:  bq.Attributes(),
		IBCat:  bq.BlockedCategories(),
		IWLang: bq.AllowedLanguage(),
		IBAdv:  bq.BlockedAdvertiserDomain(),
	}
	if s, ok := bq.Inventory().Publisher().(*restful.Site); ok {
		res.ISite = s
	} else if a, ok := bq.Inventory().Publisher().(*restful.App); ok {
		res.IApp = a
	} else {
		panic("[BUG] wrong bid request inventory")
	}
	return res
}

func init() {
	sm = &supplierManager{lock: &sync.RWMutex{}}
	mysql.Register(sm)
}
