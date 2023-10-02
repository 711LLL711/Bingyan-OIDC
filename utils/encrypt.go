package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 密码加密: pwdHash
func PasswordHash(pwd *string) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(*pwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*pwd = string(hashedPwd)
	return nil
}

// 密码验证: pwdVerify
func PasswordVerify(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
