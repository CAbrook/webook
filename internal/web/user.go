package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	emailRegExp    *regexp.Regexp
	passwordRegExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailRegExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (c *UserHandler) RegisterRoutes(server *gin.Engine) {
	// server.POST("/users", c.SignUp) 严格的REST
	ug := server.Group("/users")
	ug.POST("/signup", c.SignUp)
	ug.POST("/login", c.Login)
	ug.POST("/edit", c.Edit)
	ug.POST("/profile", c.Profile)

}

const (
	emailRegexPattern = `^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin             = "login"
)

func (c *UserHandler) SignUp(context *gin.Context) {
	type SignReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignReq
	if err := context.Bind(&req); err != nil {
		return
	}
	isEmail, _ := c.emailRegExp.MatchString(req.Email)
	if !isEmail {
		context.String(http.StatusOK, "email error")
		return
	}
	isPassword, _ := c.passwordRegExp.MatchString(req.Password)
	if !isPassword {
		context.String(http.StatusOK, "password error")
		return
	}

	if req.Password != req.ConfirmPassword {
		context.String(http.StatusOK, "password not same")
		return
	}
	context.String(http.StatusOK, "login success")
}

func (c *UserHandler) Login(context *gin.Context) {

}

func (c *UserHandler) Edit(context *gin.Context) {

}

func (c *UserHandler) Profile(context *gin.Context) {

}
