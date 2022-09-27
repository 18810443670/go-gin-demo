package Controllers

import (
	"database/sql"
	"gin-1/Common"
	"gin-1/Middlewares"
	"gin-1/Models"
	"gin-1/Services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type UserController struct {
}

//注册
func (u *UserController) Register(c *gin.Context) {
	//校验参数
	respMsg := Common.RespMsg{}
	var UserRegisterReq Models.UserRegisterReq
	if err := c.ShouldBind(&UserRegisterReq); err != nil {
		respMsg.InvalidParams()
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//判断用户是否存在
	var userInfo Models.User
	getUserSql := "select * from users where username = ?"
	Services.DB.QueryRow(getUserSql, UserRegisterReq.Username).Scan(&userInfo.Id, &userInfo.Username, &userInfo.Password, &userInfo.Status, &userInfo.Createtime)
	if userInfo.Id > 0 {
		respMsg.SetFailedRespMsg(Common.USERNAME_EXISTS_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//密码加密
	md5Password := Services.Md5(UserRegisterReq.Password)

	//插入数据库
	sql := "insert into users(username,password,status,createtime) values (?,?,?,?)"
	_, err := Services.ModifyDB(sql, UserRegisterReq.Username, md5Password, 1, time.Now().Unix())
	if err != nil {
		respMsg.SetFailedRespMsg(Common.REGISTER_ERROR)
		c.JSON(http.StatusOK, respMsg)
	} else {
		respMsg.SetSuccessRespMsg("")
		c.JSON(http.StatusOK, respMsg)
	}
}

//登录
func (u *UserController) Login(c *gin.Context) {

	//校验参数
	respMsg := Common.RespMsg{}
	var UserRegisterReq Models.UserRegisterReq
	if err := c.ShouldBind(&UserRegisterReq); err != nil {
		respMsg.InvalidParams()
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//获取用户数据
	var userInfo Models.User
	getUserSql := "select * from users where username = ?"
	err := Services.DB.QueryRow(getUserSql, UserRegisterReq.Username).Scan(&userInfo.Id, &userInfo.Username, &userInfo.Password, &userInfo.Status, &userInfo.Createtime)
	if err == sql.ErrNoRows {
		respMsg.SetFailedRespMsg(Common.NOT_FOUND_USER_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//校验密码是否正确
	md5Password := Services.Md5(UserRegisterReq.Password)
	if md5Password != userInfo.Password {
		respMsg.SetFailedRespMsg(Common.USER_PASSWORD_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//获取签名
	claims := Middlewares.MyCustomClaims{
		userInfo.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * time.Hour).Unix(), // 过期时间7小时
			Issuer:    "Go-GIN",                             // 签发人
		},
	}
	signToken, _ := Middlewares.GenToken(claims)

	//返回数据
	var loginResponse Models.LoginResponse
	loginResponse.Id = userInfo.Id
	loginResponse.Username = userInfo.Username
	loginResponse.Status = userInfo.Status
	loginResponse.Token = signToken
	respMsg.SetSuccessRespMsg(loginResponse)
	c.JSON(http.StatusOK, respMsg)
}

//获取用户信息
func (u *UserController) GetInfo(c *gin.Context) {
	respMsg := Common.RespMsg{}
	userid := c.GetInt("userid")
	var userInfo Models.UserResponse
	getUserSql := "select id,username,status,createtime from users where id = ?"
	err := Services.DB.QueryRow(getUserSql, userid).Scan(&userInfo.Id, &userInfo.Username, &userInfo.Status, &userInfo.Createtime)
	if err == sql.ErrNoRows {
		respMsg.SetFailedRespMsg(Common.NOT_FOUND_USER_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}
	respMsg.SetSuccessRespMsg(userInfo)
	c.JSON(http.StatusOK, respMsg)
}
