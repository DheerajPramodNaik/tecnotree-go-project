package utils

import (
	"tecnotree-go-project/entities"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(user *entities.Register) []byte {
	userpass := []byte(user.Password)

	if hashpass, err := bcrypt.GenerateFromPassword(userpass, 3); err == nil {
		return hashpass
	}
	return nil
}

func VerifyPassword(hashedPassword string, password *entities.Login) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password.Password))
}
