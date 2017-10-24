package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type impression struct {
	inner  *openrtb.Impression
	banner exchange.Banner
	video  exchange.Video
	native exchange.Native
}

func (m *impression) ID() string {
	return m.inner.ID
}

func (m *impression) BidFloor() float64 {
	return m.inner.BidFloor
}

func (m *impression) Banner() exchange.Banner {
	if m.Type() != exchange.AdTypeBanner {
		return nil
	}
	if m.banner == nil {
		m.banner = &banner{inner: m.inner.Banner}
	}
	return m.banner
}

func (m *impression) Video() exchange.Video {
	panic("implement video")
}

func (m *impression) Native() exchange.Native {
	panic("implement native")
}

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

func (m *impression) Secure() bool {
	if m.inner.Secure == 1 {
		return true
	}
	return false
}
