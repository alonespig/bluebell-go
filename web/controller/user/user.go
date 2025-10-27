package user

import (
	"bluebell/froms"
	"bluebell/pkg/code"
	"bluebell/pkg/response"
	"bluebell/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 注册
// @Description 注册用户
// @Tags user
// @Accept json
// @Produce json
// @Param form body froms.RegisterForm true "注册表单"
// @Success 200 {string} json "{"code":1000,"msg":"success"}"
// @Router /api/v1/signup [post]
func Register(c *gin.Context) {
	var form froms.RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateUser(form.Username, form.Password)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func Login(c *gin.Context) {
	var form froms.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidPassword, err.Error())
		return
	}

	rsp, err := service.Login(form.Username, form.Password)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InvalidPassword, err.Error())
		return
	}

	response.Success(c, rsp)
}
