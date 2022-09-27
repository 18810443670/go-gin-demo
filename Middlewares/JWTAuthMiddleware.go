package Middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// 定义加密秘钥
var mySigningKey = []byte("USIJmhMEUTHu1w9i1iimEbhivj6HRd4nasFj70V1kho40YAd9i6dj01uymUGhjoP")

// 第一步：定义结构体
type MyCustomClaims struct {
	Id       int    `json:"id"`
	jwt.StandardClaims
}

// 基于JWT的认证中间件
func JWTAuthMiddleware(c *gin.Context) {
	// 从请求头中取出
	signToken := c.Request.Header.Get("Authorization")
	//fmt.Printf("signToken =  %s \n",signToken)

	signToken = strings.Replace(signToken, "Bearer ", "", 1)
	//fmt.Printf("signToken-new =  %s \n",signToken)

	if signToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"msg":  "token为空",
		})
		c.Abort()
		return
	}
	// 校验token
	myclaims, err := parserTokenWithError(signToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "token校验失败",
		})
		c.Abort()
		return
	}
	// 将用户的id放在到请求的上下文c上
	c.Set("userid", myclaims.Id)
	c.Next() // 后续的处理函数可以用过c.Get("userid")来获取当前请求的id
}


func parserTokenWithError(signToken string) (*MyCustomClaims, error) {
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(signToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("不是一个合法的token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("token过期了")
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return nil, errors.New("无法处理这个token")
		}
	} else {
		return nil, errors.New("无法处理这个token")
	}
}


func GenToken(claims MyCustomClaims) (string, error) {
	// 使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signToken, nil
}