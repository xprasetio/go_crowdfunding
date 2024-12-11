package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}