package persistence

import (
	"best-practics/domain/entity"
	"best-practics/domain/interface/repository"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

//UserDao implements the repository.IUserRepository interface
var _ repository.IUserRepository = &UserDao{}

func (r *UserDao) CreateUser(user *entity.User) error {
	return r.db.Create(&user).Error
}

func (r *UserDao) GetUser(id int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).Take(&user).Error
	return &user, err
}

func (r *UserDao) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserDao) GetUserByEmailAndPassword(u *entity.User) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", u.Email).Take(&user).Error
	return &user, err
}
