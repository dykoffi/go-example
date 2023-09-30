package user

type Creds struct {
	Mail     string `json:"name" validate:"required,min=3,max=10"`
	Password string `json:"password" validate:"required,min=3,max10"`
}

type SuccessRes struct {
	AccessToken      string `json:"access_token"`
	IdToken          string `json:"id_token"`
	ExpiresIn        int32  `json:"expires_in"`
	RefreshExpiresIn int32  `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int32  `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type ErrorRes struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
