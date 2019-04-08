package smsapp

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 检查指定手机号是否可以继续下发短信验证码
func CanSend(phone string) (bool, error) {
	return true, nil
}

func GenerateSmsCode(codeLen int) string {
	rand.Seed(time.Now().UnixNano())
	var code string
	for i := 0; i < codeLen; i++ {
		c := rand.Intn(10)
		code += strconv.Itoa(c)
	}

	return code
}


// 发送短信验证码
func SendSMS(phone, account, password, sendUrl, code, content string)  error {
	v := url.Values{}
	now := strconv.FormatInt(time.Now().Unix(), 10)

	v.Set("account", account)
	v.Set("password", password)
	v.Set("mobile", phone)
	v.Set("content", content)
	v.Set("time", now)
	v.Set("format", "json")

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, sendUrl, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("给手机号 ", phone, "发送短信失败 ", err.Error())
	}

	defer resp.Body.Close()

	fmt.Println("给手机号: ", phone, "发送的验证码是: ", code)

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取发送短信验证码的返回消息失败 ", err.Error())
		return err
	}

	var smsRet SendSmsResponse
	err = jsoniter.Unmarshal(ret, &smsRet)
	if err != nil {
		fmt.Println("反序列化短信发送商返回的数据时失败 ", err.Error())
	}

	if smsRet.Code != SMS_RESPONSE_OK {
		fmt.Println("发送短信失败 ", smsRet.Code, smsRet.Msg)
		return fmt.Errorf("发送短信失败 %s", smsRet.Msg)
	}

	fmt.Println("发送短信验证码，返回的消息是 ", smsRet.Code, smsRet.Msg, smsRet.SmsId)

	return nil
}