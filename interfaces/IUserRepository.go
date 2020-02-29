package interfaces

import (
	"melon-service/models"
)

type IUserRepository interface {
	GetUserByName(name string) (models.UserModel, error)
}
