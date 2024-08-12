package db_service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/mnpt97/mnp-auth-service/models"

	"github.com/google/uuid"
)

type SimpleDb struct {
	users []models.User
}

func NewSimpleDb(userDataFile string) *SimpleDb {
	db := &SimpleDb{}
	db.ReadUserData(userDataFile)
	return db
}

func (sd *SimpleDb) GetUser(username string) (models.User, int, error) {
	for _, u := range sd.users {
		if u.Username == username {
			return u, http.StatusOK, nil
		}
	}
	return models.User{}, http.StatusNotFound, errors.New("user not found")
}
func (sd *SimpleDb) CreateUser(username string, password string) (models.User, int, error) {
	return models.User{}, http.StatusBadRequest, errors.New("create new users cannot be created")
}
func (sd *SimpleDb) DeleteUser(username string) (int, error) {
	return http.StatusBadRequest, errors.New("delete user not implemented")

}
func (sd *SimpleDb) UpdateUser(user models.User) (int, error) {
	return http.StatusBadRequest, errors.New("update user not implemented")
}

func (sd *SimpleDb) ReadUserData(userDataFile string) {
	userData, err := os.ReadFile(userDataFile)
	if err != nil {
		log.Fatalf("Cannot load user data file")
	}

	var simpleUsers []models.SimpleUser
	err = json.Unmarshal(userData, &simpleUsers)
	if err != nil {
		log.Fatalf("Cannot load user data file")
	}

	var users []models.User
	for _, user := range simpleUsers {
		var claims []models.Claim
		for _, claim := range user.Claims {
			claims = append(claims, models.Claim{Key: claim.Key, Value: claim.Value})
		}
		users = append(users, models.User{ID: uuid.UUID{}.String(), Username: user.Username, PasswordEnc: user.Password, Claims: claims})

	}
	sd.users = users

}
