package main

import (
	"fmt"
	"strconv"

	"github.com/qianniancn/go-dmsoft"
)

func main() {
	// 如果是x64编译或运行 需要设置export GOARCH=386
	// 使用免注册方式注册大漠插件，也可以使用命令行注册

	dmsoft.SetDllPathW("./dm.dll", 0)
	dm, _ := dmsoft.New()
	defer dm.Release() //必须释放
	fmt.Printf("当前插件版本:%s\n", dm.Ver())
	if dm.Capture(0, 0, 600, 500, "./test.png") == 1 {
		fmt.Println("截图成功")
	}
	// 打开记事本
	if dm.RunApp("C:\\Windows\\System32\\notepad.exe", 0) == 1 {
		dm.Beep(500, 500) // 蜂鸣
		dm.Delay(2000)
		hwnd1 := dm.FindWindow("", "记事本")
		hwnd2, _ := strconv.Atoi(dm.EnumWindow(hwnd1, "", "Edit", 2))
		fmt.Println(hwnd1, hwnd2)

		for i := 1; i < 10; i++ {
			dm.SendString(hwnd2, "hell world\n")
		}
	}
}
