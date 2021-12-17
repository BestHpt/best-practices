package persistence

import (
	"best-practics/common"
	"best-practics/common/initialize/viper"
	"best-practics/domain/entity"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)
var db *gorm.DB
// 这里不用小写默认，防止再测试其他单测文件时也跑到这里
func init() {
	common.Viper = viper.Init("../../conf/config.yaml") // 初始化Viper
	db = InitDB()
}

func TestCreateUser(t *testing.T) {
	var user = entity.User{}
	user.Email = "steven@example.com"
	user.FirstName = "victoria"
	user.LastName = "steven"
	user.Password = "password"
	repo := NewUserDao(db)
	CreateErr := repo.CreateUser(&user)
	assert.Nil(t, CreateErr)
}

func TestGetUser(t *testing.T) {
	repo := NewUserDao(db)
	u, getErr := repo.GetUser(1)
	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, "steven@example.com")
	assert.EqualValues(t, u.FirstName, "victoria")
	assert.EqualValues(t, u.LastName, "steven")
	fmt.Println("value is:", u)
}
