package middleware

import "golang.org/x/crypto/bcrypt"

// 密碼加密 HashPassword
func Encrypt(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), err
}

// 密碼比對 (傳入未加密的密碼即可)
func Compare(password, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	return err == nil
}
