// 后台

package dmsoft

func (com *DmSoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *DmSoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	return int(ret.Val)
}

func (com *DmSoft) DownCPU(rate int) int {
	ret, _ := com.dm.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *DmSoft) EnableBind(enable int) int {
	ret, _ := com.dm.CallMethod("EnableBind", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableFakeActive(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFakeActive", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableIme(enable int) int {
	ret, _ := com.dm.CallMethod("EnableIme", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableKeypadMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadMsg", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableKeypadPatch(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadPatch", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadSync", enable, timeOut)
	return int(ret.Val)
}

func (com *DmSoft) EnableMouseMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseMsg", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableMouseSync", enable, timeOut)
	return int(ret.Val)
}

func (com *DmSoft) EnableRealKeypad(enable int) int {
	ret, _ := com.dm.CallMethod("EnableRealKeypad", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.dm.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *DmSoft) EnableSpeedDx(enable int) int {
	ret, _ := com.dm.CallMethod("EnableSpeedDx", enable)
	return int(ret.Val)
}

func (com *DmSoft) ForceUnBindWindow() int {
	ret, _ := com.dm.CallMethod("ForceUnBindWindow")
	return int(ret.Val)
}

func (com *DmSoft) GetBindWindow() int {
	ret, _ := com.dm.CallMethod("GetBindWindow")
	return int(ret.Val)
}

func (com *DmSoft) GetFps() int {
	ret, _ := com.dm.CallMethod("GetFps")
	return int(ret.Val)
}

func (com *DmSoft) HackSpeed(rate int) int {
	ret, _ := com.dm.CallMethod("HackSpeed", rate)
	return int(ret.Val)
}

func (com *DmSoft) IsBind(hwnd int) int {
	ret, _ := com.dm.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *DmSoft) LockDisplay(lock int) int {
	ret, _ := com.dm.CallMethod("LockDisplay", lock)
	return int(ret.Val)
}

func (com *DmSoft) LockInput(lock int) int {
	ret, _ := com.dm.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *DmSoft) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod("LockMouseRect", x1, y1, x2, y2)
	return int(ret.Val)
}

func (com *DmSoft) SetAero(enable int) int {
	ret, _ := com.dm.CallMethod("SetAero", enable)
	return int(ret.Val)
}

func (com *DmSoft) SetDisplayDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayDelay", time)
	return int(ret.Val)
}

func (com *DmSoft) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayRefreshDelay", time)
	return int(ret.Val)
}

func (com *DmSoft) SwitchBindWindow(hwnd int) int {
	ret, _ := com.dm.CallMethod("SwitchBindWindow", hwnd)
	return int(ret.Val)
}

func (com *DmSoft) UnBindWindow() int {
	ret, _ := com.dm.CallMethod("UnBindWindow")
	return int(ret.Val)
}

// v7.2144
// 1. 增加了接口SetInputDm
