package LovelyCatGo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	sendUrl string //可爱猫api地址
	key     string //可爱猫api验证key
)

func init() {
	sendUrl = `http://127.0.0.1:8073/send`
	key = ``
}

//SetSendUrl 设置可爱猫API地址
func SetSendUrl(url string) {
	sendUrl = url
}

//SetKey 设置API验证KEY
func SetKey(s string) {
	key = s
}

//send 发送消息结构体
type send struct {
	Type      int         `json:"type,omitempty"`       //消息类型
	Msg       interface{} `json:"msg,omitempty"`        //消息体
	ToWxId    string      `json:"to_wxid,omitempty"`    //发送的用户id或群id
	RobotWxId string      `json:"robot_wxid,omitempty"` //机器人账户id
	AtWxId    string      `json:"at_wxid,omitempty"`    //@的用户id
	AtName    string      `json:"at_name,omitempty"`    //@的用户昵称
	IsRefresh int         `json:"is_refresh,omitempty"` //是否刷新缓存获取
	Key       string      `json:"key,omitempty"`        //API的验证key
}

//sendFriend 发送消息(好友相关)结构体
type sendFriend struct {
	Type       int         `json:"type,omitempty"`        //消息类型
	Msg        interface{} `json:"msg,omitempty"`         //消息体
	RobotWxId  string      `json:"robot_wxid,omitempty"`  //机器人账户id
	FriendWxId string      `json:"friend_wxid,omitempty"` //好友id
	Note       string      `json:"note,omitempty"`        //备注
	IsRefresh  int         `json:"is_refresh,omitempty"`  //是否刷新缓存获取
	Key        string      `json:"key,omitempty"`         //API的验证key
}

//sendGroup 发送消息(群聊相关)结构体
type sendGroup struct {
	Type       int         `json:"type,omitempty"`        //消息类型
	Msg        interface{} `json:"msg,omitempty"`         //消息体
	RobotWxId  string      `json:"robot_wxid,omitempty"`  //机器人账户id
	GroupWxId  string      `json:"group_wxid,omitempty"`  //群id
	GroupName  string      `json:"group_name,omitempty"`  //群名
	MemberWxId string      `json:"member_wxid,omitempty"` //群员id
	FriendWxId string      `json:"friend_wxid,omitempty"` //好友id
	Notice     string      `json:"notice,omitempty"`      //群公告
	Friends    []string    `json:"friends,omitempty"`     //拉新群员切片
	IsRefresh  int         `json:"is_refresh,omitempty"`  //是否刷新缓存获取
	Key        string      `json:"key,omitempty"`         //API的验证key
}

//link 分享链接结构体
type link struct {
	Title string `json:"title"` //连接标题
	Text  string `json:"text"`  //连接正文
	Url   string `json:"url"`   //连接地址
	Pic   string `json:"pic"`   //连接图片
}

// SendTextMsg
//  @Description: 发送文本消息给用户或群聊
//  @param robotWxid 发送机器人id
//  @param toWxid 接受方id(群或用户)
//  @param msg 消息内容
//  @return []byte
//  @return error
func SendTextMsg(robotWxid, toWxid, msg string) ([]byte, error) {
	data := &send{
		Type:      100,
		Msg:       url.QueryEscape(msg),
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendGroupAtMsg
//  @Description: 发送群聊消息并at某人
//  @param robotWxid 发送机器人id
//  @param toWxid 接受群id
//  @param atWxid at的用户id
//  @param atName at的用户昵称
//  @param msg 消息内容
//  @return []byte
//  @return error
func SendGroupAtMsg(robotWxid, toWxid, atWxid, atName, msg string) ([]byte, error) {
	data := &send{
		Type:      102,
		Msg:       url.QueryEscape(msg),
		ToWxId:    toWxid,
		AtWxId:    atWxid,
		AtName:    url.QueryEscape(atName),
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendImageMsg
//  @Description: 发送图片
//  @param robotWxid 发送机器人id
//  @param toWxid 接受方id(群或用户)
//  @param path 图片的绝对路径
//  @return []byte
//  @return error
func SendImageMsg(robotWxid, toWxid, path string) ([]byte, error) {
	data := &send{
		Type:      103,
		Msg:       path,
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendVideoMsg
//  @Description: 发送视频
//  @param robotWxid 发送机器人id
//  @param toWxid 接受方id(群或用户)
//  @param path 视频的绝对路径
//  @return []byte
//  @return error
func SendVideoMsg(robotWxid, toWxid, path string) ([]byte, error) {
	data := &send{
		Type:      104,
		Msg:       path,
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendFileMsg
//  @Description: 发送文件
//  @param robotWxid 发送机器人id
//  @param toWxid 接受方id(群或用户)
//  @param path 文件的绝对路径
//  @return []byte
//  @return error
func SendFileMsg(robotWxid, toWxid, path string) ([]byte, error) {
	data := &send{
		Type:      105,
		Msg:       path,
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendEmojiMsg
//  @Description: 发送动态表情
//  @param robotWxid 发送机器人id
//  @param toWxid 接受方id(群或用户)
//  @param path 表情(一般是gif)的绝对路径
//  @return []byte
//  @return error
func SendEmojiMsg(robotWxid, toWxid, path string) ([]byte, error) {
	data := &send{
		Type:      106,
		Msg:       path,
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendShareLinkMsg
//  @Description: 发送分享链接
//  @param robotWxid 发送机器人id
//  @param toWxid 接收方id(群或用户)
//  @param title 链接标题
//  @param text 链接内容
//  @param targetUrl 跳转链接
//  @param picUrl 图片链接
//  @return []byte
//  @return error
func SendShareLinkMsg(robotWxid, toWxid, title, text, targetUrl, picUrl string) ([]byte, error) {
	linkData := &link{
		Title: url.QueryEscape(title),
		Text:  url.QueryEscape(text),
		Url:   targetUrl,
		Pic:   picUrl,
	}
	data := &send{
		Type:      107,
		Msg:       linkData,
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// SendMusicMsg
//  @Description: 发送音乐分享
//  @param robotWxid 发送机器人id
//  @param toWxid 对方的id，可以是群或者好友id
//  @param name 歌曲名字
//  @return []byte
//  @return error
func SendMusicMsg(robotWxid, toWxid, name string) ([]byte, error) {
	data := &send{
		Type:      108,
		Msg:       url.QueryEscape(name),
		ToWxId:    toWxid,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// GetRobotName
//  @Description: 取当前机器人的昵称
//  @param robotWxid 机器人id
//  @return []byte
//  @return error
func GetRobotName(robotWxid string) ([]byte, error) {
	data := &send{
		Type:      201,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// GetRobotHeadImgUrl
//  @Description: 取当前机器人的头像
//  @param robotWxid 机器人id
//  @return []byte
//  @return error
func GetRobotHeadImgUrl(robotWxid string) ([]byte, error) {
	data := &send{
		Type:      202,
		RobotWxId: robotWxid,
		Key:       key,
	}
	return data.Send()
}

// GetLoggedAccountList
//  @Description: 取登录账号列表
//  @return []byte
//  @return error
func GetLoggedAccountList() ([]byte, error) {
	data := &send{
		Type: 203,
		Key:  key,
	}
	return data.Send()
}

// GetFriendList
//  @Description: 取好友列表
//  @param robotWxid 机器人id(如果填空字符串，即取所有登录账号的好友列表)
//  @param isRefresh 是否刷新获取:0从缓存获取,1刷新并获取
//  @return []byte
//  @return error
func GetFriendList(robotWxid string, isRefresh int) ([]byte, error) {
	data := &send{
		Type:      204,
		RobotWxId: robotWxid,
		IsRefresh: isRefresh,
		Key:       key,
	}
	return data.Send()
}

// GetGroupList
//  @Description: 取群聊列表
//  @param robotWxid 机器人id(如果填空字符串，即取所有登录账号的群聊列表)
//  @param isRefresh 是否刷新获取:0从缓存获取,1刷新并获取
//  @return []byte
//  @return error
func GetGroupList(robotWxid string, isRefresh int) ([]byte, error) {
	data := &send{
		Type:      205,
		RobotWxId: robotWxid,
		IsRefresh: isRefresh,
		Key:       key,
	}
	return data.Send()
}

// GetGroupMemberList
//  @Description: 取群成员列表
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param isRefresh 是否刷新获取:0从缓存获取,1刷新并获取
//  @return []byte
//  @return error
func GetGroupMemberList(robotWxid, groupWxid string, isRefresh int) ([]byte, error) {
	data := &sendGroup{
		Type:      206,
		RobotWxId: robotWxid,
		GroupWxId: groupWxid,
		IsRefresh: isRefresh,
		Key:       key,
	}
	return data.Send()
}

// GetGroupMember
//  @Description: 取群成员资料
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param memberWxid 群员id
//  @return []byte
//  @return error
func GetGroupMember(robotWxid, groupWxid, memberWxid string) ([]byte, error) {
	data := &sendGroup{
		Type:       207,
		RobotWxId:  robotWxid,
		GroupWxId:  groupWxid,
		MemberWxId: memberWxid,
		Key:        key,
	}
	return data.Send()
}

// AcceptTransfer
//  @Description: 接收转账
//  @param robotWxid 机器人id
//  @param friendWxid 好友id
//  @param jsonString 转账事件消息原文
//  @return []byte
//  @return error
func AcceptTransfer(robotWxid, friendWxid, jsonString string) ([]byte, error) {
	data := &sendFriend{
		Type:       301,
		RobotWxId:  robotWxid,
		FriendWxId: friendWxid,
		Msg:        url.QueryEscape(jsonString),
		Key:        key,
	}
	return data.Send()
}

// AgreeGroupInvite
//  @Description: 同意群聊邀请
//  @param robotWxid 机器人id
//  @param jsonString 群聊邀请事件原文
//  @return []byte
//  @return error
func AgreeGroupInvite(robotWxid, jsonString string) ([]byte, error) {
	data := &send{
		Type:      302,
		RobotWxId: robotWxid,
		Msg:       url.QueryEscape(jsonString),
		Key:       key,
	}
	return data.Send()
}

// AgreeFriendVerify
//  @Description: 同意好友请求
//  @param robotWxid 机器人id
//  @param jsonString 好友邀请事件原文
//  @return []byte
//  @return error
func AgreeFriendVerify(robotWxid, jsonString string) ([]byte, error) {
	data := &send{
		Type:      303,
		RobotWxId: robotWxid,
		Msg:       url.QueryEscape(jsonString),
		Key:       key,
	}
	return data.Send()
}

// ModifyFriendNote
//  @Description: 修改好友备注
//  @param robotWxid 机器人id
//  @param friendWxid 好友id
//  @param note 新备注
//  @return []byte
//  @return error
func ModifyFriendNote(robotWxid, friendWxid, note string) ([]byte, error) {
	data := &sendFriend{
		Type:       304,
		RobotWxId:  robotWxid,
		FriendWxId: friendWxid,
		Note:       url.QueryEscape(note),
		Key:        key,
	}
	return data.Send()
}

// DeleteFriend
//  @Description: 删除好友
//  @param robotWxid 机器人id
//  @param friendWxid 好友id
//  @return []byte
//  @return error
func DeleteFriend(robotWxid, friendWxid string) ([]byte, error) {
	data := &sendFriend{
		Type:       305,
		RobotWxId:  robotWxid,
		FriendWxId: friendWxid,
		Key:        key,
	}
	return data.Send()
}

// RemoveGroupMember
//  @Description: 踢出群成员
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param memberWxid 群成员id
//  @return []byte
//  @return error
func RemoveGroupMember(robotWxid, groupWxid, memberWxid string) ([]byte, error) {
	data := &sendGroup{
		Type:       306,
		RobotWxId:  robotWxid,
		GroupWxId:  groupWxid,
		MemberWxId: memberWxid,
		Key:        key,
	}
	return data.Send()
}

// ModifyGroupName
//  @Description: 修改群名
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param groupName 新群名
//  @return []byte
//  @return error
func ModifyGroupName(robotWxid, groupWxid, groupName string) ([]byte, error) {
	data := &sendGroup{
		Type:      307,
		RobotWxId: robotWxid,
		GroupWxId: groupWxid,
		GroupName: url.QueryEscape(groupName),
		Key:       key,
	}
	return data.Send()
}

// ModifyGroupNotice
//  @Description: 修改群公告
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param notice 新公告
//  @return []byte
//  @return error
func ModifyGroupNotice(robotWxid, groupWxid, notice string) ([]byte, error) {
	data := &sendGroup{
		Type:      308,
		RobotWxId: robotWxid,
		GroupWxId: groupWxid,
		Notice:    url.QueryEscape(notice),
		Key:       key,
	}
	return data.Send()
}

// BuildingGroup
//  @Description: 创建新群
//  @param robotWxid 机器人id
//  @param friends 三个人及以上的好友id
//  @return []byte
//  @return error
func BuildingGroup(robotWxid string, friends []string) ([]byte, error) {
	data := &sendGroup{
		Type:      309,
		RobotWxId: robotWxid,
		Friends:   friends,
		Key:       key,
	}
	return data.Send()
}

// QuitGroup
//  @Description: 退出群聊
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @return []byte
//  @return error
func QuitGroup(robotWxid, groupWxid string) ([]byte, error) {
	data := &sendGroup{
		Type:      310,
		RobotWxId: robotWxid,
		GroupWxId: groupWxid,
		Key:       key,
	}
	return data.Send()
}

// InviteInGroup
//  @Description: 邀请加入群聊
//  @param robotWxid 机器人id
//  @param groupWxid 群id
//  @param friendWxid 好友id
//  @return []byte
//  @return error
func InviteInGroup(robotWxid, groupWxid, friendWxid string) ([]byte, error) {
	data := &sendGroup{
		Type:       311,
		RobotWxId:  robotWxid,
		GroupWxId:  groupWxid,
		FriendWxId: friendWxid,
		Key:        key,
	}
	return data.Send()
}

// Send
//  @Description: 发送请求
//  @receiver t
//  @return []byte
//  @return error
func (t *sendGroup) Send() ([]byte, error) {
	params, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return post(sendUrl, string(params), 2)
}

// Send
//  @Description: 发送请求
//  @receiver t
//  @return []byte
//  @return error
func (t *sendFriend) Send() ([]byte, error) {
	params, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return post(sendUrl, string(params), 2)
}

// Send
//  @Description: 发送请求
//  @receiver t
//  @return []byte
//  @return error
func (t *send) Send() ([]byte, error) {
	params, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return post(sendUrl, string(params), 2)
}

// post
//  @Description: 发起post请求
//  @param url 地址
//  @param params 参数
//  @param types 请求content-type:0为json,1为form-urlencoded
//  @return []byte
//  @return error
func post(url, params string, types int) ([]byte, error) {
	client := &http.Client{}
	urls := url
	req, err := http.NewRequest("POST", urls, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	switch types {
	case 1:
		req.Header.Set("Content-Type", "application/json")
	case 2:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
