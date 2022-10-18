package Models

type TagReq struct {
	Name string `form:"name" json:"name" binding:"required,alphanum,max=20"`
}

type TagResponse struct {
	Id        int    `form:"id" json:"id" `
	Name      string `form:"name" json:"name" `
	CreatedAt string `form:"created_at" json:"created_at" `
}
