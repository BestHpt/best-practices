package interfaces

import (
	"best-practics/common"
	"best-practics/common/initialize/log"
	"best-practics/common/initialize/viper"
	"best-practics/domain/entity"
	"best-practics/domain/service"
	"best-practics/infrastructure/persistence"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var UserControllerInstance *UserController

func init() {
	//1、初始化Viper
	common.Viper = viper.Init("../conf/config.yaml")
	//2、初始化zap日志库
	log.InitZap()
	//3、gorm连接数据库
	repositories := persistence.NewRepositories()
	defer repositories.Close()

	userService := service.NewUserService(repositories.User)

	UserControllerInstance = NewUserController(userService)
}

//GetUser Test
func TestGetUser_Success(t *testing.T) {
	userId := strconv.Itoa(1)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/"+userId, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()

	r := gin.Default()
	// 初始化路由
	r.GET("/api/v1/users/:id", UserControllerInstance.GetUser)
	//r.Use(middleware.SetLoggerMiddleware(),middleware.GinRecovery(true))
	r.ServeHTTP(rr, req)

	var user *entity.User

	err = json.Unmarshal(rr.Body.Bytes(), &user)

	//assert.Nil(t, err)
	assert.Equal(t, rr.Code, 200)
	assert.NotNil(t, user,"返回用户为空")
	assert.EqualValues(t, user.FirstName, "victor")
	assert.EqualValues(t, user.LastName, "steven")
}

//IF YOU HAVE TIME, YOU CAN TEST ALL FAILURE CASES TO IMPROVE COVERAGE
func TestCreateUser_Success(t *testing.T) {
	inputJSON := `{
		"first_name": "victor",
		"last_name": "steven",
		"email": "steven@example.com",
		"password": "password"
	}`
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(inputJSON))
	req.Header.Add("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r := gin.Default()
	r.ServeHTTP(resp, req)

	user := &entity.User{}

	err := json.Unmarshal(resp.Body.Bytes(), &user)

	assert.Nil(t, err)
	assert.Equal(t, resp.Code, 201)
	assert.EqualValues(t, user.FirstName, "victor")
	assert.EqualValues(t, user.LastName, "steven")
}

