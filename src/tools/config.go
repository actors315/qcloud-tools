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
	Region string `yaml:"region"`
}

type QcloudConfig struct {
	SecretId     string              `yaml:"secretId"`
	SecretKey    string              `yaml:"secretKey"`
	Certificates map[string]CertItem `yaml:"certificates"`
	Cvms         map[string]CvmItem  `yaml:"cvms"`
}

func (config *QcloudConfig) GetCvmItem(group string) CvmItem {
	cvmItem,found := config.Cvms[group]
	if !found {
		panic("配置不存在")
	}
	return cvmItem
}

func (config *QcloudConfig) GetCertItem(group string) CertItem {
	certItem,found := config.Certificates[group]
	if !found {
		panic("配置不存在")
	}
	return certItem
}

func (config *QcloudConfig) GetCertRequestParam(group string) string  {

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