package middleware

import "golang.org/x/crypto/bcrypt"

// 密碼加密
func Encrypt(source string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// 密碼比對 (傳入未加密的密碼即可)
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
