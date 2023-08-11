package textCheck

const (
	LabelOk    = 0
	LabelPorn  = 100
	LabelAd    = 200
	LabelAdLaw = 260
)

// result.antispam.suggestion 取值常量
const (
	SuggestionOk         = 0 //通过
	SuggestionSuspicious = 1 //嫌疑
	SuggestionReject     = 2 //不通过
)

// result.antispam.Label.level 取值常量
const (
	LevelSuspicious = 1
	LevelReject     = 2 //不通过
)
