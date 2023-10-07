package model

type User struct {
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	UserID   string `gorm:"primaryKey;column:userid"`
	Password string `gorm:"column:password"`
}
