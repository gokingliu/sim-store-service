package entity

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk                    = Status{0, "Success"}
	ParamErr                 = Status{-1, "请求参数不合法"}
	ResFail                  = Status{-2, ""}
	JsonMarshalErr           = Status{-100, "JSON 序列化出错"}
	InnerQueryDBError        = Status{-101, "查询数据库失败"}
	InnerUpdateDBError       = Status{-102, "更新数据库失败"}
	InnerAddDBError          = Status{-103, "添加数据库失败"}
	ClientCheckUserNameError = Status{302, "用户名重复"}
	ClientUPInvalid          = Status{303, "用户名或密码不规范"}
	ClientLoginError         = Status{304, "登录失败"}
	NOAuthError              = Status{403, "无权限"}
)
