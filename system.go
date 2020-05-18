// 系统

package dmsoft

func (com *DmSoft) Beep(f, duration int) int {
	ret, _ := com.dm.CallMethod("Beep", f, duration)
	return int(ret.Val)
}

func (com *DmSoft) CheckFontSmooth() int {
	ret, _ := com.dm.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *DmSoft) CheckUAC() int {
	ret, _ := com.dm.CallMethod("CheckUAC")
	return int(ret.Val)
}

func (com *DmSoft) Delay(mis int) int {
	ret, _ := com.dm.CallMethod("Delay")
	return int(ret.Val)
}
func (com *DmSoft) Delays(misMin, misMax int) int {
	ret, _ := com.dm.CallMethod("Delays", misMin, misMax)
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *DmSoft) DisableCloseDisplayAndSleep() int {
	ret, _ := com.dm.CallMethod("DisableCloseDisplayAndSleep")
	return int(ret.Val)
}

func (com *DmSoft) DisableFontSmooth() int {
	ret, _ := com.dm.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *DmSoft) DisablePowerSave() int {
	ret, _ := com.dm.CallMethod("DisablePowerSave")
	return int(ret.Val)
}

func (com *DmSoft) DisableScreenSave() int {
	ret, _ := com.dm.CallMethod("DisableScreenSave")
	return int(ret.Val)
}

func (com *DmSoft) EnableFontSmooth() int {
	ret, _ := com.dm.CallMethod("EnableFontSmooth")
	return int(ret.Val)
}

func (com *DmSoft) ExitOs(types int) int {
	ret, _ := com.dm.CallMethod("ExitOs", types)
	return int(ret.Val)
}

func (com *DmSoft) GetClipboard() string {
	ret, _ := com.dm.CallMethod("GetClipboard")
	return ret.ToString()
}

func (com *DmSoft) GetCPUType() int {
	ret, _ := com.dm.CallMethod("GetCpuType")
	return int(ret.Val)
}

func (com *DmSoft) GetCPUUsage() int {
	ret, _ := com.dm.CallMethod("GetCpuUsage")
	return int(ret.Val)
}

func (com *DmSoft) GetDir(types int) string {
	ret, _ := com.dm.CallMethod("GetDir", types)
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *DmSoft) GetDiskModel(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskModel", index)
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *DmSoft) GetDiskReversion(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskReversion", index)
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *DmSoft) GetDiskSerial(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskSerial", index)
	return ret.ToString()
}

func (com *DmSoft) GetDisplayInfo() string {
	ret, _ := com.dm.CallMethod("GetDisplayInfo")
	return ret.ToString()
}

func (com *DmSoft) GetDPI() int {
	ret, _ := com.dm.CallMethod("GetDPI")
	return int(ret.Val)
}

func (com *DmSoft) GetLocale() int {
	ret, _ := com.dm.CallMethod("GetLocale")
	return int(ret.Val)
}

func (com *DmSoft) GetMachineCode() string {
	ret, _ := com.dm.CallMethod("GetMachineCode")
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *DmSoft) GetMachineCodeNoMac() string {
	ret, _ := com.dm.CallMethod("GetMachineCodeNoMac")
	return ret.ToString()
}

func (com *DmSoft) GetMemoryUsage() int {
	ret, _ := com.dm.CallMethod("GetMemoryUsage")
	return int(ret.Val)
}

func (com *DmSoft) GetNetTime() string {
	ret, _ := com.dm.CallMethod("GetNetTime")
	return ret.ToString()
}

func (com *DmSoft) GetNetTimeByIP(ip string) string {
	ret, _ := com.dm.CallMethod("GetNetTimeByIp", ip)
	return ret.ToString()
}

func (com *DmSoft) GetOsBuildNumber() int {
	ret, _ := com.dm.CallMethod("GetOsBuildNumber")
	return int(ret.Val)
}

func (com *DmSoft) GetOsType() int {
	ret, _ := com.dm.CallMethod("GetOsType")
	return int(ret.Val)
}

func (com *DmSoft) GetScreenDepth() int {
	ret, _ := com.dm.CallMethod("GetScreenDepth")
	return int(ret.Val)
}

func (com *DmSoft) GetScreenHeight() int {
	ret, _ := com.dm.CallMethod("GetScreenHeight")
	return int(ret.Val)
}

func (com *DmSoft) GetScreenWidth() int {
	ret, _ := com.dm.CallMethod("GetScreenWidth")
	return int(ret.Val)
}

func (com *DmSoft) GetTime() int {
	ret, _ := com.dm.CallMethod("GetTime")
	return int(ret.Val)
}

func (com *DmSoft) Is64Bit() int {
	ret, _ := com.dm.CallMethod("Is64Bit")
	return int(ret.Val)
}

func (com *DmSoft) IsSurrpotVt() int {
	ret, _ := com.dm.CallMethod("IsSurrpotVt")
	return int(ret.Val)
}

func (com *DmSoft) Play(mediaFile string) int {
	ret, _ := com.dm.CallMethod("Play", mediaFile)
	return int(ret.Val)
}

func (com *DmSoft) RunApp(appPath string, mode int) int {
	ret, _ := com.dm.CallMethod("RunApp", appPath, mode)
	return int(ret.Val)
}

func (com *DmSoft) SetClipboard(value string) int {
	ret, _ := com.dm.CallMethod("SetClipboard", value)
	return int(ret.Val)
}

func (com *DmSoft) SetDisplayAcceler(level int) int {
	ret, _ := com.dm.CallMethod("SetDisplayAcceler", level)
	return int(ret.Val)
}

func (com *DmSoft) SetLocale() int {
	ret, _ := com.dm.CallMethod("SetLocale")
	return int(ret.Val)
}

func (com *DmSoft) SetScreen(width, height, depth int) int {
	ret, _ := com.dm.CallMethod("SetScreen", width, height, depth)
	return int(ret.Val)
}

func (com *DmSoft) SetUAC(enable int) int {
	ret, _ := com.dm.CallMethod("SetUAC", enable)
	return int(ret.Val)
}

func (com *DmSoft) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.dm.CallMethod("ShowTaskBarIcon", hwnd, isShow)
	return int(ret.Val)
}

func (com *DmSoft) Stop(id int) int {
	ret, _ := com.dm.CallMethod("Stop", id)
	return int(ret.Val)
}
