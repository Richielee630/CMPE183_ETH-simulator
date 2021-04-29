package e

var MsgFlags = map[int]string{
	SUCCESS:                                "ok",
	ERROR:                                  "fail",
	INVALID_PARAMS:                         "请求参数错误",
	ERROR_JSON_PARSING:                     "JSON 解析错误",
	ERROR_DB_NOT_FIND_BLOCK:                "该区块不存在",
	ERROR_DB_NOT_FIND_TRANSACTION:          "该交易不存在",
	ERROR_DB_NOT_FIND_ACCOUNT:              "该地址不存在",
	ERROR_DB_SERARCH_LIST:                  "分页查询溢出",
	ERROR_DB_SERARCH_LIST_IS_NULL:          "列表为空",
	ERROR_SYSTEM_GET_HEIGHT:                "系统获取当前高度错误",
	ERROR_SYSTEM_SERACH_DB:                 "数据库查找错误",
	ERROR_NODE_GET_NONCE:                   "获取用户随机值",
	ERROR_SEND_MSG_TO_CONTRACT:             "往合约发送交易失败",
	ERROR_GET_MSG_FROM_CONTRACT:            "从合约查询数据 NOT FOUND",
	ERROR_DB_NOT_FIND_FILE:                 "未发现文件",
	ERROR_DB_FIND_FILE_FAIL:                "查找文件失败",
	ERROR_FS_CID_FILENAME_TO_CONTRACT_FAIL: "文件系统 cid，文件名上链失败",
	ERROR_FS_CONTRACT_THIS_KEY_IS_USED:     "该文件cid已上链",
	ERROR_AUTH_CHECK_TOKEN_FAIL:            "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:         "Token已超时",
	ERROR_AUTH_TOKEN:                       "Token生成失败",
	ERROR_AUTH:                             "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:           "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:          "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT:        "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
