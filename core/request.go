package core

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/tcpadmin/netease-antispam-go/common"
)

var ErrHttpCode = errors.New("HTTP状态码非200")

// IRequest 易盾的请求接口；不同的内容检查 url不同，签名计算方式不同，请求参数也不同
type IRequest interface {
	ApiUrl() string
	PostData() url.Values
}

// PostForm 使用 form 表单的方式请求，目前内容安全相关接口都是该方式
func PostForm(ctx context.Context, cfg *Config, request IRequest) ([]byte, error) {
	postData := request.PostData()
	//设置公共参数
	now := time.Now()
	postData.Set("secretId", cfg.secretID)
	postData.Set("timestamp", strconv.FormatInt(now.UnixMilli(), 10))
	postData.Set("nonce", strconv.FormatInt(rand.Int63(), 10))

	sign := GenSignature(postData, cfg.secretKey)
	postData.Set("signature", sign)

	paramsStr := postData.Encode()
	resp, err := cfg.client.PostForm(request.ApiUrl(), postData)
	consume := time.Since(now).Milliseconds()

	if err != nil {
		cfg.log.Error(ctx, "易盾PostForm错误",
			common.LogField("err", err),
			common.LogField("url", request.ApiUrl()),
			common.LogField("参数", paramsStr),
			common.LogField("耗时", consume),
		)
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		cfg.log.Error(ctx, "易盾PostForm失败",
			common.LogField("url", request.ApiUrl()),
			common.LogField("参数", paramsStr),
			common.LogField("响应", resp),
			common.LogField("耗时", consume),
		)
		return nil, ErrHttpCode
	}

	if err != nil {
		cfg.log.Error(ctx, "易盾PostForm读响应错误",
			common.LogField("err", err),
			common.LogField("url", request.ApiUrl()),
			common.LogField("参数", paramsStr),
			common.LogField("响应", resp),
			common.LogField("耗时", consume))
		return nil, err
	}

	cfg.log.Info(ctx, "易盾PostForm",
		common.LogField("url", request.ApiUrl()),
		common.LogField("参数", paramsStr),
		common.LogField("响应", string(respBytes)),
		common.LogField("耗时", consume))
	return respBytes, err
}

// GenSignature 生成签名信息
func GenSignature(params url.Values, secretKey string) string {
	var paramStr string
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		paramStr += key + params[key][0]
	}
	paramStr += secretKey
	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))
	return hex.EncodeToString(md5Reader.Sum(nil))
}
