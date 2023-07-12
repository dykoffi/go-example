package auth

// TODO: Define all policies you want to apply on a route

var (
	IsAdminUser                = Policy{Roles: []string{"Admin"}}
	IsAuthenticatedUserAtNight = Policy{Times: TimeInterval{StartTime: "2023-07-12 07:00:00", ExpireTime: "2023-08-13 23:59:00", Repeat: true, RepeatFrequency: "hourly"}}
)
