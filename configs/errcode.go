package configs

// ErrCode
const (
	ERR_GETINFO = 1000
)

// ErrMap
var ErrMap = map[int]string{
	1000: "获取信息失败",
}

// GetErrMsg
func GetErrMsg(code int) string {
	return ErrMap[code]
}
