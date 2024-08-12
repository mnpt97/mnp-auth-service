package db_service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/mnpt97/mnp-auth-service/models"
	"google.golang.org/api/iterator"
)

type FirestoreDb struct {
	client *firestore.Client
}

func NewFirestoreDb(projectId string) (*FirestoreDb, error) {
	client, err := initClient(projectId)
	if err != nil {
		return nil, err
	}
	return &FirestoreDb{client: client}, nil
}

func initClient(projectId string) (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err

	}
	return client, err
}

func (fs *FirestoreDb) GetUser(email string) (models.User, int, error) {
	ctx := context.Background()
	iter := fs.client.Collection("users").Where("email", "==", email).Documents(ctx)
	userCount := 0

	var user models.User

	userData, code, err := func(iter *firestore.DocumentIterator) (map[string]interface{}, int, error) {
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				if userCount <= 1 {
					return doc.Data(), http.StatusOK, nil
				} else {
					return make(map[string]interface{}), http.StatusBadRequest, errors.New("more than one user with the same email address")
				}
			}
			userCount += 1
			if err != nil {
				return make(map[string]interface{}), http.StatusInternalServerError, errors.New("something went wrong")
			}
			return make(map[string]interface{}), http.StatusBadRequest, errors.New("bad request")
		}
	}(iter)

	if err != nil {
		return user, code, err
	}

	jsonBody, err := json.Marshal(userData)
	if err != nil {
		return user, http.StatusInternalServerError, errors.New("something went wrong")
	}
	if err = json.Unmarshal(jsonBody, &user); err != nil {
		return user, http.StatusInternalServerError, errors.New("something went wrong")
	}
	return user, http.StatusOK, nil
}

func (fs *FirestoreDb) CreateUser(username string, password string) (models.User, int, error) {
	return models.User{}, http.StatusBadRequest, errors.New("new users cannot be created")
}
func (fs *FirestoreDb) DeleteUser(username string) (int, error) {
	return http.StatusBadRequest, errors.New("delete user not implemented")

}
func (fs *FirestoreDb) UpdateUser(user models.User) (int, error) {
	return http.StatusBadRequest, errors.New("update user not implemented")
}
