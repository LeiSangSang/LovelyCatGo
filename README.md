# LovelyCatGo
go语言与可爱猫机器人HTTP插件对接sdk

## Install
```shell
go get github.com/LeiSangSang/LovelyCatGo
```

## Use
```go
//设置api地址
LovelyCatGo.SetSendUrl(`http://127.0.0.1:8073/send`)
//设置api验证key
LovelyCatGo.SetKey(`E946****************93B7F`)
//可用消息结构体
var msg LovelyCatGo.Message
//可用的消息发送方法
LovelyCatGo.SendTextMsg(msg.RobotWxId,msg.FromWxId,`hello world`)
```

## 事件类型
100                       私聊消息

200                       群聊消息

300                       暂无

400                       群成员增加

410                       群成员减少

500                       收到好友请求

600                       二维码收款

700                       收到转账

800                       软件开始启动

900                       新的账号登录完成

910                       账号下线

##可用方法
参阅 send.go