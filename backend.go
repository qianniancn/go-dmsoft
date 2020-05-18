// 后台

package Dmsoft

func (com *Dmsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *Dmsoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	return int(ret.Val)
}

func (com *Dmsoft) DownCPU(rate int) int {
	ret, _ := com.dm.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *Dmsoft) EnableBind(enable int) int {
	ret, _ := com.dm.CallMethod("EnableBind", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableFakeActive(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFakeActive", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableIme(enable int) int {
	ret, _ := com.dm.CallMethod("EnableIme", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadMsg", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadPatch(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadPatch", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadSync", enable, timeOut)
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseMsg", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableMouseSync", enable, timeOut)
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealKeypad(enable int) int {
	ret, _ := com.dm.CallMethod("EnableRealKeypad", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.dm.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *Dmsoft) EnableSpeedDx(enable int) int {
	ret, _ := com.dm.CallMethod("EnableSpeedDx", enable)
	return int(ret.Val)
}

func (com *Dmsoft) ForceUnBindWindow() int {
	ret, _ := com.dm.CallMethod("ForceUnBindWindow")
	return int(ret.Val)
}

func (com *Dmsoft) GetBindWindow() int {
	ret, _ := com.dm.CallMethod("GetBindWindow")
	return int(ret.Val)
}

func (com *Dmsoft) GetFps() int {
	ret, _ := com.dm.CallMethod("GetFps")
	return int(ret.Val)
}

func (com *Dmsoft) HackSpeed(rate int) int {
	ret, _ := com.dm.CallMethod("HackSpeed", rate)
	return int(ret.Val)
}

func (com *Dmsoft) IsBind(hwnd int) int {
	ret, _ := com.dm.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) LockDisplay(lock int) int {
	ret, _ := com.dm.CallMethod("LockDisplay", lock)
	return int(ret.Val)
}

func (com *Dmsoft) LockInput(lock int) int {
	ret, _ := com.dm.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *Dmsoft) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod("LockMouseRect", x1, y1, x2, y2)
	return int(ret.Val)
}

func (com *Dmsoft) SetAero(enable int) int {
	ret, _ := com.dm.CallMethod("SetAero", enable)
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayDelay", time)
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayRefreshDelay", time)
	return int(ret.Val)
}

func (com *Dmsoft) SwitchBindWindow(hwnd int) int {
	ret, _ := com.dm.CallMethod("SwitchBindWindow", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) UnBindWindow() int {
	ret, _ := com.dm.CallMethod("UnBindWindow")
	return int(ret.Val)
}
