package auth

const (
	MessageUserRequiredRole      = "This user has not the required role to access to this resource"
	MessageNotCorrespondantTime  = "This user can not access to this resource at this time"
	MessageUnAuthorizedUser      = "This user can not access to this resource"
	MessageExpiredOrInvalidToken = "The user token is expired or not valid"
)

const (
	ErrorPermissionUser        = "ErrorPermissionUSer"
	ErrorPermissionRole        = "ErrorPermissionRole"
	ErrorPermissionTime        = "ErrorPermissionTime"
	ErrorExpiredOrInvalidToken = "ExpiredOrInvalidToken"
)
