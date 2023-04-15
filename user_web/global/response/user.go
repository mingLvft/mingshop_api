package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	NikeName string `json:"name"`
	//sting类型的时间戳转换成时间格式（自定义格式）
	//Birthday string `json:"birthday"`
	//time类型的时间戳转换成时间格式-两种均可
	//Birthday time.Time `json:"birthday"`
	//自定义类型的时间戳转换成时间格式（自定义格式）
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}
