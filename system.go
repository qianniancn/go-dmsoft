// 系统

package dmsoft

func (com *dmsoft) Beep(f, duration int) int {
	ret, _ := com.dm.CallMethod("Beep", f, duration)
	return int(ret.Val)
}

func (com *dmsoft) CheckFontSmooth() int {
	ret, _ := com.dm.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *dmsoft) CheckUAC() int {
	ret, _ := com.dm.CallMethod("CheckUAC")
	return int(ret.Val)
}

func (com *dmsoft) Delay(mis int) int {
	ret, _ := com.dm.CallMethod("Delay")
	return int(ret.Val)
}
func (com *dmsoft) Delays(misMin, misMax int) int {
	ret, _ := com.dm.CallMethod("Delays", misMin, misMax)
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *dmsoft) DisableCloseDisplayAndSleep() int {
	ret, _ := com.dm.CallMethod("DisableCloseDisplayAndSleep")
	return int(ret.Val)
}

func (com *dmsoft) DisableFontSmooth() int {
	ret, _ := com.dm.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *dmsoft) DisablePowerSave() int {
	ret, _ := com.dm.CallMethod("DisablePowerSave")
	return int(ret.Val)
}

func (com *dmsoft) DisableScreenSave() int {
	ret, _ := com.dm.CallMethod("DisableScreenSave")
	return int(ret.Val)
}

func (com *dmsoft) EnableFontSmooth() int {
	ret, _ := com.dm.CallMethod("EnableFontSmooth")
	return int(ret.Val)
}

func (com *dmsoft) ExitOs(types int) int {
	ret, _ := com.dm.CallMethod("ExitOs", types)
	return int(ret.Val)
}

func (com *dmsoft) GetClipboard() string {
	ret, _ := com.dm.CallMethod("GetClipboard")
	return ret.ToString()
}

func (com *dmsoft) GetCPUType() int {
	ret, _ := com.dm.CallMethod("GetCpuType")
	return int(ret.Val)
}

func (com *dmsoft) GetCPUUsage() int {
	ret, _ := com.dm.CallMethod("GetCpuUsage")
	return int(ret.Val)
}

func (com *dmsoft) GetDir(types int) string {
	ret, _ := com.dm.CallMethod("GetDir", types)
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *dmsoft) GetDiskModel(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskModel", index)
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *dmsoft) GetDiskReversion(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskReversion", index)
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *dmsoft) GetDiskSerial(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskSerial", index)
	return ret.ToString()
}

func (com *dmsoft) GetDisplayInfo() string {
	ret, _ := com.dm.CallMethod("GetDisplayInfo")
	return ret.ToString()
}

func (com *dmsoft) GetDPI() int {
	ret, _ := com.dm.CallMethod("GetDPI")
	return int(ret.Val)
}

func (com *dmsoft) GetLocale() int {
	ret, _ := com.dm.CallMethod("GetLocale")
	return int(ret.Val)
}

func (com *dmsoft) GetMachineCode() string {
	ret, _ := com.dm.CallMethod("GetMachineCode")
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *dmsoft) GetMachineCodeNoMac() string {
	ret, _ := com.dm.CallMethod("GetMachineCodeNoMac")
	return ret.ToString()
}

func (com *dmsoft) GetMemoryUsage() int {
	ret, _ := com.dm.CallMethod("GetMemoryUsage")
	return int(ret.Val)
}

func (com *dmsoft) GetNetTime() string {
	ret, _ := com.dm.CallMethod("GetNetTime")
	return ret.ToString()
}

func (com *dmsoft) GetNetTimeByIP(ip string) string {
	ret, _ := com.dm.CallMethod("GetNetTimeByIp", ip)
	return ret.ToString()
}

func (com *dmsoft) GetOsBuildNumber() int {
	ret, _ := com.dm.CallMethod("GetOsBuildNumber")
	return int(ret.Val)
}

func (com *dmsoft) GetOsType() int {
	ret, _ := com.dm.CallMethod("GetOsType")
	return int(ret.Val)
}

func (com *dmsoft) GetScreenDepth() int {
	ret, _ := com.dm.CallMethod("GetScreenDepth")
	return int(ret.Val)
}

func (com *dmsoft) GetScreenHeight() int {
	ret, _ := com.dm.CallMethod("GetScreenHeight")
	return int(ret.Val)
}

func (com *dmsoft) GetScreenWidth() int {
	ret, _ := com.dm.CallMethod("GetScreenWidth")
	return int(ret.Val)
}

func (com *dmsoft) GetTime() int {
	ret, _ := com.dm.CallMethod("GetTime")
	return int(ret.Val)
}

func (com *dmsoft) Is64Bit() int {
	ret, _ := com.dm.CallMethod("Is64Bit")
	return int(ret.Val)
}

func (com *dmsoft) IsSurrpotVt() int {
	ret, _ := com.dm.CallMethod("IsSurrpotVt")
	return int(ret.Val)
}

func (com *dmsoft) Play(mediaFile string) int {
	ret, _ := com.dm.CallMethod("Play", mediaFile)
	return int(ret.Val)
}

func (com *dmsoft) RunApp(appPath string, mode int) int {
	ret, _ := com.dm.CallMethod("RunApp", appPath, mode)
	return int(ret.Val)
}

func (com *dmsoft) SetClipboard(value string) int {
	ret, _ := com.dm.CallMethod("SetClipboard", value)
	return int(ret.Val)
}

func (com *dmsoft) SetDisplayAcceler(level int) int {
	ret, _ := com.dm.CallMethod("SetDisplayAcceler", level)
	return int(ret.Val)
}

func (com *dmsoft) SetLocale() int {
	ret, _ := com.dm.CallMethod("SetLocale")
	return int(ret.Val)
}

func (com *dmsoft) SetScreen(width, height, depth int) int {
	ret, _ := com.dm.CallMethod("SetScreen", width, height, depth)
	return int(ret.Val)
}

func (com *dmsoft) SetUAC(enable int) int {
	ret, _ := com.dm.CallMethod("SetUAC", enable)
	return int(ret.Val)
}

func (com *dmsoft) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.dm.CallMethod("ShowTaskBarIcon", hwnd, isShow)
	return int(ret.Val)
}

func (com *dmsoft) Stop(id int) int {
	ret, _ := com.dm.CallMethod("Stop", id)
	return int(ret.Val)
}
