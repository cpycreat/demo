package user

func DoLogin(req *LoginReq) (res *LoginRes) {
	// //查询用户名是否存在
	// _, err := model.QueryUserByUsername(req.Username)
	// if err == nil {
	// 	fmt.Println("用户名存在")
	// 	//查询密码是否正确
	// 	_, err := model.QueryUserByUsernameAndPassword(req.Username, req.Password)
	// 	if err == nil {
	// 		fmt.Println("密码正确")
	// 		//返回
	// 		return &LoginRes{
	// 			StatusCode: 0,
	// 			StatusMsg:  "登录成功",
	// 			Token:      req.Password + req.Username,
	// 			UserID:     1,
	// 		}
	// 	} else { //密码错误

	// 	}
	// 	return
	// } else { //用户名不存在

	// }

	//返回
	return &LoginRes{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		Token:      req.Password + req.Username,
		UserID:     1,
	}
}
