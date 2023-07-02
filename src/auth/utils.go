package auth

import (
	"fmt"
	"strings"
)

func GetBearerToken(authorization string) (string, error) {

	tkArray := strings.Split(authorization, " ")

	if strings.ToLower(tkArray[0]) != "bearer" {
		return "", fmt.Errorf("It is not a Bearer token")
	} else {
		if token := tkArray[1]; strings.Trim(token, "") == "" {
			return "", fmt.Errorf("Token is not present")
		} else {
			return token, nil
		}
	}

}

func GetTokenClaims(token string) (string, error) {

	return "", nil
}

func GetUserRoles(user string) {}
func GetUserId(user string)    {}
