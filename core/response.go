package core

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AsrDto struct {
	TaskId   string          `json:"taskId"`
	DataId   string          `json:"dataId"`
	Callback string          `json:"callback"`
	Details  []*AsrDetailDto `json:"details"`
}

type AsrDetailDto struct {
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	Content   string `json:"content"`
}
