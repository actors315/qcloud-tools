package main

import (
	"flag"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn/v20191012"
	"qcloud-tools/src/tools"
)

func main() {

	var group string
	var file string

	flag.StringVar(&group, "group", "profile", "分组")
	flag.StringVar(&file, "config", "", "配置文件地址")
	flag.Parse()

	myUtil := new(tools.Utils)
	config := myUtil.GetConfig(file)

	credential,cpf := myUtil.GetCredential(config.SecretId,config.SecretKey)
	cpf.HttpProfile.Endpoint = "ecdn.tencentcloudapi.com"

	params := config.GetCertItem(group)

	client, _ := cdn.NewClient(credential, "", cpf)

	request := cdn.NewUpdateDomainConfigRequest()

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.UpdateDomainConfig(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", response.ToJsonString())

}
