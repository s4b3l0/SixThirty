package domain

import (
	"crud.com/packages/configuration"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CredentialDTO struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	UserPassword string `json:"user_password"`
}

type Credential struct {
	Username     string `json:"username" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"password_hash"`
}

func SetPassword(cred *Credential, pw string) {
	var hashPassword, _ = bcrypt.GenerateFromPassword([]byte(pw), 14)
	cred.PasswordHash = string(hashPassword)
}

func IsPasswordDiff(currentPassword []byte, loginPassword []byte) error {
	return bcrypt.CompareHashAndPassword(currentPassword, loginPassword)
}

func Update(cred *Credential) *gorm.DB {
	return configuration.Pdatabase.Updates(&cred)
}

func Save(cred *Credential) *gorm.DB {
	return configuration.Pdatabase.Save(&cred)
}

//func Signup(db *gorm.DB, req *http.Request) (*http.Request, error) {
//	//passwordHash, err :=  passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	return &req
//}
