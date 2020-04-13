package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

const AliAppKey = "网上申请"
const AliAppSecret = "网上申请"

func main() {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", AliAppKey, AliAppSecret)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = "137XXXX9608"
	request.QueryParams["SignName"] = "XXXX"
	request.QueryParams["TemplateCode"] = "SMS_Code"
	request.QueryParams["TemplateParam"] = "{user:\"felix\",senior:\"felix-senior\",relation:\"dd\"}"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())
}
