package dtos

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
}

type UpdateLangRequest struct {
	Lang string `json:"lang" binding:"required"`
}

type InitPasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required"`
}

/************** Response **************/

type LoginResponse struct {
	User      UserResponse `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

type UserResponse struct {
	Username string `json:"username"`
	Lang     string `json:"lang"`
}

type InitInfoResponse struct {
	IsInit bool `json:"isInit"`
}

type TokenDetail struct {
	AccessId     string `json:"access_id"`
	RefreshId    string `json:"refresh_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}
