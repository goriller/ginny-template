package constants

// ErrCode
const (
	SUCCESS        = 0
	FAIL           = 1
	PARAMS_INVALID = 2
	NOT_FOUND      = 3
	INTERNAL_ERR   = 4
	UNAUTHORIZED   = 5
)

// ErrMap
var ErrMap = map[int]string{
	0: "成功",
	1: "失败",
	2: "参数不合法",
	3: "请求的资源不存在",
	4: "服务器内部错误",
	5: "非法访问",
}

// GetErrMsg
func GetErrMsg(code int) string {
	return ErrMap[code]
}
