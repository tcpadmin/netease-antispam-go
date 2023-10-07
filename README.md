# 易盾 - 内容安全/反垃圾 - go-sdk

备注：
- 图片审核结果和其他的不同。根目录就是 antispam

## 业务说明
- 不用于反作弊，不用于智能风控，只用于内容安全
- 一般情况下，一个账户开通的所有业务共享 secretId 和 secretKey， 但是使用不同的业务ID（businessId）
- 内容安全的各个业务签名算法是相同的，但是与反作弊、智能风控的签名算法不同
- 不支持 sm3 签名

## 代码说明
层级结构；只允许上级包调用下级包，不允许调用同级包
- 业务代码
- instance
  - audioCheck
  - textCheck
- core
- common
  - log

## 使用说明

step1 - 初始化
```go
//全局初始化配置；可以只初始化一次；接口调用并发安全；
//但是修改日志级别并非并发安全
cfg := core.NewConfig("123", "456", core.WithTimeout(0)) //不设超时

//logger := NewLogger() //使用自定义logger
//cfg := core.NewConfig("123", "456", core.WithLog(logger))

cfg.SetLogLevel(common.Info)
audioV2Client := audioCheck.NewClientV2(cfg, "123456789")

```

step2 - 业务调用
```go
// 扩展参数
extra := url.Values{}
extra.Set("ip", "1.2.3.4")
extra.Set("phone", "13800000000")
extra.Set("isPremiumUse", "0")

req := &audioCheck.RequestV2{
    AudioUrl: "https://test.test.com/test.m4a",
    Extra:    extra,
}
resp, err := audioV2Client.Check(context.Background(), req)
fmt.Println(resp, err)
```

## 目前支持的 内容安全 业务

音频检测
- 点播音频接口 - done
  - 同步检测 https://support.dun.163.com/documents/588434426518708224?docId=588884842603749376

图片检测
- 图片接口 - todo
  - 同步检测 https://support.dun.163.com/documents/588434277524447232?docId=791822473634779136

文本检测
- 文本接口 - done
  - 单次同步 https://support.dun.163.com/documents/588434200783982592?docId=791131792583602176