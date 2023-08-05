package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RegisterReq 用户注册请求
type RegisterReq struct {
	Password string `json:"password"` // 密码，最长32个字符
	Username string `json:"username"` // 注册用户名，最长32个字符
}

// RegisterRes 用户登录返回
type RegisterRes struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

func Register(c *gin.Context) {
	var req RegisterReq
	c.BindJSON(&req)
	fmt.Println("发送的消息为", req.Password, req.Username)
	//判断用户名和密码是否过长
	if len(req.Username) > 32 || len(req.Password) > 32 {
		c.JSON(200, RegisterRes{
			StatusCode: 1,
			StatusMsg:  "注册失败,用户名或密码过长,最长32个字符",
			Token:      "nil",
			UserID:     -1,
		})
		return
	}
	res := DoRegister(&req)
	c.JSON(200, res)
}
