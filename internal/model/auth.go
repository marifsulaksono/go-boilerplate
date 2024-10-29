package model

type (
	Login struct {
		GrantType string `json:"grant_type"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string                 `json:"access_token"`
		RefreshToken string                 `json:"refresh_token"`
		Metadata     map[string]interface{} `json:"metadata"`
	}
)

type (
	TokenAuth struct {
		RefreshToken string `json:"refresh_token" gorm:"not null"`
		UserID       string `json:"user_id" gorm:"not null"`
		IP           string `json:"ip" gorm:"not null;char(128)"`
	}
)
