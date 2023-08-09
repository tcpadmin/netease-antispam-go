package audioCheck

import (
	"context"
	"encoding/json"
	"gitee.com/httpadmin/netease-antispam-go/core"
	"net/url"
)

type ClientV2 struct {
	config *core.Config
	bizId  string
}

func NewClientV2(cfg *core.Config, bizId string) *ClientV2 {
	return &ClientV2{config: cfg, bizId: bizId}
}

func (c *ClientV2) Check(ctx context.Context, req *RequestV2) (*RespV2, error) {
	respByte, err := c.CheckRaw(ctx, req)
	if err != nil {
		return nil, err
	}
	res := &RespV2{}
	err = json.Unmarshal(respByte, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CheckRaw 原生响应值，可自行用 gJson 处理（推荐）
func (c *ClientV2) CheckRaw(ctx context.Context, req *RequestV2) ([]byte, error) {
	req.bizId = c.bizId
	respByte, err := core.PostForm(ctx, c.config, req)
	if err != nil {
		return nil, err
	}
	return respByte, nil
}

type RequestV2 struct {
	bizId string

	AudioUrl string     //音频url
	DataId   string     //唯一标识
	Custom   url.Values //自定义数据
}

func NewRequestV2(audioUrl, dataId string) *RequestV2 {
	return &RequestV2{AudioUrl: audioUrl, DataId: dataId}
}

func (c *RequestV2) ApiUrl() string {
	return "http://as.dun.163.com/v2/audio/check"
}

func (c *RequestV2) PostData() url.Values {
	postData := url.Values{}
	//公共参数；其他参数在postForm中处理
	postData.Set("businessId", c.bizId)

	//独有参数
	postData.Set("version", "v2")

	//追加自定义参数
	if c.Custom != nil && len(c.Custom) != 0 {
		for k, v := range c.Custom {
			postData[k] = v
		}
	}
	return postData
}

type RespV2 struct {
	core.BaseResp
	Result *ResultV2 `json:"result"`
}
type ResultV2 struct {
	Antispam *AntispamDtoV2 `json:"antispam"`
}

type AntispamDtoV2 struct {
	TaskId       string `json:"taskId"`
	Status       int    `json:"status "`
	Suggestion   int    `json:"suggestion"`
	ResultType   int    `json:"resultType"`
	Callback     string `json:"callback"`
	DataId       string `json:"dataId"`
	CensorSource int    `json:"censorSource"`
	Duration     int    `json:"duration"`
	CensorTime   int64  `json:"censorTime"`

	Segments []*ACV2SegmentDto `json:"segments"`
}

// ACV2SegmentDto audioCheckV2Segment
type ACV2SegmentDto struct {
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Content   string `json:"content"`
	Url       string `json:"url"`

	Labels []*ACV2LabelDto `json:"labels"`
}

type ACV2LabelDto struct {
	Label int `json:"label"` // 100：色情，200：广告，260：广告法，300：暴恐，400：违禁，500：涉政，600：谩骂，1100：涉价值观
	Level int `json:"level"` // 0：通过，1：嫌疑，2：不通过

	SubLabels []*ACV2SubLabel `json:"subLabels,omitempty"`
}

type ACV2SubLabel struct {
	SubLabel string `json:"subLabel"`

	Details json.RawMessage `json:"details"`
}
