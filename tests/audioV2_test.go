package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/tcpadmin/netease-antispam-go/audioCheck"
	"github.com/tcpadmin/netease-antispam-go/common"
	"github.com/tcpadmin/netease-antispam-go/core"
)

func TestAudioV2(t *testing.T) {
	extra := url.Values{}
	extra.Set("ip", "1.2.3.4")
	extra.Set("phone", "13800000000")
	extra.Set("isPremiumUse", "0")
	cfg := core.NewConfig("123", "456", core.WithTimeout(0)) //不设超时
	cfg.SetLogLevel(common.Info)
	audioV2Client := audioCheck.NewClientV2(cfg, "123456789")
	req := &audioCheck.RequestV2{
		AudioUrl: "https://test.test.com/test.m4a",
		Extra:    extra,
	}
	resp, err := audioV2Client.Check(context.Background(), req)
	fmt.Println(resp, err)
}
