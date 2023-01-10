package dmsoft

func (com *Dmsoft) EnablePicCache(enable int) int {
	ret, _ := com.dm.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}

func (com *Dmsoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *Dmsoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod("GetDmCount")
	return int(ret.Val)
}

func (com *Dmsoft) GetID() int {
	ret, _ := com.dm.CallMethod("GetID")
	return int(ret.Val)
}

func (com *Dmsoft) GetLastError() int {
	ret, _ := com.dm.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *Dmsoft) GetPath() string {
	ret, _ := com.dm.CallMethod("GetPath")
	return ret.ToString()
}

func (com *Dmsoft) Reg(regCode string, verInfo string) int {
	ret, _ := com.dm.CallMethod("Reg", regCode, verInfo)
	return int(ret.Val)
}

func (com *Dmsoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegEx", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *Dmsoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegExNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *Dmsoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayInput(mode string) int {
	ret, _ := com.dm.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}

func (com *Dmsoft) SetEnumWindowDelay(delay int) int {
	ret, _ := com.dm.CallMethod("SetEnumWindowDelay", delay)
	return int(ret.Val)
}

func (com *Dmsoft) SetPath(path string) int {
	ret, _ := com.dm.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *Dmsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.dm.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *Dmsoft) SpeedNormalGraphic(enable int) int {
	ret, _ := com.dm.CallMethod("SpeedNormalGraphic", enable)
	return int(ret.Val)
}

// Ver get version
func (com *Dmsoft) Ver() string {
	ver, _ := com.dm.CallMethod("Ver")
	return ver.ToString()
}
