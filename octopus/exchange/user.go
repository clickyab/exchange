package exchange

// User is the identity of the user in system
// Design note:
// ID/BuyerID/BuyerUID are one thing in any level. so decide based on implementation and version
// XXX : YOB/Gender/Keywords/CustomData/Data are not supported
// Geo is used from device.
type User interface {
	// Unique consumer ID of this user on the exchange
	// Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	// Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	ID() string
}
