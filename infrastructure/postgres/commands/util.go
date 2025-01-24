package commands

const (
	AccountTable                 = "accounts"
	AccountIdField               = "id"
	AccountEmailField            = "email"
	AccountPasswordField         = "password"
	AccountRegistrationDateField = "registration_date"
	AccountIsVerifiedField       = "is_verified"
	AccountRoleField             = "role_"
	AccountUserIdField           = "user_id"
)

const (
	SessionTable             = "sessions"
	SessionIdField           = "id"
	SessionAccessTokenField  = "access_token"
	SessionRefreshTokenField = "refresh_token"
	SessionAccountIdField    = "account_id"
	SessionExpiresAtField    = "expires_at"
)
