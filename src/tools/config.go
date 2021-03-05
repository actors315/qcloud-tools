package tools

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type CertItem struct {
	Domain         string `yaml:"domain"`
	PublicKeyPath  string `yaml:"publicKeyPath"`
	PrivateKeyPath string `yaml:"privateKeyPath"`
	Alias          string `yaml:"alias"`
}

type CvmItem struct {

}

type QcloudConfig struct {
	SecretId       string `yaml:"secretId"`
	SecretKey      string `yaml:"secretKey"`
	Certificates map[string]CertItem `yaml:"certificates"`
}

func (config *QcloudConfig) GetCertItem(group string) string  {

	certItem,found := config.Certificates[group]
	if !found {
		return ""
	}

	publicData, _ := ioutil.ReadFile(certItem.PublicKeyPath)
	privateData, _ := ioutil.ReadFile(certItem.PrivateKeyPath)

	publicKeyData := strings.TrimSpace(string(publicData))
	publicKeyData = strings.ReplaceAll(publicKeyData, "\n", "\\n")

	privateKeyData := strings.TrimSpace(string(privateData))
	privateKeyData = strings.ReplaceAll(privateKeyData, "\n", "\\n")

	params := "{\"Domain\":\"%s\",\"Https\":{\"Switch\":\"on\",\"Http2\":\"on\",\"CertInfo\":{\"Certificate\":\"%s\",\"PrivateKey\":\"%s\",\"Message\":\"%s\"}}}"
	params = fmt.Sprintf(params, certItem.Domain, publicKeyData, privateKeyData, time.Now().Format("2006-01-02"))

	return params
}