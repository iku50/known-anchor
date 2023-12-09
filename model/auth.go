package model

type AuthLoginPostReq struct {
	// email
	Email string `json:"email"`
	// password
	Password string `json:"password"`
}

type AuthLoginPostResp struct {
	// token
	Token string `json:"token"`
}

type AuthRegisterPostReq struct {
	// username
	Username string `json:"username"`
	// email
	Email string `json:"email"`
	// password
	Password string `json:"password"`
}

type AuthRegisterPostResp struct {
}

type AuthConfirmPostReq struct {
	Email        string `json:"email"`
	ConfirmToken string `json:"confirm_token"`
}

type AuthConfirmPostResp struct {
}

type AuthActivatePostReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthActivatePostResp struct {
}
