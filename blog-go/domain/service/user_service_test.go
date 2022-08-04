package service

import (
	"best-practics/common"
	"best-practics/common/initialize/viper"
	"best-practics/domain/entity"
	"best-practics/infrastructure/persistence"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userSeviceInstance *userService

// 这里不用小写默认，防止再测试其他单测文件时也跑到这里
func init() {
	common.Viper = viper.Init("../../conf/config.yaml") // 初始化Viper
	db := persistence.InitDB()
	repo := persistence.NewUserDao(db)
	userSeviceInstance = &userService{repo}
}

func TestCreateUser_Success(t *testing.T) {
	user := &entity.User{
		ID:        1,
		FirstName: "victor",
		LastName:  "steven",
		Email:     "steven@example.com",
		Password:  "password",
	}
	err := userSeviceInstance.CreateUserService(context.TODO(), user)
	assert.Nil(t, err)
}

func TestGetUser_Success(t *testing.T) {
	userId := int64(1)
	u, err := userSeviceInstance.GetUserService(context.TODO(), userId)
	assert.Nil(t, err)
	assert.EqualValues(t, u.FirstName, "victoria")
	assert.EqualValues(t, u.LastName, "steven")
	assert.EqualValues(t, u.Email, "steven@example.com")
	fmt.Println("value is:", u)
}
