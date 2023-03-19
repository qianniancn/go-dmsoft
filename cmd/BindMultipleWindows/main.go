package main

import (
	"fmt"
	"github.com/qianniancn/go-dmsoft"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	DmRegCode   = ""
	DmExtraCode = ""
)

// 全局大漠对象
var globalDm *dmsoft.Dmsoft

// 多个大漠对象
var dmColl []*dmsoft.Dmsoft

// 句柄集
var handleColl []int

func main() {
	globalDm = CreateDmObj()
	log.Printf("插件版本:%s", globalDm.Ver())
	ret := globalDm.Reg(DmRegCode, DmExtraCode)
	switch ret {
	case 1:
		log.Println("付费功能注册成功")
	case -1:
		log.Fatal("无法连接网络")
	case -2:
		log.Fatal("进程没有以管理员方式运行")
	default:
		log.Fatal("失败 (未知错误)")
	}
	defer globalDm.Release()

	EnumWindowHwnd() // 枚举所有窗口句柄
	BindMultipleWindowsAsync()
}

func BindMultipleWindowsAsync() {
	res := make(chan int) // 创建一个 channel,接收结果

	for i := 0; i < len(handleColl); i++ {
		dm := dmsoft.NewDmsoft() // 创建多个大漠对象
		dmColl = append(dmColl, dm)
		go BindMultipleWindows(dmColl[i], handleColl[i], res) // 启动 goroutine
	}

	for i := 0; i < len(handleColl); i++ {
		v := <-res // 接收 channel 中的结果
		log.Printf("窗口句柄:%d,截图结果:%v", handleColl[i], v != 0)
	}
}

// BindMultipleWindows 绑定多个窗口
func BindMultipleWindows(dm *dmsoft.Dmsoft, handle int, result chan<- int) {
	// 绑定窗口
	ret := dm.BindWindowEx(handle, "dx2", "windows3", "windows", "", 0)
	if ret != 0 {
		log.Printf("绑定窗口成功,句柄:%d", handle)
	}
	// 截图
	r1 := dm.Capture(0, 0, 2000, 2000, "test_"+strconv.Itoa(handle)+".bmp")
	result <- r1 // 将结果发送到 channel

	log.Printf("释放大漠对象:%v,窗口句柄:%d", dm, handle)
	dm.UnBindWindow() // 解除绑定
	dm.Release()
}

// EnumWindowHwnd 枚举出逍遥模拟器的窗口句柄
func EnumWindowHwnd() {
	handleStr := globalDm.EnumWindow(0, "RenderWindowWindow", "Qt5QWindowIcon", 1+2)
	handleStrSlice := strings.Split(handleStr, ",")
	// 将string类型转换为int类型
	for _, str := range handleStrSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		handleColl = append(handleColl, num)
	}
	log.Printf("获取到 %d 个窗口句柄:%v", len(handleColl), handleColl)
}

func CreateDmObj() *dmsoft.Dmsoft {
	// 获取当前工作目录
	dir, _ := os.Getwd()
	// 设置dm.dll路径,并进行注册
	ret := dmsoft.SetDllPathW(dir+"\\dm.dll", 0)
	if ret {
		log.Println("插件注册成功！")
	} else {
		log.Println("插件注册失败！")
	}
	return dmsoft.NewDmsoft()
}
