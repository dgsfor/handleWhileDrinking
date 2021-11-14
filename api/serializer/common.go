package serializer

// Response 基础序列化器
// omitempty如果返回的数据这个字段为空，则序列化出来的数据没有这个字段
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

type SreResponse struct {
	Response
	ReCode int `json:"recode"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// user相关
	LOGIN_SUCCESS     = 10001 // 登录成功
	LOGIN_FAILURE     = 10002 // 登录失败
	SETCOOKIE_FAILURE = 10003 // 设置cookie失败
	LOGOUT_SUCCESS    = 10004 // 退出登录成功
	PARAMS_ERROR      = 40000 // 参数错误
	ALL_SUCCESS       = 20000 // 成功
	COM_ERROR         = 20001 // 通用错误
)
