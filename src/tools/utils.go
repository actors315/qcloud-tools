package tools

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Utils struct {
	config *QcloudConfig
}

func (u *Utils) GetConfig(file string) *QcloudConfig {

	if u.config != nil {
		return u.config
	}

	if file == "" {
		rootDir, _ := os.Executable()
		rootDir = filepath.Dir(rootDir)
		file = rootDir + "/../config/qcloud.yaml"
	}

	config := new(QcloudConfig)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("failed to read yaml file : %v\n", err)
		panic(err)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	u.config = config

	return u.config

}

func (u *Utils) GetCredential(secretId string,secretKey string) (*common.Credential,*profile.ClientProfile) {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)

	cpf := profile.NewClientProfile()

	return credential,cpf
}