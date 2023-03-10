package services

import (
	"github.com/albanybuipe96/books-users-api/domain/users"
	"github.com/albanybuipe96/books-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return &user, nil
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
