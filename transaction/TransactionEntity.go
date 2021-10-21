package transaction

import (
	"golang-practice/campaign"
	"golang-practice/user"
	"time"
)

//representasi dari table transaction
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	User       user.User
	Campaign   campaign.Campaign
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
