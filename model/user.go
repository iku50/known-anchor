package model

type UserUpdateReq struct {
	// username
	Username string `json:"username"`
	// email
	Email string `json:"email"`
	// password
	Password string `json:"password"`
}

type UserUpdateResp struct {
}
