package models

type (
	Users struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	DisplayUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
