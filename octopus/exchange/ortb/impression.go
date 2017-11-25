package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/random"
)

// impression ortb imp
type impression struct {
	inner  *openrtb.Impression
	banner exchange.Banner
	sup    exchange.Supplier
	cid    string
}

func (m *impression) Currency() string {
	return m.sup.Currency()
}

// CID return ortb CID
func (m *impression) CID() string {
	if m.cid == "" {
		m.cid = <-random.ID
	}
	return m.cid
}

// ID return ortb ID
func (m *impression) ID() string {
	return m.inner.ID
}

// BidFloor return ortb BidFloor
func (m *impression) BidFloor() float64 {
	return m.inner.BidFloor
}

// Banner return ortb Banner
func (m *impression) Banner() exchange.Banner {
	if m.Type() != exchange.AdTypeBanner {
		return nil
	}
	if m.banner == nil {
		m.banner = &banner{inner: m.inner.Banner}
	}
	return m.banner
}

// Video return ortb Video
func (m *impression) Video() exchange.Video {
	panic("implement video")
}

// Native return ortb Native
func (m *impression) Native() exchange.Native {
	panic("implement native")
}

// Attributes return ortb Attributes
func (m *impression) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Audio":             m.inner.Audio,
		"BidFloorCurrency":  m.inner.BidFloorCurrency,
		"DisplayManager":    m.inner.DisplayManager,
		"DisplayManagerVer": m.inner.DisplayManagerVer,
		"Exp":               m.inner.Exp,
		"Ext":               m.inner.Ext,
		"IFrameBuster":      m.inner.IFrameBuster,
		"Instl":             m.inner.Instl,
		"Pmp":               m.inner.Pmp,
		"TagID":             m.inner.TagID,
	}
}

// Type return ortb Type
func (m *impression) Type() exchange.ImpressionType {
	if m.inner.Banner != nil {
		return exchange.AdTypeBanner
	}
	if m.inner.Video != nil {
		return exchange.AdTypeVideo
	}
	if m.inner.Native != nil {
		return exchange.AdTypeNative
	}
	panic("not valid ad type")
}

// Secure return ortb Secure
func (m *impression) Secure() bool {
	if m.inner.Secure == 1 {
		return true
	}
	return false
}
