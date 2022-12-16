package globalError

type GlobalError struct {
	Code             int    `json:"code"`    // 业务码
	Message          string `json:"message"` // 业务码
	RealErrorMessage string `json:"err_msg"`
}

func (e *GlobalError) Error() string {
	return e.Message
}

// 2、定义errorCode
const (
	AuthErr = 405

	ServerError        = 10101 // Internal Server Error
	ParamBindError     = 10102 // 参数信息有误
	AuthorizationError = 10103 // 签名信息有误
	CallHTTPError      = 10104 // 调用第三方HTTP接口失败
	ResubmitMsg        = 10105 // 请勿重复提交

	GetError    = 20101
	CreateError = 20102
	DeleteError = 20103
	UpdateError = 20104

	LoginErr  = 30101
	LogoutErr = 30102
)

// 3、定义errorCode对应的文本信息
var codeTag = map[int]string{
	AuthErr: "权限不足，请联系管理员",

	ServerError:        "Internal Server Error",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败",
	ResubmitMsg:        "请勿重复提交",

	GetError:    "查询失败",
	CreateError: "添加失败",
	UpdateError: "修改失败",
	DeleteError: "删除失败",

	LoginErr:  "登录失败",
	LogoutErr: "注销失败",
}

func GetErrorMsg(code int) string {
	return codeTag[code]
}

// NewGlobalError 4、新建自定义error实例化
func NewGlobalError(code int, err error) error {
	// 初次调用得用Wrap方法，进行实例化
	return &GlobalError{
		Code:             code,
		Message:          codeTag[code],
		RealErrorMessage: err.Error(),
	}
}
