package auth

import (
	"fmt"
	"lab/exp1/src/utils"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TimeInterval struct {
	StartTime       string // format YYYY-MM-DD HH:MM:SS
	ExpireTime      string // format YYYY-MM-DD HH:MM:SS
	Repeat          bool
	RepeatFrequency string // daily, weekly, monthly
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
		return "", fmt.Errorf("it is not a bearer token")
	} else {
		if token := tkArray[1]; strings.Trim(token, "") == "" {
			return "", fmt.Errorf("token is not present")
		} else {
			return token, nil
		}
	}
}

func enforcePolicies(token jwt.MapClaims, policy Policy, opts PolicyOpts) (bool, error) {

	keycloak_client_id := os.Getenv("KEYCLOAK_CLIENT_ID")
	userID := token["sub"].(string)
	userRoles := []string{}
	var err error

	// GEt isers roles via client id
	if resource_access, exist := token["resource_access"].(map[string]interface{}); exist {
		if client, exist := resource_access[keycloak_client_id].(map[string]interface{}); exist {
			if roles, exist := client["roles"].([]interface{}); exist {
				userRoles = utils.ConvertSlice(roles)
			}
		}
	}

	rolePolicyResult := applyRolePolicy(userRoles, policy.Roles)
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
		decisionStrategy = "unanymous"
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
	if strings.Trim(times.StartTime, "") == "" && strings.Trim(times.ExpireTime, "") == "" {
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

		if time.Now().Before(expireTime) {
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
	if strings.Trim(times.StartTime, "") != "" && strings.Trim(times.ExpireTime, "") != "" {

		startTime, err1 := time.Parse(timeLayout, times.StartTime)
		expireTime, err2 := time.Parse(timeLayout, times.ExpireTime)

		nowTime := time.Now()
		var err error

		if err1 != nil || err2 != nil {
			fmt.Println(err1, err2)
			return false
		}

		if times.Repeat {

			repeatFrequency := strings.ToLower(times.RepeatFrequency)

			if repeatFrequency != "daily" && repeatFrequency != "weekly" && repeatFrequency != "monthly" {
				repeatFrequency = "daily"
			}

			if repeatFrequency == "daily" {

				hour, min, sec := nowTime.Clock()
				year, month, day := startTime.Date()
				nowTime, err = time.Parse(timeLayout, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec))

				if err != nil {
					fmt.Println(err)
					return false
				}

			} else if repeatFrequency == "weekly" {

			} else if repeatFrequency == "monthly" {

				hour, min, sec := nowTime.Clock()
				_, _, day := nowTime.Date()
				year, month, _ := startTime.Date()

				nowTime, err = time.Parse(timeLayout, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec))

				if err != nil {
					fmt.Println(err)
					return false
				}

			}

		}

		if nowTime.After(startTime) && nowTime.Before(expireTime) {
			return true
		}
	}

	return false
}

func applyRolePolicy(userRoles []string, permittedRoles []string) bool {

	if len(permittedRoles) == 0 {
		return true
	}

	for _, permitted := range permittedRoles {
		for _, role := range userRoles {
			if role == permitted {
				return true
			}
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
