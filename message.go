package LovelyCatGo

// Message
//  @Description: 消息结构体
//  @事件类型:
//   100 私聊消息
//   200 群聊消息
//   300 暂无
//   400 群成员增加
//   410 群成员减少
//   500 收到好友请求
//   600 二维码收款
//   700 收到转账
//   800 软件开始启动
//   900 新的账号登录完成
//   910 账号下线
type Message struct {
	Type          int    `json:"type"`            //事件类型（事件列表可参考 - 事件列表demo）
	MsgType       int    `json:"msg_type"`        //消息类型（仅在私聊和群消息事件中，代表消息的表现形式，如文字消息、语音、等等）
	FromWxId      string `json:"from_wxid"`       //1级来源id（比如发消息的人的id）
	FromName      string `json:"from_name"`       //1级来源昵称（比如发消息的人昵称）
	FinalFromWxId string `json:"final_from_wxid"` //2级来源id（群消息事件下，1级来源为群id，2级来源为发消息的成员id，私聊事件下都一样）
	FinalFromName string `json:"final_from_name"` //2级来源昵称
	RobotWxId     string `json:"robot_wxid"`      //当前登录的账号（机器人）标识id
	FileUrl       string `json:"file_url"`        //如果是文件消息（图片、语音、视频、动态表情），这里则是可直接访问的网络地址，非文件消息时为空
	Msg           string `json:"msg"`             //消息内容
	Parameters    string `json:"parameters"`      //附加参数（暂未用到，请忽略）
	Time          int    `json:"time"`            //请求时间(时间戳10位版本)
	Rid           int    `json:"rid"`             //请求标识
}

