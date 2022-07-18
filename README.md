# ECHO-CLI

## package

### linux arm

```linux
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -a -o bin/echo-cli_arm
```

### linux indel

```linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/echo-cli_intel
```

## 使用方法

### 执行方法

```linux
chmod +x echo-cli
./echo-cli -h 3.236.42.248 -p 4000
```

### 参数说明

```linux
-h 访问IP地址
-p 访问端口
```
