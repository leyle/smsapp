package smsapp

import "github.com/go-redis/redis"

var (
	SMS_RESPONSE_OK = 2
)

// 发送短信的配置文件
type SmsOption struct {
	Phone string
	Account string
	Passwd string
	Url string
	Rclient *redis.Client

	// 以下两者可以为空，如果为空时，就是内部默认的生成
	Default bool // 是否默认发送，默认为 true
	Code string
	CodeLen int
	Content string
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