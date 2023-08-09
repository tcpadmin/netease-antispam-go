# 易盾 - 内容安全/反垃圾 - go-sdk

说明
- 一般情况下，一个账户开通的所有业务共享 secretId 和 secretKey 但是使用不同的业务ID businessId
- 内容安全的各个业务签名算法是相同的，但是与反作弊、智能风控的签名算法不同
- 目前不支持 sm3

层级结构；只允许上级包调用下级包，不允许调用同级包
- 业务代码
- instance
  - audioV2
  - textV4
- core
  - formClient
  - jsonClient
- common
  - logger

示例
```go
logger := newLogger() //自定义的logger
audioClient := core.NewClient(secretId, secretKey, core.WithLog(logger), core.WithBizId("123456789"))
request := &audioV2.Request{DataId:"", Token:""}
client.FormPost(ctx, request)

textClient := core.NewClient(secretId, secretKey, core.WithLog(logger), core.WithBizId("abcdefg"))
requestText := &textV4.Request{DataId:"", Token:""}
textRes,err := textClient.FormPost(ctx, requestText)
```

示例2
```
// 全局变量 logger 和 client
logger := newLogger() //自定义的logger
audioClient := audioV2.NewAudioClient(secretId, secretKey, core.WithLog(logger), core.WithBizId("123456789"))
textClient := textV4.NewTextClient(secretId, secretKey, core.WithLog(logger), core.WithBizId("hello world"))

//具体业务
checkRes,err := audioClient.Check(ctx, &AudioReq{})
```

## 目前支持的业务

### 内容安全

音频检测
- 点播音频接口
  - 同步检测 https://support.dun.163.com/documents/588434426518708224?docId=588884842603749376

图片检测
- 图片接口
  - 同步检测 https://support.dun.163.com/documents/588434277524447232?docId=791822473634779136

文本检测
- 文本接口
  - 单次同步 https://support.dun.163.com/documents/588434200783982592?docId=791131792583602176