# 腾讯云小工具

目前我的项目都是用 letsencrypt 来生成免费证书，同时用了腾讯云 cdn 加速。letsencrypt 证书有效期只有三个月，每次生成完要手动更新也挺烦，于是就有了这个工具   

## 配置项

config 目录下,重建 qcloud.simple.yaml 为 qcloud.yaml

修改对应配置

```
secretId: secretId
secretKey: secretKey

certificates:
  test:
    domain: "www.test.com"
    publicKeyPath: "/usr/local/public.cer"
    privateKeyPath: "/usr/local/private.pem"
    alias: test
```

默认可执行文件是在 bin 目录，配置文件在 config 目录。否则执行时需要手动指定配置文件路径 `--config={dir}/qcloud.yaml`

## 同步证书到 cdn

```
# 编译可执行文件
make cert-sync-cdn

# 或者也可以直接 build
# go build -o bin/cert-sync-cdn src/cmd/certificate-sync/cdn.go

# 同步
{DIR}/bin/cert-sync-cdn --group=test

```

### acme.sh

```
acme.sh --issue -d example.com --renew-hook "{DIR}/bin/cert-sync-cdn --group=test" # test对应配置的分组
```

注意替换可执行文件完整路径

## 同步证书到 ecdn

```
# 编译可执行文件
make cert-sync-ecdn

# 同步
{DIR}/bin/cert-sync-ecdn --group=wildcard

```

## 重建 cvm

以过期的 cvm 中创建镜像，以该镜像重建新机器