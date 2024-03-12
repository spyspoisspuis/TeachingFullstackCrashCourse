package entities

type (
	Users struct {
		Id       string `gorm:"primaryKey"`
		Username string
		Password string
	}
)