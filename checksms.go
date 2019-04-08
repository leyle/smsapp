package smsapp

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func (s *SmsOption) CheckSms() (bool, error) {
	dbcode, err := s.Rclient.Get(s.Phone).Result()
	if err != nil && err != redis.Nil {
		emsg := fmt.Sprintf("从 redis 读取手机号[%s]对应的 code 失败, %s", s.Phone, err.Error())
		fmt.Println(emsg)
		return false, errors.New(emsg)
	}

	if dbcode != s.Code {
		return false, nil
	}

	return true, nil
}
