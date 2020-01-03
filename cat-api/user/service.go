package user

import "cat-api/models"

type Service interface {
	Login(login models.UserLogin) (string, bool, error)
}