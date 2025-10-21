package gdmodel

import "github.com/thinhphamnls/gd/entity/gd"

type QuickBooksAuthQuery struct {
	AppUserName string
	AppTenant   string
	Deleted     gdentity.QuickbooksOAuthDeletedType
}

type QuickBooksAuthView struct {
	QuickbooksOathId       uint   `gorm:"column:quickbooks_oauth_id"`
	OauthAccessToken       string `gorm:"column:oauth_access_token"`
	OauthAccessTokenSecret string `gorm:"column:oauth_access_token_secret"`
	QbRealm                string `gorm:"column:qb_realm"`
}
