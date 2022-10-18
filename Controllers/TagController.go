package Controllers

import (
	"gin-1/Common"
	"gin-1/Models"
	"gin-1/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
}

func (t *TagController) CreateTags(c *gin.Context) {
	//校验参数
	respMsg := Common.RespMsg{}
	var TagRequest Models.TagReq
	if err := c.ShouldBind(&TagRequest); err != nil {
		respMsg.InvalidParams()
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//插入数据
	sql := "insert into tags(name,created_at) values(?,?) "
	res, err := Services.InsertDB(sql, TagRequest.Name, Services.GetFormatTime())
	if err != nil {
		respMsg.SetFailedRespMsg(Common.SQL_INSERT_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}
	//获取返回的ID
	id, err2 := res.LastInsertId()
	if err2 != nil {
		respMsg.SetFailedRespMsg(Common.SQL_INSERT_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}

	//查询数据结果
	sql2 := "select id,name,created_at from tags where id = ?"
	var TagResponse Models.TagResponse
	err3 := Services.QueryRowDB(sql2, id).Scan(&TagResponse.Id, &TagResponse.Name, &TagResponse.CreatedAt)
	if err3 != nil {
		respMsg.SetFailedRespMsg(Common.NOT_FOUND_TAG_ERROR)
		c.JSON(http.StatusOK, respMsg)
		return
	}
	respMsg.SetSuccessRespMsg(TagResponse)
	c.JSON(http.StatusOK, respMsg)
}
