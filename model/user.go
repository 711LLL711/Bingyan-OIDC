package model

import (
	"time"
)

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Password string
	Email    string
	Username string
	Avatar   string
	Bio      string
	//用于邮箱激活
	ActivationToken string `gorm:"varchar(255)"`
	Activated       int    `gorm:"type:int(10)"`
	EmailVerifiedAt *time.Time
}
