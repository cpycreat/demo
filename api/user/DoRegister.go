package user

import (
	"dydemo/model"
	"fmt"
)

func DoRegister(req *RegisterReq) *RegisterRes {
	// 查询用户是否存在
	_, err := model.QueryUserByUsername(req.Username)
	if err == nil {
		return &RegisterRes{
			StatusCode: 1,
			StatusMsg:  "用户已存在",
		}
	}
	// 插入用户
	fmt.Println(req.Username, req.Password)
	newuser := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	uid, err := model.InsertUser(&newuser)
	if err != nil {
		return &RegisterRes{
			StatusCode: 1,
			StatusMsg:  "注册失败",
		}
	}
	//返回
	return &RegisterRes{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		Token:      req.Password + req.Username,
		UserID:     uid,
	}
}
