package middleware

import "golang.org/x/crypto/bcrypt"

// 密碼加密 HashPassword
func Encrypt(password string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// 密碼比對 (傳入未加密的密碼即可)
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
