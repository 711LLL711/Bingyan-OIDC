package model

type User struct {
	Username string `gorm:"column:username"`
	UserID   string `gorm:"primaryKey;column:userid"`
	Password string `gorm:"column:password"`
}
