package audioCheck

import (
	"context"
	"github.com/tcpadmin/netease-antispam-go/audioCheck"
	"github.com/tcpadmin/netease-antispam-go/core"
	"net/url"
)

var YdConfig *core.Config
var AudioClient *audioCheck.ClientV2

const SecretId = ""
const SecretKey = ""
const bizId = ""

func init() {
	YdConfig = core.NewConfig(SecretId, SecretKey)
	AudioClient = audioCheck.NewClientV2(YdConfig, bizId)
}

func main() {
	extra := url.Values{}
	extra.Set("ip", "1.2.3.4")
	extra.Set("phone", "13800000000")
	extra.Set("isPremiumUse", "0")

	req := &audioCheck.RequestV2{
		AudioUrl: "https://test.test.com/test.m4a",
		DataId:   "111111111", //唯一标识
		Extra:    extra,
		//BizId:    "", //可选
	}
	resByte, err := AudioClient.CheckRaw(context.TODO(), req)
	if err != nil {
		// process error
	}
	_ = resByte //解析结果；建议使用 gjson; demo 见下面
}

// 获取音频审核结果
//func parseResDemo(resByte []byte) {
//	checkSuggestion := gjson.GetBytes(resByte, "result.antispam.suggestion").Int()
//	if checkSuggestion == common.SuggestionReject {
//		// 语音机审拒绝
//	}
//	if checkSuggestion == common.SuggestionDoubt {
//		// 语音机审疑似
//	}
//}

// 获取音转文结果
//func getAudioContent(resByte []byte) {
//	asrDetails := gjson.GetBytes(resByte, "asr.details")
//	if !asrDetails.Exists() {
//		return
//	}
//	var res []string
//	for _, asrItem := range asrDetails.Array() {
//		durationSeconds := asrItem.Get("startTime").Int()
//		res = append(res, fmt.Sprintf("[%02d:%02d]", durationSeconds/60, durationSeconds%60)+asrItem.Get("content").String())
//	}
//	_ = res
//}
