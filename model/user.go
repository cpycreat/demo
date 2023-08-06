package model

// User 用户登录表
type User struct {
	UID             int64    `gorm:"primarykey" unique:"true"` //用户ID
	Username        string   `gorm:"type:varchar(20);unique"`  //用户名
	Password        string   `gorm:"type:varchar(20);"`        //密码
	Avatar          string   `gorm:"type:varchar(255);"`       //头像
	BackgroundImage string   `gorm:"type:varchar(255);"`       //背景图
	FavoriteCount   int64    `gorm:"type:int(11);"`            //点赞数
	FollowCount     int64    `gorm:"type:int(11);"`            //关注数
	FansCount       int64    `gorm:"type:int(11);"`            //粉丝数
	IsFollow        bool     `gorm:"type:bool;"`               //是否被关注,true-已关注，false-未关注
	Signature       string   `gorm:"type:varchar(255);"`       //个性签名
	TotalFavorited  int64    `gorm:"type:int(11);"`            //获赞数
	WorkCount       int64    `gorm:"type:int(11);"`            //作品数
	UpVedio         []*Vedio `gorm:"foreignKey:AuthorID;"`     //上传的视频
	FavorVedio      []*Vedio `gorm:"many2many:favor_vedio;"`   //点赞的视频
	Follower        []*User  `gorm:"many2many:follow;"`        //关注的人
	Following       []*User  `gorm:"many2many:follow;"`        //粉丝
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
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

//通过用户名和密码查询用户
func QueryUserByUsernameAndPassword(username string, password string) (user *User, err error) {
	user = new(User)
	err = DB.Where("username=? and password=?", username, password).First(user).Error
	return
}
