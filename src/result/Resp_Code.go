package result

type ResCode int

const EmptyData = ""

// 业务逻辑状态码
const (
	Success   ResCode = 200
	NeedLogin ResCode = 1000 + iota
	InvalidPermission
	InvalidParam
	WrongPassword
	NotInClass
	UnmatchedPassword
	ServerBusy
	RecordNotFound
)

// 业务逻辑状态信息描述
var recodeText = map[ResCode]string{
	Success:           "success",
	NeedLogin:         "请登录后重试",
	WrongPassword:     "用户名或密码错误",
	InvalidPermission: "无此权限",
	InvalidParam:      "参数错误",
	ServerBusy:        "服务繁忙",
	RecordNotFound:    "找不到记录",
}

// StatusText 返回状态码的文本。如果代码为空或未知状态码则返回error
func (code ResCode) StatusText() string {
	msg, ok := recodeText[code]
	if ok {
		return msg
	}
	return recodeText[ServerBusy]
}
