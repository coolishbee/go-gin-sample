package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "유효하지 않은 파라미터입니다",

	ERROR_NOT_EXIST_USER: "이미 존재하는 사용자입니다.",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "인증 토큰이 만료되었습니다.",
	ERROR_AUTH_INVALID_TOKEN:       "토큰 정보가 일치하지 않습니다.",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "인증 시간초과",
	ERROR_INVALID_LOGIN_TYPE:       "잘못된 로그인 타입을 요청했습니다.",

	ERROR_DB: "DB 오류",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
