package audioCheck

// 文本、音频、图片等各个业务的标签不尽相同
// 音频检测：https://support.dun.163.com/documents/588434426518708224?docId=588884842603749376

const (
	LabelOk    = 0
	LabelPorn  = 100
	LabelAd    = 200
	LabelAdLaw = 260
)

const (
	SuggestionOk        = 0 //通过
	SuggestionSuspicion = 1 //嫌疑
	SuggestionReject    = 2 //删除
)
