package model

// UserLogin 用户登录表
type User struct {
	UID      int64  `gorm:"primary_key" auto_increment:"true"`
	Username string `gorm:"type:varchar(20);unique"`
	Password string `gorm:"type:varchar(20);"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}

//创建users表
func CreateUserTable() {
	if !DB.Migrator().HasTable(&User{}) {
		DB.Migrator().CreateTable(&User{})
	}
}

//对users表插入数据,返回插入User的UID
func InsertUser(user *User) (uid int64, err error) {
	err = DB.Create(user).Error
	uid = user.UID
	return
}

//通过用户名查询用户
func QueryUserByUsername(username string) (user *User, err error) {
	user = new(User)
	err = DB.Where("username=?", username).First(user).Error
	return
}
