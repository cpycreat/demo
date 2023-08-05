package user

import (
	"dydemo/model"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func DoRegister(req *RegisterReq) *RegisterRes {
	// 插入用户
	fmt.Println(req.Username, req.Password)
	newuser := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	uid, err := model.InsertUser(&newuser)
	if err != nil {
		//判断是否是用户名重复
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			// 判断是否是唯一约束冲突
			if mysqlErr.Number == 1062 {
				return &RegisterRes{
					StatusCode: 1,
					StatusMsg:  "注册失败,用户名已存在",
					Token:      "nil",
					UserID:     -1,
				}
			} else {
				return &RegisterRes{
					StatusCode: 1,
					StatusMsg:  "注册失败," + err.Error(),
					Token:      "nil",
					UserID:     -1,
				}
			}
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
