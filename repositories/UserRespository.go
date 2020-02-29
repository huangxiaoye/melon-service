package repositories

import (
	"melon-service/databases"
	"melon-service/models"
)

type UserRepository struct {
	DB *databases.DB
}

func NewUserRepository(db *databases.DB) UserRepository {
	return UserRepository{DB: db}
}

func (repository UserRepository) GetUserByName(name string) (models.UserModel, error) {
	var u models.UserModel

	req := repository.DB.SelectFrom("users").Where("name=?", name)
	if err := req.One(&u); err != nil {
		return models.UserModel{}, err
	}
	return u, nil
}
