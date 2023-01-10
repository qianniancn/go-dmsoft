// 系统

package dmsoft

func (com *Dmsoft) Beep(f, duration int) int {
	ret, _ := com.dm.CallMethod("Beep", f, duration)
	return int(ret.Val)
}

func (com *Dmsoft) CheckFontSmooth() int {
	ret, _ := com.dm.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *Dmsoft) CheckUAC() int {
	ret, _ := com.dm.CallMethod("CheckUAC")
	return int(ret.Val)
}

// Delay 函数不推荐使用
// 建议time.Sleep()
func (com *Dmsoft) Delay(mis int) int {
	ret, _ := com.dm.CallMethod("Delay")
	return int(ret.Val)
}
func (com *Dmsoft) Delays(misMin, misMax int) int {
	ret, _ := com.dm.CallMethod("Delays", misMin, misMax)
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *Dmsoft) DisableCloseDisplayAndSleep() int {
	ret, _ := com.dm.CallMethod("DisableCloseDisplayAndSleep")
	return int(ret.Val)
}

func (com *Dmsoft) DisableFontSmooth() int {
	ret, _ := com.dm.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *Dmsoft) DisablePowerSave() int {
	ret, _ := com.dm.CallMethod("DisablePowerSave")
	return int(ret.Val)
}

func (com *Dmsoft) DisableScreenSave() int {
	ret, _ := com.dm.CallMethod("DisableScreenSave")
	return int(ret.Val)
}

func (com *Dmsoft) EnableFontSmooth() int {
	ret, _ := com.dm.CallMethod("EnableFontSmooth")
	return int(ret.Val)
}

func (com *Dmsoft) ExitOs(types int) int {
	ret, _ := com.dm.CallMethod("ExitOs", types)
	return int(ret.Val)
}

func (com *Dmsoft) GetClipboard() string {
	ret, _ := com.dm.CallMethod("GetClipboard")
	return ret.ToString()
}

func (com *Dmsoft) GetCPUType() int {
	ret, _ := com.dm.CallMethod("GetCpuType")
	return int(ret.Val)
}

func (com *Dmsoft) GetCPUUsage() int {
	ret, _ := com.dm.CallMethod("GetCpuUsage")
	return int(ret.Val)
}

func (com *Dmsoft) GetDir(types int) string {
	ret, _ := com.dm.CallMethod("GetDir", types)
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *Dmsoft) GetDiskModel(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskModel", index)
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *Dmsoft) GetDiskReversion(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskReversion", index)
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *Dmsoft) GetDiskSerial(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskSerial", index)
	return ret.ToString()
}

func (com *Dmsoft) GetDisplayInfo() string {
	ret, _ := com.dm.CallMethod("GetDisplayInfo")
	return ret.ToString()
}

func (com *Dmsoft) GetDPI() int {
	ret, _ := com.dm.CallMethod("GetDPI")
	return int(ret.Val)
}

func (com *Dmsoft) GetLocale() int {
	ret, _ := com.dm.CallMethod("GetLocale")
	return int(ret.Val)
}

func (com *Dmsoft) GetMachineCode() string {
	ret, _ := com.dm.CallMethod("GetMachineCode")
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *Dmsoft) GetMachineCodeNoMac() string {
	ret, _ := com.dm.CallMethod("GetMachineCodeNoMac")
	return ret.ToString()
}

func (com *Dmsoft) GetMemoryUsage() int {
	ret, _ := com.dm.CallMethod("GetMemoryUsage")
	return int(ret.Val)
}

func (com *Dmsoft) GetNetTime() string {
	ret, _ := com.dm.CallMethod("GetNetTime")
	return ret.ToString()
}

func (com *Dmsoft) GetNetTimeByIP(ip string) string {
	ret, _ := com.dm.CallMethod("GetNetTimeByIp", ip)
	return ret.ToString()
}

func (com *Dmsoft) GetOsBuildNumber() int {
	ret, _ := com.dm.CallMethod("GetOsBuildNumber")
	return int(ret.Val)
}

func (com *Dmsoft) GetOsType() int {
	ret, _ := com.dm.CallMethod("GetOsType")
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenDepth() int {
	ret, _ := com.dm.CallMethod("GetScreenDepth")
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenHeight() int {
	ret, _ := com.dm.CallMethod("GetScreenHeight")
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenWidth() int {
	ret, _ := com.dm.CallMethod("GetScreenWidth")
	return int(ret.Val)
}

func (com *Dmsoft) GetTime() int {
	ret, _ := com.dm.CallMethod("GetTime")
	return int(ret.Val)
}

func (com *Dmsoft) Is64Bit() int {
	ret, _ := com.dm.CallMethod("Is64Bit")
	return int(ret.Val)
}

func (com *Dmsoft) IsSurrpotVt() int {
	ret, _ := com.dm.CallMethod("IsSurrpotVt")
	return int(ret.Val)
}

func (com *Dmsoft) Play(mediaFile string) int {
	ret, _ := com.dm.CallMethod("Play", mediaFile)
	return int(ret.Val)
}

func (com *Dmsoft) RunApp(appPath string, mode int) int {
	ret, _ := com.dm.CallMethod("RunApp", appPath, mode)
	return int(ret.Val)
}

func (com *Dmsoft) SetClipboard(value string) int {
	ret, _ := com.dm.CallMethod("SetClipboard", value)
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayAcceler(level int) int {
	ret, _ := com.dm.CallMethod("SetDisplayAcceler", level)
	return int(ret.Val)
}

func (com *Dmsoft) SetLocale() int {
	ret, _ := com.dm.CallMethod("SetLocale")
	return int(ret.Val)
}

func (com *Dmsoft) SetScreen(width, height, depth int) int {
	ret, _ := com.dm.CallMethod("SetScreen", width, height, depth)
	return int(ret.Val)
}

func (com *Dmsoft) SetUAC(enable int) int {
	ret, _ := com.dm.CallMethod("SetUAC", enable)
	return int(ret.Val)
}

func (com *Dmsoft) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.dm.CallMethod("ShowTaskBarIcon", hwnd, isShow)
	return int(ret.Val)
}

func (com *Dmsoft) Stop(id int) int {
	ret, _ := com.dm.CallMethod("Stop", id)
	return int(ret.Val)
}
