package repository

import (
	"best-practics/domain/entity"
)

type UserRepository interface {
	CreateUser(*entity.User) error
	GetUser(int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}
