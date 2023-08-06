package model

// VedioList 视频列表,一对多
type Vedio struct {
	VID      int64  `gorm:"primarykey"  unique:"true"  ` //视频ID
	PlayURL  string `gorm:"type:varchar(255);"`          //视频地址
	CoverURL string `gorm:"type:varchar(255);"`          //封面地址
	AuthorID int64  `gorm:"type:int(11);"`               //作者ID

	FavoriteCount int64   `gorm:"type:int(11);"`          //点赞数
	Title         string  `gorm:"type:varchar(255);"`     //视频标题
	IsFavorite    bool    `gorm:"type:bool;"`             //是否被点赞
	Liked         []*User `gorm:"many2many:favor_vedio;"` //点赞的人
}

// TableName 设置表名
func (Vedio) TableName() string {
	return "vedio_list"
}

// InsertVedioList 插入视频列表
func InsertVedioList(vedioList *Vedio) (vid int64, err error) {
	err = DB.Create(vedioList).Error
	vid = vedioList.VID
	return
}
