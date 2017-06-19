package models

import (
	"clickyab.com/exchange/crane/models/internal/ad"
	"clickyab.com/exchange/crane/models/internal/campaign"
	"clickyab.com/exchange/crane/models/internal/publisher"
	"clickyab.com/exchange/crane/models/internal/user"
)

// NewAdManager returns publisher manager
func NewAdManager() *ad.Manager {
	return ad.NewAdManager()
}

// NewCampaignManager returns publisher manager
func NewCampaignManager() *campaign.Manager {
	return campaign.NewCampaignManager()
}

// NewUserManager returns publisher manager
func NewUserManager() *user.Manager {
	return user.NewUserManager()
}

// NewPublisherManager returns publisher manager
func NewPublisherManager() *publisher.Manager {
	return publisher.NewPublisherManager()
}
