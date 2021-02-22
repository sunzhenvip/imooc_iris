package datamodels

// User模型
type User struct {
	ID       int64  `json:"id" form:"ID" sql:"ID"`
	NickName string `json:"nickName" form:"nickName" sql:"nickName"`
	UserName string `json:"userName" form:"userName" sql:"userName"`
	// json- 代表映射值为空
	HashPassword string `json:"-" form:"passWord" sql:"passWord"`
}
