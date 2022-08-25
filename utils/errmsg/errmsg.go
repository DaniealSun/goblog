package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code以1000开头 用户模块错误
	ERROR_USERNAME_USED     = 1001
	ERROR_PASSWORD_WRONG    = 1002
	ERROR_USER_NOT_EXIST    = 1003
	ERROR_TOKEN_NOT_EXIST   = 1004
	ERROR_TOKEN_TIMEOUT     = 1005
	ERROR_TOKEN_WRONG       = 1006
	ERROR_TOKEN_TYPE_WRONG  = 1007
	ERROR_USER_NO_PRIVILEGE = 1008
	// code以2000开头 文章模块错误
	ERROR_ART_NOT_EXIST = 2001
	// code以3000开头 分类模块错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCSE:                  "OK",
	ERROR:                   "FAIL",
	ERROR_USERNAME_USED:     "用户名已存在",
	ERROR_PASSWORD_WRONG:    "密码错误",
	ERROR_USER_NOT_EXIST:    "用户不存在",
	ERROR_USER_NO_PRIVILEGE: "用户无权限",
	ERROR_TOKEN_NOT_EXIST:   "Token不存在",
	ERROR_TOKEN_TIMEOUT:     "Token已过期",
	ERROR_TOKEN_WRONG:       "Token不正确",
	ERROR_TOKEN_TYPE_WRONG:  "Token格式不正确",

	ERROR_CATENAME_USED:  "分类已存在",
	ERROR_CATE_NOT_EXIST: "分类不存在",

	ERROR_ART_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
