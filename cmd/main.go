package main

import (
	"github.com/qianniancn/go-dmsoft"
	"log"
	"os"
	"unsafe"
)

const (
	// 填写自己的注册码
	DmRegCode   = ""
	DmExtraCode = ""
)

var dm *dmsoft.Dmsoft

type _bytes struct {
	Data int
	Len  int
}

func main() {
	dm = CreateDmObj()
	log.Printf("插件版本:%s", dm.Ver())
	ret := dm.Reg(DmRegCode, DmExtraCode)
	switch ret {
	case 1:
		log.Println("付费功能注册成功")
	case -1:
		log.Println("无法连接网络")
	case -2:
		log.Println("进程没有以管理员方式运行")
	default:
		log.Println("失败 (未知错误)")
	}
	defer dm.Release()

	var data, size int
	dm.GetScreenDataBmp(0, 0, 800, 800, &data, &size)
	bs := *(*[]byte)(unsafe.Pointer(&_bytes{
		data,
		size,
	}))
	os.WriteFile("test.bmp", bs, os.ModePerm)

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
