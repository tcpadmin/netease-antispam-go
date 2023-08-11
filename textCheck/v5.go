package textCheck

// 文本检测
// 接口文档 https://support.dun.163.com/documents/588434200783982592?docId=791131792583602176

import (
	"context"
	"errors"
	"net/url"

	"github.com/tcpadmin/netease-antispam-go/core"
)

type ClientV5 struct {
	config *core.Config
	bizId  string //全局通用，也可以在后续请求中指定
}

func NewClientV5(cfg *core.Config, bizId string) *ClientV5 {
	return &ClientV5{config: cfg, bizId: bizId}
}

// CheckRaw 原生响应值，可自行用 gJson 处理（推荐）
func (c *ClientV5) CheckRaw(ctx context.Context, req *RequestV5) ([]byte, error) {
	if req.Content == "" {
		return nil, errors.New("内容不能为空")
	}

	if req.BizId == "" {
		req.BizId = c.bizId
	}

	respByte, err := core.PostForm(ctx, c.config, req)
	if err != nil {
		return nil, err
	}
	return respByte, nil
}

type RequestV5 struct {
	BizId string //非必传；如果没有赋值，会用 client 中的字段赋值；这个字段是为了方便在请求中灵活调整 bizId；

	DataId  string //唯一标识
	Content string //聊天内容

	//其他参数和扩展参数（用户扩展参数 设备扩展参数）
	//扩展参数在内容安全业务中，是通用的 https://support.dun.163.com/documents/588434200783982592?docId=476559002902757376
	//有些自定义扩展参数，易盾对接各个甲方时使用的有可能使用不同的字段名称，以自行商定为准。也放在这里。
	Extra url.Values
}

func NewRequestV5(content, dataId string) *RequestV5 {
	return &RequestV5{Content: content, DataId: dataId}
}

func (c *RequestV5) ApiUrl() string {
	return "http://as.dun.163.com/v5/text/check"
}

func (c *RequestV5) PostData() url.Values {
	postData := url.Values{}
	//公共参数；其他参数在postForm中处理
	postData.Set("businessId", c.BizId)

	//独有参数
	postData.Set("version", "v5.3")
	postData.Set("content", c.Content)
	if c.DataId != "" {
		postData.Set("dataId", c.DataId)
	}

	//追加自定义参数
	if c.Extra != nil && len(c.Extra) != 0 {
		for k, v := range c.Extra {
			postData[k] = v
		}
	}
	return postData
}
