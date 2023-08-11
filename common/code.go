package common

// 接口响应码 https://support.dun.163.com/documents/588434426518708224?docId=588872765795962880

const ApiOk = 200

// result.antispam.suggestion 取值常量
const (
	SuggestionOk     = 0 //通过
	SuggestionDoubt  = 1 //嫌疑
	SuggestionReject = 2 //不通过
)

// result.antispam.Label.level 取值常量
const (
	LabelLevelOk     = 0 //正常
	LabelLevelDoubt  = 1 //不确定/嫌疑
	LabelLevelReject = 2 //确定/不通过
)
