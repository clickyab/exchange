package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type impression struct {
	id         string
	bidFloor   float64
	banner     exchange.Banner
	video      exchange.Video
	native     exchange.Native
	fType      exchange.ImpressionType
	secure     bool
	attributes map[string]interface{}
}

func (i impression) ID() string {
	return i.id
}

func (i impression) BidFloor() float64 {
	return i.bidFloor
}

func (i impression) Banner() exchange.Banner {
	return i.banner
}

func (i impression) Video() exchange.Video {
	return i.video
}

func (i impression) Native() exchange.Native {
	return i.native
}

func (i impression) Attributes() map[string]interface{} {
	return i.attributes
}

func (i impression) Type() exchange.ImpressionType {
	return i.fType
}

func (i impression) Secure() bool {
	return i.secure
}

func impressionAttributes(m openrtb.Impression) map[string]interface{} {
	return map[string]interface{}{
		"Audio":             m.Audio,
		"BidFloorCurrency":  m.BidFloorCurrency,
		"DisplayManager":    m.DisplayManager,
		"DisplayManagerVer": m.DisplayManagerVer,
		"Exp":               m.Exp,
		"Ext":               m.Ext,
		"IFrameBuster":      m.IFrameBuster,
		"Instl":             m.Instl,
		"Pmp":               m.Pmp,
		"TagID":             m.TagID,
	}
}

func newImpressions(m []openrtb.Impression) []exchange.Impression {
	ms := make([]exchange.Impression, 0)
	for _, v := range m {
		ms = append(ms, newImpression(v))
	}
	return ms
}

func newImpression(m openrtb.Impression) exchange.Impression {
	return impression{
		attributes: impressionAttributes(m),
		id:         m.ID,
		fType:      impressionType(m),
		secure: func() bool {
			if m.Secure == 1 {
				return true
			}
			return false
		}(),
		bidFloor: m.BidFloor,
		banner:   newBanner(m.Banner),
	}
}

func impressionType(m openrtb.Impression) exchange.ImpressionType {
	if m.Banner != nil {
		return exchange.AdTypeBanner
	}
	if m.Video != nil {
		return exchange.AdTypeVideo
	}
	if m.Native != nil {
		return exchange.AdTypeNative
	}
	panic("not a valid impression type")
}
