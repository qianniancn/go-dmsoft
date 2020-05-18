package dmsoft

func (com *DmSoft) EnablePicCache(enable int) int {
	ret, _ := com.dm.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}

func (com *DmSoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *DmSoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod("GetDmCount")
	return int(ret.Val)
}

func (com *DmSoft) GetID() int {
	ret, _ := com.dm.CallMethod("GetID")
	return int(ret.Val)
}

func (com *DmSoft) GetLastError() int {
	ret, _ := com.dm.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *DmSoft) GetPath() string {
	ret, _ := com.dm.CallMethod("GetPath")
	return ret.ToString()
}

func (com *DmSoft) Reg(regCode string, verInfo string) int {
	ret, _ := com.dm.CallMethod("Reg", regCode, verInfo)
	return int(ret.Val)
}

func (com *DmSoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegEx", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DmSoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegExNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DmSoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DmSoft) SetDisplayInput(mode string) int {
	ret, _ := com.dm.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}

func (com *DmSoft) SetEnumWindowDelay(delay int) int {
	ret, _ := com.dm.CallMethod("SetEnumWindowDelay", delay)
	return int(ret.Val)
}

func (com *DmSoft) SetPath(path string) int {
	ret, _ := com.dm.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *DmSoft) SetShowErrorMsg(show int) int {
	ret, _ := com.dm.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *DmSoft) SpeedNormalGraphic(enable int) int {
	ret, _ := com.dm.CallMethod("SpeedNormalGraphic", enable)
	return int(ret.Val)
}

// Ver get version
func (com *DmSoft) Ver() string {
	ver, _ := com.dm.CallMethod("Ver")
	return ver.ToString()
}
