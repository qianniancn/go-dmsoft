package dmsoft

import "github.com/go-ole/go-ole/oleutil"

func (com *dmsoft) EnablePicCache(enable int) int {
	ret, _ := oleutil.CallMethod(com.dm, "EnablePicCache", enable)
	return int(ret.Val)
}

func (com *dmsoft) GetBasePath() string {
	ret, _ := oleutil.CallMethod(com.dm, "GetBasePath")
	return ret.ToString()
}

func (com *dmsoft) GetDmCount() int {
	ret, _ := oleutil.CallMethod(com.dm, "GetDmCount")
	return int(ret.Val)
}

func (com *dmsoft) GetID() int {
	ret, _ := oleutil.CallMethod(com.dm, "GetID")
	return int(ret.Val)
}

func (com *dmsoft) GetLastError() int {
	ret, _ := oleutil.CallMethod(com.dm, "GetLastError")
	return int(ret.Val)
}

func (com *dmsoft) GetPath() string {
	ret, _ := oleutil.CallMethod(com.dm, "GetPath")
	return ret.ToString()
}

func (com *dmsoft) Reg(regCode string, verInfo string) int {
	ret, _ := oleutil.CallMethod(com.dm, "Reg", regCode, verInfo)
	return int(ret.Val)
}

func (com *dmsoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := oleutil.CallMethod(com.dm, "RegEx", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := oleutil.CallMethod(com.dm, "RegExNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := oleutil.CallMethod(com.dm, "RegNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) SetDisplayInput(mode string) int {
	ret, _ := oleutil.CallMethod(com.dm, "SetDisplayInput", mode)
	return int(ret.Val)
}

func (com *dmsoft) SetEnumWindowDelay(delay int) int {
	ret, _ := oleutil.CallMethod(com.dm, "SetEnumWindowDelay", delay)
	return int(ret.Val)
}

func (com *dmsoft) SetPath(path string) int {
	ret, _ := oleutil.CallMethod(com.dm, "SetPath", path)
	return int(ret.Val)
}

func (com *dmsoft) SetShowErrorMsg(show int) int {
	ret, _ := oleutil.CallMethod(com.dm, "SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *dmsoft) SpeedNormalGraphic(enable int) int {
	ret, _ := oleutil.CallMethod(com.dm, "SpeedNormalGraphic", enable)
	return int(ret.Val)
}

func (com *dmsoft) Ver() string {
	ver, _ := oleutil.CallMethod(com.dm, "Ver")
	return ver.ToString()
}
