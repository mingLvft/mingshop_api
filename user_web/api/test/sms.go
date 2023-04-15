package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"mingshop_api/user_web/global"
)

func main() {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(global.ServerConfig.AliSmsInfo.AccessKeyId, global.ServerConfig.AliSmsInfo.AccessKeySecret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.SignName = global.ServerConfig.AliSmsInfo.SignName
	request.TemplateCode = global.ServerConfig.AliSmsInfo.TemplateCode
	request.PhoneNumbers = "15328356316"
	request.TemplateParam = "{\"code\":\"1234\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
