package smsapp

import "github.com/go-redis/redis"

var (
	SMS_RESPONSE_OK = 2
)

// 发送短信的配置文件
type SmsOption struct {
	Account string
	Passwd  string
	Url     string
	R       *redis.Client
	Debug   bool // 如果 true，就不真的调用发送接口
	Default bool // 是否内部生成 code 来发送
}

type Sms struct {
	Phone string `json:"phone"`
	Code string `json:"code"`
	First int64 `json:"first"`
	Ban bool `json:"ban"`
}

// 短信上发送短信后的返回信息结构
type SendSmsResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	SmsId string `json:"smsid"`
}