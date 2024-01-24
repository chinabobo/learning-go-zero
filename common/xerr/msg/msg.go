package msg

import "lightleap-c/common/xerr/code"

var MsgMap = map[uint32]string{
	code.CODE_OK:    "OK",
	code.CODE_ERROR: "ERROR",

	code.CODE_SERVER_ERROR:              "Server error",
	code.CODE_DB_ERROR:                  "DB error",
	code.CODE_REDIS_ERROR:               "Cache error",
	code.CODE_PARAMS_LACK:               "parameter missing",
	code.CODE_PARAMS_INVALID:            "parameter invalid",
	code.CODE_SIGNATURE_INVALID:         "signature invalid",
	code.CODE_JSON_UNMARSHAL_ERROR:      "json unmarshal error",
	code.CODE_SAVE_DATA_ERROR:           "save data error",
	code.CODE_DATA_UPDATE_ERROR:         "update data error",
	code.CODE_QUERY_DATA_ERROR:          "query data error",
	code.CODE_DATA_NOT_FOUND:            "data not found",
	code.CODE_IMAGE_ENCODE_ERROR:        "image encode error",
	code.CODE_IMAGE_DECODE_ERROR:        "image decode error",
	code.CODE_UNAUTHORIZED:              "unauthorized",
	code.CODE_WS_EXCEED_MAX_CONNECTIONS: "WS exceed max connections",
	code.CODE_HTTP_UPGRADER_FAILED:      "HTTP upgrade failed",
	code.CODE_TOKEN_NONACTIVATED:        "token not activated",
	code.CODE_TOKEN_EXPIRED:             "token expired",
	code.CODE_TOKEN_FORMAT_INVALID:      "token format invalid",
	code.CODE_TOKEN_INVALID:             "token invalid",
	code.CODE_READ_FILE_ERROR:           "read file error",

	code.CODE_REQUEST_THIRD_ERROR:               "request third error",
	code.CODE_TMS_CLIENT_INIT_ERROR:             "create text audit sdk client error",
	code.CODE_REQUEST_TMS_AUDIT_ERROR:           "request text audit error",
	code.CODE_TMS_AUDIT_NOPASS:                  "text audit not pass",
	code.CODE_GPT_CLIENT_INIT_ERROR:             "create chatGPT sdk client error",
	code.CODE_SMS_CLIENT_INIT_ERROR:             "create sms sdk client error",
	code.CODE_REQUEST_SEND_SMS_ERROR:            "request send sms error",
	code.CODE_IMS_CLIENT_INIT_ERROR:             "create image audit sdk client error",
	code.CODE_REQUEST_IMS_AUDIT_ERROR:           "request image audit error",
	code.CODE_IMS_AUDIT_NOPASS:                  "image audit not pass",
	code.CODE_FILE_UPLOAD_FAIL:                  "file upload fail",
	code.CODE_REQUEST_FACE_IDENTIFICATION_ERROR: "face identification error",
	code.CODE_FACE_IDENTIFICATION_NO_PASS:       "face identification no pass",
	code.CODE_REQUEST_CLEAR_IMAGE_ERROR:         "request clear image error",
	code.CODE_REQUEST_FACE_SWAP_ERROR:           "request face swap error",
	code.CODE_TMT_CLIENT_INIT_ERROR:             "create translate sdk client error",
	code.CODE_REQUEST_TMT_ERROR:                 "request translate error",
	code.CODE_REQUEST_SD_ERROR:                  "request SD error",
	code.CODE_REQUEST_GPT_ERROR:                 "request chatGPT error",
	code.CODE_BASE64_TO_BYTES_ERROR:             "base64 to bytes error",
	code.CODE_REQUEST_MJ_ERROR:                  "request MJ error",
	code.CODE_SIGN_ERROR:                        "sign error",
	code.CODE_FACE_TOO_SMALL:                    "face too small",
	code.CODE_FACE_INCOMPLETE:                   "face incomplete",
	/* 通用模块 end */

	/* 账户模块(货币系统) start */
	code.CodeInsufficientBalance: "Insufficient balance",
	/* 账户模块(货币系统) end */

	/* 对话模块 start */
	code.CODE_TEMPLATE_NOT_FOUND:        "template not found",
	code.CODE_TEMPLATE_COVER_NOT_FOUNT:  "template cover not found",
	code.CODE_USER_MESSAGE_SAVE_FAIL:    "user message save fail",
	code.CODE_GEN_REPLY_MESSAGE_FAIL:    "gen reply message fail",
	code.CODE_GEN_PROMPT_FAIL:           "gen prompt fail",
	code.CODE_SAVE_PROMPT_FAIL:          "save prompt fail",
	code.CODE_SD_GEN_IMAGE_FAIL:         "SD gen image fail",
	code.CODE_MJ_GEN_IMAGE_FAIL:         "MJ gen image fail",
	code.CODE_REQUEST_SD_SRV_ERROR:      "request SD srv error",
	code.CODE_SAVE_SD_TASK_FAIL:         "save SD task fail",
	code.CODE_CONTENT_UNAPPROVED:        "content unapproved",
	code.CODE_CONTENT_UNAPPROVED_REPEAT: "content unapproved repeat",
	code.CODE_GEN_TASK_TIMEOUT:          "gen task timeout",
	code.CODE_MJ_RESULT_FAIL:            "MJ result fail",
	code.CODE_SD_TASK_FAIL:              "SD task fail",
	code.CODE_MJ_TASK_FAIL:              "MJ task fail",
	code.CODE_MSG_TYPE_INVALID:          "msg type invalid",
	code.CODE_BASIC_IMAGE_EMPTY:         "basic image empty",
	/* 对话模块 end */

	/* 机器人模块 start */
	/* 机器人模块 end */

	/* 图片编辑模块 start */
	code.CODE_DOWNLOAD_BASIC_IMAGE_FAIL:       "download basic image fail",
	code.CODE_DOWNLOAD_MASK_IMAGE_FAIL:        "download mask image fail",
	code.CODE_REMOVE_TRANSPARENT_CHANNEL_FAIL: "remove transparent channel fail",
	// 请求涂抹接口失败
	code.CODE_DOWNLOAD_FACE_IMAGE_FAIL:             "download face image fail",
	code.CODE_DOWNLOAD_TEMPLATE_IMAGE_FAIL:         "download template image fail",
	code.CODE_GEN_MASK_IMAGE_FAIL:                  "gen mask image fail",
	code.CODE_REMOVE_MASK_TRANSPARENT_CHANNEL_FAIL: "remove mask transparent channel fail",
	code.CODE_NOT_NEED_OPERATION:                   "not need operation",
	/* 图片编辑模块 end */

	/* 文件上传模块 start */
	code.CODE_GET_FORM_PARAM_FAIL:    "get form param fail",
	code.CODE_GET_UPLOAD_FILE_FAIL:   "get upload file fail",
	code.CODE_FILE_NUM_TO_LITTLE:     "file num to little",
	code.CODE_FILE_TYPE_INVALID:      "file type invalid",
	code.CODE_SAVE_FILE_INFO_FAIL:    "save file info fail",
	code.CODE_SAVE_FILE_CONTENT_FAIL: "save file content fail",
	code.CODE_FACE_NUM_TO_MANY:       "face num to many",
	/* 文件上传模块 end */

	/* 图片生成 start */
	code.CODE_FACESWAP_LIMIT: "face swap limit",
	/* 图片生成 end */

	/* 用户模块 start */
	code.CODE_NOT_IN_WHITE_LIST:              "not in white list",
	code.CODE_SMS_DAY_LIMIT:                  "sms day limit",
	code.CODE_SMS_HOUR_LIMIT:                 "sms hour limit",
	code.CODE_SMS_INTERVAL_LIMIT:             "sms interval limit",
	code.CODE_SAVE_SMS_FAIL:                  "save sms fail",
	code.CODE_SMS_NOT_FOUND:                  "sms not found",
	code.CODE_VERIFY_CODE_FAIL_COUNT_LIMIT:   "verify code fail count limit",
	code.CODE_VERIFY_CODE_INVALID:            "verify code invalid",
	code.CODE_VERIFY_CODE_EXPIRED:            "verify code expired",
	code.CODE_SAVE_FAIL_COUNT_ERROR:          "save fail count error",
	code.CODE_UPDATE_CODE_STATUS_ERROR:       "update code status error",
	code.CODE_QUERY_USER_INFO_ERROR:          "query user info error",
	code.CODE_CREATE_USER_ERROR:              "create user error",
	code.CODE_GEN_TOKEN_ERROR:                "gen token error",
	code.CODE_SAVE_TOKEN_ERROR:               "save token error",
	code.CODE_QUERY_TOKEN_ERROR:              "query token error",
	code.CODE_CANCELLATION_TOKEN_ERROR:       "cancellation token error",
	code.CODE_USER_NOT_EXIST:                 "user not exist",
	code.CODE_PASSWORD_ERROR:                 "password error",
	code.CODE_APPLE_LOGIN_GEN_SECRET_ERROR:   "apple secret gen error",
	code.CODE_APPLE_LOGIN_VALID_ERROR:        "apple login valid error",
	code.CODE_APPLE_LOGIN_GET_UNIQUEID_ERROR: "apple login get uniqueid error",
	/* 用户 end */

	/* 活动 start */
	code.CODE_HAVE_ALREADY_ATTENDED: "have already attended",
	/* 活动 end */

	/* 管理后台模块 start */
	code.CODE_ADMIN_PERMISSION_INVALID:     "User does not have permission",
	code.CODE_ADMIN_AUTHENTICATE_FAIL:      "Authenticate error",
	code.CODE_ADMIN_LIST_PERMISSION_FAIL:   "List permission error",
	code.CODE_ADMIN_UPDATE_PERMISSION_FAIL: "Update permission error",
	code.CODE_TEMPLATE_CHECK_ERROR:         "template check error",
	/* 管理后台模块 end */
}

func GetMsg(code uint32) string {
	if msg, ok := MsgMap[code]; ok {
		return msg
	}
	return ""
}

func IsCodeErr(code uint32) bool {
	_, ok := MsgMap[code]
	return ok
}
