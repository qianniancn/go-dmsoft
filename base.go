package dmsoft

func (com *dmsoft) EnablePicCache(enable int) int {
	ret, _ := com.dm.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}

func (com *dmsoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *dmsoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod("GetDmCount")
	return int(ret.Val)
}

func (com *dmsoft) GetID() int {
	ret, _ := com.dm.CallMethod("GetID")
	return int(ret.Val)
}

func (com *dmsoft) GetLastError() int {
	ret, _ := com.dm.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *dmsoft) GetPath() string {
	ret, _ := com.dm.CallMethod("GetPath")
	return ret.ToString()
}

func (com *dmsoft) Reg(regCode string, verInfo string) int {
	ret, _ := com.dm.CallMethod("Reg", regCode, verInfo)
	return int(ret.Val)
}

func (com *dmsoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegEx", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegExNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *dmsoft) SetDisplayInput(mode string) int {
	ret, _ := com.dm.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}

func (com *dmsoft) SetEnumWindowDelay(delay int) int {
	ret, _ := com.dm.CallMethod("SetEnumWindowDelay", delay)
	return int(ret.Val)
}

func (com *dmsoft) SetPath(path string) int {
	ret, _ := com.dm.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *dmsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.dm.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *dmsoft) SpeedNormalGraphic(enable int) int {
	ret, _ := com.dm.CallMethod("SpeedNormalGraphic", enable)
	return int(ret.Val)
}

// Ver get version
func (com *dmsoft) Ver() string {
	ver, _ := com.dm.CallMethod("Ver")
	return ver.ToString()
}
