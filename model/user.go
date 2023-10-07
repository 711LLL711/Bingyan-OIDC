package model

type User struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Password string
	Email    string
	Username string
	Avatar   string
	Bio      string
	//用于邮箱激活
	VerificationToken string `gorm:"varchar(255)"`
	Verified          bool
}
