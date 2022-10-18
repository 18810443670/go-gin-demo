package Common

const (
	SUCCESS               = 20000
	ERROR                 = 50000
	INVALID_PARAMS        = 40000
	REGISTER_ERROR        = 40001
	NOT_FOUND_USER_ERROR  = 40002
	USER_PASSWORD_ERROR   = 40003
	USERNAME_EXISTS_ERROR = 40004
	SQL_INSERT_ERROR      = 40005
	NOT_FOUND_TAG_ERROR   = 40006
)

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "error",
	INVALID_PARAMS:        "参数错误",
	REGISTER_ERROR:        "注册失败",
	NOT_FOUND_USER_ERROR:  "查无用户数据",
	USER_PASSWORD_ERROR:   "用户密码错误",
	USERNAME_EXISTS_ERROR: "用户名称已经存在",
	SQL_INSERT_ERROR:      "数据写入失败",
	NOT_FOUND_TAG_ERROR:   "标签数据不存在",
}

type RespMsg struct {
	Code int
	Data interface{}
	Msg  string
}

// 不合法参数
func (resp *RespMsg) InvalidParams() {
	resp.Code = INVALID_PARAMS
	resp.Msg = MsgFlags[INVALID_PARAMS]
	return
}

func (resp *RespMsg) SetFailedRespMsg(code int) {
	resp.Code = code
	resp.Data = make(map[string]string)
	resp.Msg = MsgFlags[code]
	return

}

// 根据响应码设置响应数据
func (resp *RespMsg) SetSuccessRespMsg(data interface{}) {
	resp.Code = SUCCESS
	resp.Data = data
	resp.Msg = MsgFlags[SUCCESS]
	return
}
