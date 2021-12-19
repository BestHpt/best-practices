package repository

import (
	"best-practics/domain/entity"
)

type IUserRepository interface {
	CreateUser(*entity.User) error
	GetUser(int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}
