package code

const (
	/* 通用模块 start */
	CODE_OK    uint32 = 0
	CODE_ERROR uint32 = 1

	CODE_SERVER_ERROR              uint32 = 10000 // 服务器出错
	CODE_DB_ERROR                  uint32 = 10001 // 数据库出错
	CODE_REDIS_ERROR               uint32 = 10002 // redis出错
	CODE_PARAMS_LACK               uint32 = 10003 // 参数缺失
	CODE_PARAMS_INVALID            uint32 = 10004 // 参数不符合要求
	CODE_SIGNATURE_INVALID         uint32 = 10005 // 签名错误
	CODE_JSON_UNMARSHAL_ERROR      uint32 = 10006 // json反序列化失败
	CODE_SAVE_DATA_ERROR           uint32 = 10007 // 数据保存失败
	CODE_DATA_UPDATE_ERROR         uint32 = 10008 // 数据更新失败
	CODE_QUERY_DATA_ERROR          uint32 = 10009 // 数据查询失败
	CODE_DATA_NOT_FOUND            uint32 = 10010 // 记录不存在
	CODE_IMAGE_ENCODE_ERROR        uint32 = 10011 // 图片编码失败
	CODE_IMAGE_DECODE_ERROR        uint32 = 10012 // 图片解码失败
	CODE_WS_EXCEED_MAX_CONNECTIONS uint32 = 10014 // WS超出最大连接数
	CODE_HTTP_UPGRADER_FAILED      uint32 = 10015 // 协议升级失败
	CODE_JSON_MARSHAL_ERROR        uint32 = 10020 // json序列化失败
	CODE_READ_FILE_ERROR           uint32 = 10021 // 读取文件出错
	CODE_ALREADY_EXISTED_ERROR     uint32 = 10021 // 读取文件出错

	CODE_REQUEST_THIRD_ERROR               uint32 = 10100 // 请求第三方接口失败
	CODE_TMS_CLIENT_INIT_ERROR             uint32 = 10101 // 创建文本审核SDK客户端失败
	CODE_REQUEST_TMS_AUDIT_ERROR           uint32 = 10102 // 请求文本审核出错
	CODE_TMS_AUDIT_NOPASS                  uint32 = 10103 // 文本审核不通过
	CODE_GPT_CLIENT_INIT_ERROR             uint32 = 10104 // chatGPT客户端初始化失败
	CODE_SMS_CLIENT_INIT_ERROR             uint32 = 10105 // 创建短信SDK客户端失败
	CODE_REQUEST_SEND_SMS_ERROR            uint32 = 10106 // 请求发送短信失败
	CODE_IMS_CLIENT_INIT_ERROR             uint32 = 10107 // 图片审核客户端创建失败
	CODE_REQUEST_IMS_AUDIT_ERROR           uint32 = 10108 // 请求图片审核出错
	CODE_IMS_AUDIT_NOPASS                  uint32 = 10109 // 图片审核不通过
	CODE_FILE_UPLOAD_FAIL                  uint32 = 10110 // 文件上传失败
	CODE_REQUEST_FACE_IDENTIFICATION_ERROR uint32 = 10111 // 人脸识别失败
	CODE_FACE_IDENTIFICATION_NO_PASS       uint32 = 10112 // 未检测到人脸
	CODE_REQUEST_CLEAR_IMAGE_ERROR         uint32 = 10113 // 请求涂抹接口失败
	CODE_REQUEST_FACE_SWAP_ERROR           uint32 = 10114 // 请求换脸接口失败
	CODE_TMT_CLIENT_INIT_ERROR             uint32 = 10115 // 创建翻译SDK客户端失败
	CODE_REQUEST_TMT_ERROR                 uint32 = 10116 // 请求翻译失败
	CODE_REQUEST_SD_ERROR                  uint32 = 10117 // 请求SD出错
	CODE_BASE64_TO_BYTES_ERROR             uint32 = 10118 // base64转bytes失败
	CODE_REQUEST_GPT_ERROR                 uint32 = 10119 // chatGPT请求出错
	CODE_REQUEST_MJ_ERROR                  uint32 = 10120 // 请求MJ出错
	CODE_NOT_SUPPORT                       uint32 = 10121 // 参数目前不支持
	CODE_LIB_THIRD_ERROR                   uint32 = 10122 // 第三方包错误
	CODE_REDIS_LOCK_ERROR                  uint32 = 10123 // 没有获取到redis锁
	CODE_SIGN_ERROR                        uint32 = 10124 // 验签失败
	CODE_FACE_TOO_SMALL                    uint32 = 10125 // 人脸过小
	CODE_FACE_INCOMPLETE                   uint32 = 10126 // 人脸不完整
	/* 通用模块 end */

	/* 用户模块 start */
	CODE_NOT_IN_WHITE_LIST            uint32 = 11001 // 用户不在白名单
	CODE_SMS_DAY_LIMIT                uint32 = 11002 // 手机号当日发送短信达到上限
	CODE_SMS_HOUR_LIMIT               uint32 = 11003 // 手机号一小时内发送短信达到上限
	CODE_SMS_INTERVAL_LIMIT           uint32 = 11004 // 手机号距离上一次发送时间过短
	CODE_SAVE_SMS_FAIL                uint32 = 11005 // 记录短信消息失败
	CODE_SMS_NOT_FOUND                uint32 = 11006 // 找不到短信发送记录
	CODE_VERIFY_CODE_FAIL_COUNT_LIMIT uint32 = 11007 // 当前验证码已经输错3次
	CODE_VERIFY_CODE_INVALID          uint32 = 11008 // 验证码不正确
	CODE_VERIFY_CODE_EXPIRED          uint32 = 11009 // 验证码已过期
	CODE_SAVE_FAIL_COUNT_ERROR        uint32 = 11010 // 记录验证失败次数出错
	CODE_UPDATE_CODE_STATUS_ERROR     uint32 = 11011 // 更新验证码状态失败
	CODE_QUERY_USER_INFO_ERROR        uint32 = 11012 // 查询用户信息出错
	CODE_CREATE_USER_ERROR            uint32 = 11013 // 用户第一次登录创建用户出错
	CODE_GEN_TOKEN_ERROR              uint32 = 11014 // 用户登录创建token出错
	CODE_SAVE_TOKEN_ERROR             uint32 = 11015 // 设置token到redis失败
	CODE_QUERY_TOKEN_ERROR            uint32 = 11016 // 查询redis的token失败
	CODE_CANCELLATION_TOKEN_ERROR     uint32 = 11017 // 作废旧token失败

	CODE_USER_NOT_EXIST                 uint32 = 11018 // 用户不存在
	CODE_PASSWORD_ERROR                 uint32 = 11019 // 密码错误
	CODE_APPLE_LOGIN_GEN_SECRET_ERROR   uint32 = 11020 // apple secret生成出错
	CODE_APPLE_LOGIN_VALID_ERROR        uint32 = 11021 // apple登录验证不通过
	CODE_APPLE_LOGIN_GET_UNIQUEID_ERROR uint32 = 11022 // apple获取uniqueid出错
	CODE_UNAUTHORIZED                   uint32 = 11023 // 用户未登录
	CODE_TOKEN_NONACTIVATED             uint32 = 11024 // token未激活
	CODE_TOKEN_EXPIRED                  uint32 = 11025 // token过期
	CODE_TOKEN_FORMAT_INVALID           uint32 = 11026 // token格式不合法
	CODE_TOKEN_INVALID                  uint32 = 11027 // token无效
	// 18-50用作用户登录态相关
	/* 用户模块 end */

	/* 账户模块(货币系统) start */
	CodeInsufficientBalance uint32 = 12001 // 余额不足
	/* 账户模块(货币系统) end */

	/* 对话模块 start */
	CODE_TEMPLATE_NOT_FOUND        uint32 = 13001 // 未找到模板
	CODE_TEMPLATE_COVER_NOT_FOUNT  uint32 = 13002 // 未找到模板封面
	CODE_USER_MESSAGE_SAVE_FAIL    uint32 = 13003 // 用户消息保存失败
	CODE_GEN_REPLY_MESSAGE_FAIL    uint32 = 13004 // 生成回复失败
	CODE_GEN_PROMPT_FAIL           uint32 = 13005 // 生成prompt和参考图信息失败
	CODE_SAVE_PROMPT_FAIL          uint32 = 13006 // 保存消息prompt失败
	CODE_SD_GEN_IMAGE_FAIL         uint32 = 13007 // 请求SD生图失败
	CODE_MJ_GEN_IMAGE_FAIL         uint32 = 13008 // 请求MJ生图失败
	CODE_REQUEST_SD_SRV_ERROR      uint32 = 13009 // 请求SD_SRV出错
	CODE_SAVE_SD_TASK_FAIL         uint32 = 13010 // 保存SD任务失败
	CODE_CONTENT_UNAPPROVED        uint32 = 13011 // 第一次生图内容审核不通过
	CODE_CONTENT_UNAPPROVED_REPEAT uint32 = 13012 // 第二次生图内容审核不通过
	CODE_GEN_TASK_TIMEOUT          uint32 = 13013 // 生图超时
	CODE_MJ_RESULT_FAIL            uint32 = 13014 // MJ消息报错
	CODE_SD_TASK_FAIL              uint32 = 13015 // SD生图失败
	CODE_MJ_TASK_FAIL              uint32 = 13016 // MJ生图失败
	CODE_MSG_TYPE_INVALID          uint32 = 13017 // MsgType不合规
	CODE_BASIC_IMAGE_EMPTY         uint32 = 13018 // 底图不能为空
	/* 对话模块 end */

	/* 机器人模块 start */
	/* 机器人模块 end */

	/* 图片编辑模块 start */
	CODE_DOWNLOAD_BASIC_IMAGE_FAIL            uint32 = 14001 // 下载底图失败
	CODE_DOWNLOAD_MASK_IMAGE_FAIL             uint32 = 14002 // 下载蒙版失败
	CODE_REMOVE_TRANSPARENT_CHANNEL_FAIL      uint32 = 14003 // 删除透明通道失败
	CODE_DOWNLOAD_FACE_IMAGE_FAIL             uint32 = 14005 // 下载人脸图片失败
	CODE_DOWNLOAD_TEMPLATE_IMAGE_FAIL         uint32 = 14006 // 下载模板图片失败
	CODE_GEN_MASK_IMAGE_FAIL                  uint32 = 14007 // 生成蒙版图片失败
	CODE_REMOVE_MASK_TRANSPARENT_CHANNEL_FAIL uint32 = 14008 // 删除蒙版图片的透明通道失败
	CODE_NOT_NEED_OPERATION                   uint32 = 14009 // 此次操作不需要更新
	/* 图片编辑模块 end */

	/* 文件上传模块 start */
	CODE_GET_FORM_PARAM_FAIL    uint32 = 15001 // 获取表单参数出错
	CODE_GET_UPLOAD_FILE_FAIL   uint32 = 15002 // 获取上传的文件出错
	CODE_FILE_NUM_TO_LITTLE     uint32 = 15003 // 上传的文件数量小于限制的数量
	CODE_FILE_TYPE_INVALID      uint32 = 15004 // 上传的文件类型不符合要求
	CODE_SAVE_FILE_INFO_FAIL    uint32 = 15005 // 保存文件信息失败
	CODE_SAVE_FILE_CONTENT_FAIL uint32 = 15006 // 获取文件内容失败
	CODE_FACE_NUM_TO_MANY       uint32 = 15007 // 人脸数量过多
	/* 文件上传模块 end */

	/* 生图模块 start */
	CALLBACK_TASK_STATUS_INVALID uint32 = 16001 // 任务当前的状态不对
	CODE_FACESWAP_LIMIT          uint32 = 16002
	/* 生图模块 end */

	/* 活动 start */
	CODE_HAVE_ALREADY_ATTENDED = 18001 // 已经参加过
	/* 活动 end */

	/* 管理后台模块 start */
	CODE_ADMIN_PERMISSION_INVALID     uint32 = 17001 // 无访问权限
	CODE_ADMIN_AUTHENTICATE_FAIL      uint32 = 17002 // 鉴权失败
	CODE_ADMIN_LIST_PERMISSION_FAIL   uint32 = 17003 // 获取权限列表错误
	CODE_ADMIN_UPDATE_PERMISSION_FAIL uint32 = 17004 // 更新权限列表错误

	// 1-100用作用户权限相关

	CODE_TEMPLATE_CHECK_ERROR uint32 = 17102 // 模板机审报错

	/* 管理后台模块 end */

	CODE_TASK_CANCEL uint32 = 20004
)
