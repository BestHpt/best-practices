package service

import (
	"best-practics/domain/entity"
	"best-practics/domain/interface/repository"
	"context"
)

// 表名app依赖下一层的哪些接口，需要手工注入
type userService struct {
	userDao repository.IUserRepository
}

func NewUserService(userDao repository.IUserRepository) IUserService {
	return &userService{userDao: userDao}
}

type IUserService interface {
	CreateUserService(ctx context.Context,user *entity.User) error
	GetUserService(ctx context.Context,userId int64) (*entity.User, error)
	GetUserList(ctx context.Context) ([]entity.User, error)
	GetUserByEmailAndPasswordService(ctx context.Context,user *entity.User) (*entity.User, error)
}

func (u *userService) CreateUserService(ctx context.Context,user *entity.User) error {
	return u.userDao.CreateUser(user)
}

func (u *userService) GetUserService(ctx context.Context,userId int64) (*entity.User, error) {
	return u.userDao.GetUser(userId)
}

func (u *userService) GetUserList(ctx context.Context) ([]entity.User, error) {
	return u.userDao.GetUsers()
}

func (u *userService) GetUserByEmailAndPasswordService(ctx context.Context,user *entity.User) (*entity.User, error) {
	return u.userDao.GetUserByEmailAndPassword(user)
}
