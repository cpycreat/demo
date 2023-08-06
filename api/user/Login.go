package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginReq 用户登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRes 用户登录返回
type LoginRes struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

func Login(c *gin.Context) {
	var req LoginReq
	c.BindJSON(&req)
	fmt.Println("发送的消息为", req.Password, req.Username)
	//进行登录
	res := DoLogin(&req)
	c.JSON(200, res)

}
