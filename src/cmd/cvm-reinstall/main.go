package main

import (
	"flag"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"qcloud-tools/src/tools"
)

func main() {

	var group string
	var file string

	flag.StringVar(&group, "group", "test", "分组")
	flag.StringVar(&file, "config", "E:\\go\\cert-syn-tencent-cloud\\config\\qcloud.yaml", "配置文件地址")
	flag.Parse()

	myUtil := new(tools.Utils)
	config := myUtil.GetConfig(file)
	credential,cpf := myUtil.GetCredential(config.SecretId,config.SecretKey)
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"

	cvmItem := config.GetCvmItem(group)

	client, _ := cvm.NewClient(credential, cvmItem.Region, cpf)

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())

}
