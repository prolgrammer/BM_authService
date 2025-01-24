package entities

import "time"

const (
	AccountIdClaimName = "sub"
	ExpiresAtClaimName = "exp"
	RoleClaimName      = "role"
)

type AccessTokenClaims map[string]any

func NewClaims(accountId, role string, expiresAt time.Time) AccessTokenClaims {
	return AccessTokenClaims{
		AccountIdClaimName: accountId,
		ExpiresAtClaimName: expiresAt.Unix(),
		RoleClaimName:      role,
	}
}
