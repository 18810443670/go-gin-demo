package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//程序运行时间计算
func MyTime(c *gin.Context) {
	start := time.Now().UnixNano() / 1e6
	fmt.Println("程序开始：", start)
	c.Next()
	// 统计时间
	end := time.Now().UnixNano() / 1e6
	fmt.Println("程序结束：", end)
	since := end - start
	fmt.Println("程序用时：", since ,"ms")
}