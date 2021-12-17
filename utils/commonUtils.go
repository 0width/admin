package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetError(errors interface{}, r interface{}) string {
	var errs validator.ValidationErrors
	var ok bool
	if errs, ok = errors.(validator.ValidationErrors); !ok {
		return "请求参数或请求类型有误"
	}
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_err"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("err")
		if errTagText != "" {
			return errTagText
		}
		if errText != "" {
			return errText
		}
		return fieldError.Field() + ":" + fieldError.Tag()
	}
	return ""
}
