package cvmVo

import (
	"fmt"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ecdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn/v20191012"
	"qcloud-tools/src/tools"
)

type CredentialItem struct {
	File string
	Group string
}

func (item *CredentialItem) updateCdnCredential(credential *common.Credential,cpf *profile.ClientProfile) {
	client, _ := cdn.NewClient(credential, "", cpf)
	request := cdn.NewUpdateDomainConfigRequest()

	config := tools.NewQcloudConfig(item.File)

	params := config.GetCertRequestParam(item.Group)

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

func (item *CredentialItem) updateEcdnCredential(credential *common.Credential, cpf *profile.ClientProfile) {
	client, _ := ecdn.NewClient(credential, "", cpf)

	request := ecdn.NewUpdateDomainConfigRequest()

	config := tools.NewQcloudConfig(item.File)

	params := config.GetCertRequestParam(item.Group)

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

func (item *CredentialItem) UpdateCredential() {
	config := tools.NewQcloudConfig(item.File)

	certItem, found := config.Certificates[item.Group]
	if !found {
		fmt.Print("配置不存在")
		return
	}

	credential,cpf := tools.GetCredential(config.SecretId,config.SecretKey)

	switch certItem.Product {
		case "ecdn":
			cpf.HttpProfile.Endpoint = "ecdn.tencentcloudapi.com"
			item.updateEcdnCredential(credential,cpf)
		default:
			cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
			item.updateCdnCredential(credential,cpf)
	}

}
