package smsapp

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// 检查指定手机号是否可以继续下发短信验证码
func CanSend(phone string) (bool, error) {
	if len(phone) != PHONE_LEN {
		return false, errors.New("手机号长度不对")
	}

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