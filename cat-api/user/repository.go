package user

import "cat-api/models"

type Repository interface {
	FindUserByCredentials(login models.UserLogin) (bool, error)
}
