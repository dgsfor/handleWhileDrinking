package util

import "golang.org/x/crypto/bcrypt"

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
	}
	return string(hash)
}

// 验证密码
/*
	hashedPwd 保存在数据库的加密后密码
	plainPwd  前端传入byte后的密码
*/
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
