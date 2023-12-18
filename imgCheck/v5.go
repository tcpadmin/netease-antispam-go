package imgCheck

// 图片检测
// 接口文档：https://support.dun.163.com/documents/588434277524447232?docId=791822473634779136

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/tcpadmin/netease-antispam-go/core"
)

type ClientV5 struct {
	config *core.Config
	bizId  string //全局通用，也可以在后续请求中指定
}

func NewClientV5(cfg *core.Config, bizId string) *ClientV5 {
	return &ClientV5{config: cfg, bizId: bizId}
}

var errEmptyImg = errors.New("待审核图片不能为空")
var errTooMuch = errors.New("待审核图片不能超过32张")

// CheckRaw 原生响应值，可自行用 gJson 处理（推荐）
func (c *ClientV5) CheckRaw(ctx context.Context, req *RequestV5) ([]byte, error) {
	if len(req.Images) == 0 {
		return nil, errEmptyImg
	}
	if len(req.Images) > 32 {
		return nil, errTooMuch
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

const ImgTypeUrl = 1
const ImgTypeBase64 = 2

type RequestV5 struct {
	apiUrl string
	BizId  string //非必传；如果没有赋值，会用 client 中的字段赋值；这个字段是为了方便在请求中灵活调整 bizId；

	Images  []*ImgDto //要审核的图片列表
	ImgType int       //图片类型

	//其他参数和扩展参数（用户扩展参数 设备扩展参数）
	//扩展参数在内容安全业务中，是通用的 https://support.dun.163.com/documents/588434200783982592?docId=476559002902757376
	//有些自定义扩展参数，易盾对接各个甲方时使用的有可能使用不同的字段名称，以自行商定为准。也放在这里。
	Extra url.Values
}

type ImgDto struct {
	Name   string `json:"name"` //该字段为必须；如果没有会取 dataId 字段
	Type   int    `json:"type"`
	Data   string `json:"data"`
	DataId string `json:"dataId,omitempty"`

	CallbackUrl string `json:"callbackUrl,omitempty"`
}

// NewImgRequest 创建图片审核请求
// 此处的 imgType 会替换所有的 dto 里的 Type
func NewImgRequest(imgType int, imgItemList ...*ImgDto) *RequestV5 {
	for k, item := range imgItemList {
		item.Type = imgType
		// Name为必传字段，优先取dataId；如果为空则取序号
		if item.Name == "" {
			item.Name = item.DataId
		}
		if item.Name == "" {
			item.Name = strconv.Itoa(k)
		}
	}
	res := &RequestV5{Images: imgItemList}

	if imgType == ImgTypeBase64 {
		res.apiUrl = "http://as.dun.163.com/v5/image/base64Check"
	} else {
		res.apiUrl = "http://as.dun.163.com/v5/image/check"
	}
	return res
}

func (c *RequestV5) ApiUrl() string {
	return c.apiUrl
}

func (c *RequestV5) PostData() url.Values {
	postData := url.Values{}
	//公共参数；其他参数在postForm中处理
	postData.Set("businessId", c.BizId)

	//独有参数
	postData.Set("version", "v5.2")
	imgTodo, _ := json.Marshal(c.Images)
	postData.Set("images", string(imgTodo))

	//追加自定义参数
	if c.Extra != nil && len(c.Extra) != 0 {
		for k, v := range c.Extra {
			postData[k] = v
		}
	}
	return postData
}
