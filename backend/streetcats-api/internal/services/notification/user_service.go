package notification

import (
	"github.com/golang-jwt/jwt/v5"
)

// GetUserIdByPreferredUsername - Retrieve id using username provided
func (s *Service) GetUserIdByPreferredUsername(claims any) (uint64, error) {
	var (
		claimsMap jwt.MapClaims
		username  string
		ok        bool
	)

	if claimsMap, ok = claims.(jwt.MapClaims); ok {
		username = claimsMap["preferred_username"].(string)
	}

	return s.notificationRepo.GetUserIdByPreferredUsername(username)
}
