package models

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	CREATED:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_REPORT:             "作战报告已经存在",
	ERROR_NOT_EXIST_REPORT:         "作战报告不存在或者未找到",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "20001",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "20002",
	ERROR_AUTH_TOKEN:               "20003",
	ERROR_AUTH:                     "20004",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
