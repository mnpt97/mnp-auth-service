package auth_service

import (
	"github.com/mnpt97/mnp-auth-service/models"

	"golang.org/x/crypto/bcrypt"
)

type UserDB interface {
	GetUser(string) (models.User, int, error)
	CreateUser(string, string) (models.User, int, error)
	DeleteUser(string) (int, error)
	UpdateUser(models.User) (int, error)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
