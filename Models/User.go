package Models

type UserRegisterReq struct {
	Username string `form:"username" json:"user_name" binding:"required,alphanum"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=12"`
}

//注册用户信息返回
type LoginResponse struct {
	UserResponse
	Token string `form:"token" json:"token" `
}

type UserResponse struct {
	Id         int    `form:"id" json:"id" `
	Username   string `form:"username" json:"username" `
	Status     int    `form:"status" json:"status" ` // 0 正常状态， 1删除
	Createtime int64  `form:"createtime" json:"createtime" `
}

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}
