package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TimeInterval struct {
	StartTime  string // format YYYY-MM-DD HH:MM:SS
	ExpireTime string // format YYYY-MM-DD HH:MM:SS
	Repeat     bool
}

type Policy struct {
	Users []string     // users granted by this policy
	Roles []string     // roles granted by this policy
	Times TimeInterval // times for applying this policy
}

type PolicyOpts struct {
	Logic            bool   // True or false
	DecisionStrategy string // Unanymous ou Affirmative
}

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

func enforcePolicies(token jwt.MapClaims, policy Policy, opts PolicyOpts) (bool, error) {

	userID := fmt.Sprintf("%v", token["sub"])
	var err error

	roles := strings.Join(policy.Roles, "")

	rolePolicyResult := applyRolePolicy([]string{}, roles)
	userPolicyResult := applyUserPolicy(userID, policy.Users)
	timePolicyResult := applyTimePolicy(policy.Times)

	if !rolePolicyResult {
		err = fmt.Errorf(ErrorPermissionRole)
	} else if !userPolicyResult {
		err = fmt.Errorf(ErrorPermissionUser)
	} else if !timePolicyResult {
		err = fmt.Errorf(ErrorPermissionTime)
	}

	decisionStrategy := strings.ToLower(opts.DecisionStrategy)

	if decisionStrategy != "affirmative" && decisionStrategy != "unanymous" {
		opts.DecisionStrategy = "unanymous"
	}

	logic := opts.Logic
	result := false

	switch decisionStrategy {
	case "affirmative":
		result = (rolePolicyResult || userPolicyResult || timePolicyResult)
		if logic {
			return result, err
		}
		return !result, err

	case "unanymous":
		result := (rolePolicyResult && userPolicyResult && timePolicyResult)
		if logic {
			return result, err
		}
		return !result, err
	}

	return result, err
}

func applyTimePolicy(times TimeInterval) bool {

	timeLayout := "2006-01-02 15:04:05"

	// If any of the times is defined, we consider this policy like not applicable, so return true
	if strings.Trim(times.StartTime, "") == "" && strings.Trim(times.StartTime, "") == "" {
		return true
	}

	// If only the expire time is defined
	if strings.Trim(times.StartTime, "") == "" && strings.Trim(times.ExpireTime, "") != "" {

		// Verify just if now time is before expire time
		expireTime, err := time.Parse(timeLayout, times.ExpireTime)

		if err != nil {
			fmt.Println(err)
			return false
		}

		if time.Now().After(expireTime) {
			return true
		}
	}

	// If only the start time is specified
	if strings.Trim(times.StartTime, "") != "" && strings.Trim(times.ExpireTime, "") == "" {

		// Verify just if now time is after the start time
		startTime, err := time.Parse(timeLayout, times.StartTime)

		if err != nil {
			fmt.Println(err)
			return false
		}

		if time.Now().After(startTime) {
			return true
		}
	}

	// If the both startTime and endTime are defined
	if strings.Trim(times.StartTime, "") != "" && strings.Trim(times.StartTime, "") != "" {

		startTime, err1 := time.Parse(timeLayout, times.StartTime)
		expireTime, err2 := time.Parse(timeLayout, times.ExpireTime)

		if err1 != nil || err2 != nil {
			fmt.Println(err1, err2)
			return false
		}

		if time.Now().After(startTime) && time.Now().Before(expireTime) {
			return true
		}
	}

	return false
}

func applyRolePolicy(userRole []string, permittedRoles string) bool {

	if len(permittedRoles) == 0 {
		return true
	}

	for _, role := range userRole {
		if strings.Contains(permittedRoles, role) {
			return true
		}
	}

	return false
}

func applyUserPolicy(userId string, permittedIds []string) bool {

	if len(permittedIds) == 0 {
		return true
	}

	for _, permitted := range permittedIds {
		if userId == permitted {
			return true
		}
	}

	return false

}

func Protect(policy Policy, opts PolicyOpts) func(*fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		// Get user token and verify if it is in correct format
		userToken, err := GetBearerToken(ctx.Get("Authorization"))
		if err != nil {
			fmt.Println(err.Error())
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": MessageExpiredOrInvalidToken, "error": ErrorExpiredOrInvalidToken})
		}

		// Once the user token got, try to check it validity and get user information
		_, tokenClaims, err := Kc.Gocloak.DecodeAccessToken(ctx.Context(), userToken, Kc.Realm)
		if err != nil {
			fmt.Println(err.Error())
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": MessageExpiredOrInvalidToken, "error": ErrorExpiredOrInvalidToken})
		}

		// Check if user is authorized by enforcing policy rules required for accessing resource
		userIsAuthorized, err := enforcePolicies(*tokenClaims, policy, opts)
		if !userIsAuthorized {
			switch err.Error() {
			case ErrorPermissionRole:
				return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]string{"message": MessageUserRequiredRole, "error": ErrorPermissionRole})

			case ErrorPermissionUser:
				return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]string{"message": MessageUnAuthorizedUser, "error": ErrorPermissionUser})

			case ErrorPermissionTime:
				return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]string{"message": MessageNotCorrespondantTime, "error": ErrorPermissionTime})
			}
		}

		// Go to the next function
		return ctx.Next()
	}
}
