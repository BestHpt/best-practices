package interfaces

import (
	"best-practics/common/error_code"
	"best-practics/common/trace"
	"best-practics/domain/entity"
	"best-practics/domain/service"
	"best-practics/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//UserController struct defines the dependencies that will be used
type UserController struct {
	us service.IUserService
}

//UserController constructor
func NewUserController(us service.IUserService) *UserController {
	return &UserController{
		us: us,
	}
}

// @Tags SysUser
// @Summary 获取用户信息
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Hello victoria steven}!"}"
// @Router /user/getUserInfo [get]
func (s *UserController) GetUser(c *gin.Context) {
	ctx := trace.GetTraceCtx(c)

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage(c, error_code.ParamBindError)
		return
	}

	user, err := s.us.GetUserService(ctx, int64(userId))
	if err != nil {
		response.FailWithMessage(c, error_code.ConfigMySQLSelectError)
		return
	}
	response.SuccessWithMessage(c, fmt.Sprintf("Hello %s!", user.PublicUser()))
}

func (s *UserController) CreateUser(c *gin.Context) {
	ctx := trace.GetTraceCtx(c)
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}
	err := s.us.CreateUserService(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (s *UserController) GetUsers(c *gin.Context) {
	ctx := trace.GetTraceCtx(c)
	users := entity.Users{} //customize user
	var err error
	//us, err = application.UserApp.GetUsers()
	users, err = s.us.GetUserList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users.PublicUsers())
}

