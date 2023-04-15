package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断手机号是否合法
	ok, _ := regexp.MatchString("^1[3|4|5|7|8][0-9]{9}$", mobile)
	if !ok {
		return false
	}
	return true
}
