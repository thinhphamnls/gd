package gdentity

type QuickbooksOAuthDeletedType int32

const (
	QuickbooksOAuthNotDelete QuickbooksOAuthDeletedType = 0
	QuickbooksOAuthDeleted   QuickbooksOAuthDeletedType = 1
)

type QuickbooksOAuthType int32

const (
	QuickbooksOAuthOAuth1 QuickbooksOAuthType = 1
	QuickbooksOAuthOAuth2 QuickbooksOAuthType = 2
)
