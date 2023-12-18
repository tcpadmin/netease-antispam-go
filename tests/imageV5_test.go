package tests

import (
	"context"
	"fmt"
	"github.com/tcpadmin/netease-antispam-go/common"
	"github.com/tcpadmin/netease-antispam-go/core"
	"github.com/tcpadmin/netease-antispam-go/imgCheck"
	"net/url"
	"testing"
)

func TestImageV5(t *testing.T) {
	secretId := ""
	secretKey := ""
	bizId := ""

	extra := url.Values{}
	extra.Set("ip", "1.2.3.4")
	extra.Set("phone", "13800000000")
	extra.Set("isPremiumUse", "0")
	cfg := core.NewConfig(secretId, secretKey, core.WithTimeout(0)) //不设超时
	cfg.SetLogLevel(common.Info)
	client := imgCheck.NewClientV5(cfg, bizId)

	req := imgCheck.NewImgRequest(imgCheck.ImgTypeUrl, &imgCheck.ImgDto{
		Data:   "https://img.img.com/path/to/img.png",
		DataId: "11111",
	})
	req.Extra = extra

	resp, err := client.CheckRaw(context.Background(), req)
	fmt.Println(string(resp), err)
}
