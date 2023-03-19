# go-dmsoft

大漠是付费插件，不充值无法使用

## 安装

`go get -u github.com/qianniancn/go-dmsoft`

## 注意事项

由于大漠插件是 win32 的 Dll，所以在 windows64 位运行和编译的时候需要设置环境变量。

### 打开终端 Bash

执行 `go env -w GOARCH=386` 或者 `export GOARCH=386`

### Windows Cmd

执行 `set GOARCH=386`

### Goland 设置环境

![image](./docs/2023-01-10-172752.png)

## 工具下载

大漠插件(当前最新版本 7.2302) [下载地址](http://121.204.253.175:8088/file/dm.rar)

绑定测试工具(v70) [下载地址](http://121.204.253.175:8088/file/%E7%BB%91%E5%AE%9A%E5%B7%A5%E5%85%B7.rar)

免注册 DLL(v11) [下载地址](http://121.204.253.175:8088/file/%E5%85%8D%E6%B3%A8%E5%86%8C.rar)

压缩包解压密码是 1234. 当前版本号 7.2302
