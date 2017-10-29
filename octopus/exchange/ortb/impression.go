package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/random"
)

type Impression struct {
	inner  *openrtb.Impression
	banner exchange.Banner
	video  exchange.Video
	native exchange.Native
	cid    string
}

func (m *Impression) CID() string {
	if m.cid == "" {
		m.cid = <-random.ID
	}
	return m.cid
}

func (m *Impression) ID() string {
	return m.inner.ID
}

func (m *Impression) BidFloor() float64 {
	return m.inner.BidFloor
}

func (m *Impression) Banner() exchange.Banner {
	if m.Type() != exchange.AdTypeBanner {
		return nil
	}
	if m.banner == nil {
		m.banner = &banner{inner: m.inner.Banner}
	}
	return m.banner
}

func (m *Impression) Video() exchange.Video {
	panic("implement video")
}

func (m *Impression) Native() exchange.Native {
	panic("implement native")
}

func (m *Impression) Attributes() map[string]interface{} {
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

func (m *Impression) Type() exchange.ImpressionType {
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

func (m *Impression) Secure() bool {
	if m.inner.Secure == 1 {
		return true
	}
	return false
}
