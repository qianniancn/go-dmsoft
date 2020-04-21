# go-dmsoft
使用go调用大漠插件,使用[例子](https://github.com/qianniancn/go-dmsoft/tree/master/examples/windows/main.go) 包里包含了最新版本的大漠插件dll，然后就是大漠是付费插件

## 安装
获取这个包的方法就不说了，大家都会。使用了mod

## 注意事项，关于windows64报错的问题
由于大漠插件是win32的Dll，所以在windows64位运行和编译的时候需要设置环境变量。

### 打开终端 Bash
执行 `go env -w GOARCH=386` 或者 `export GOARCH=386`

#### Windows Cmd
执行 `set GOARCH=386`
