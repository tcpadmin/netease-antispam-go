package tests

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/url"
	"testing"

	"github.com/tcpadmin/netease-antispam-go/common"
	"github.com/tcpadmin/netease-antispam-go/core"
	"github.com/tcpadmin/netease-antispam-go/textCheck"
)

func TestTextV5(t *testing.T) {
	secretId := ""
	secretKey := ""
	bizId := ""

	extra := url.Values{}
	extra.Set("ip", "1.2.3.4")
	extra.Set("phone", "13800000000")
	extra.Set("isPremiumUse", "0")
	cfg := core.NewConfig(secretId, secretKey, core.WithTimeout(0)) //不设超时
	cfg.SetLogLevel(common.Info)
	client := textCheck.NewClientV5(cfg, bizId)

	content := "正规发票。欢迎来我们平台；"
	dataId := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	req := &textCheck.RequestV5{
		Content: content,
		DataId:  dataId,
		Extra:   extra,
	}
	resp, err := client.CheckRaw(context.Background(), req)
	fmt.Println(string(resp), err)
}
