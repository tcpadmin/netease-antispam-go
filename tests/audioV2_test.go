package tests

import (
	"context"
	"fmt"
	"gitee.com/httpadmin/netease-antispam-go/audioCheck"
	"gitee.com/httpadmin/netease-antispam-go/core"
	"testing"
)

func TestAudioV2(t *testing.T) {
	cfg := core.NewConfig("123", "456", core.WithTimeout(0)) //不设超时
	audioV2Client := audioCheck.NewClientV2(cfg, "sss")
	resp, err := audioV2Client.Check(context.Background(), &audioCheck.RequestV2{})
	fmt.Println(resp, err)
}
