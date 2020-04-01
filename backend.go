// 后台

package dmsoft

func (com *dmsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *dmsoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	return int(ret.Val)
}

func (com *dmsoft) DownCPU(rate int) int {
	ret, _ := com.dm.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *dmsoft) EnableBind(enable int) int {
	ret, _ := com.dm.CallMethod("EnableBind", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableFakeActive(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFakeActive", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableIme(enable int) int {
	ret, _ := com.dm.CallMethod("EnableIme", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableKeypadMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadMsg", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableKeypadPatch(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadPatch", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadSync", enable, timeOut)
	return int(ret.Val)
}

func (com *dmsoft) EnableMouseMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseMsg", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableMouseSync", enable, timeOut)
	return int(ret.Val)
}

func (com *dmsoft) EnableRealKeypad(enable int) int {
	ret, _ := com.dm.CallMethod("EnableRealKeypad", enable)
	return int(ret.Val)
}

func (com *dmsoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.dm.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *dmsoft) EnableSpeedDx(enable int) int {
	ret, _ := com.dm.CallMethod("EnableSpeedDx", enable)
	return int(ret.Val)
}

func (com *dmsoft) ForceUnBindWindow() int {
	ret, _ := com.dm.CallMethod("ForceUnBindWindow")
	return int(ret.Val)
}

func (com *dmsoft) GetBindWindow() int {
	ret, _ := com.dm.CallMethod("GetBindWindow")
	return int(ret.Val)
}

func (com *dmsoft) GetFps() int {
	ret, _ := com.dm.CallMethod("GetFps")
	return int(ret.Val)
}

func (com *dmsoft) HackSpeed(rate int) int {
	ret, _ := com.dm.CallMethod("HackSpeed", rate)
	return int(ret.Val)
}

func (com *dmsoft) IsBind(hwnd int) int {
	ret, _ := com.dm.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *dmsoft) LockDisplay(lock int) int {
	ret, _ := com.dm.CallMethod("LockDisplay", lock)
	return int(ret.Val)
}

func (com *dmsoft) LockInput(lock int) int {
	ret, _ := com.dm.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *dmsoft) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod("LockMouseRect", x1, y1, x2, y2)
	return int(ret.Val)
}

func (com *dmsoft) SetAero(enable int) int {
	ret, _ := com.dm.CallMethod("SetAero", enable)
	return int(ret.Val)
}

func (com *dmsoft) SetDisplayDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayDelay", time)
	return int(ret.Val)
}

func (com *dmsoft) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayRefreshDelay", time)
	return int(ret.Val)
}

func (com *dmsoft) SwitchBindWindow(hwnd int) int {
	ret, _ := com.dm.CallMethod("SwitchBindWindow", hwnd)
	return int(ret.Val)
}

func (com *dmsoft) UnBindWindow() int {
	ret, _ := com.dm.CallMethod("UnBindWindow")
	return int(ret.Val)
}
