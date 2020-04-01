// 其他

package dmsoft

func (com *dmsoft) DmGuard(enable, lType int) int {
	ret, _ := com.dm.CallMethod("DmGuard", enable, lType)
	return int(ret.Val)
}

func (com *dmsoft) DmGuardParams(cmd, subcmd, param string) string {
	ret, _ := com.dm.CallMethod("DmGuardParams", cmd, subcmd, param)
	return ret.ToString()
}

func (com *dmsoft) UnLoadDriver() int {
	ret, _ := com.dm.CallMethod("UnLoadDriver")
	return int(ret.Val)
}

// long ActiveInputMethod(hwnd,input_method)
// long CheckInputMethod(hwnd,input_method)
// long EnterCri()
// string ExecuteCmd(cmd,current_dir,time_out)
// long FindInputMethod(input_method)
// long InitCri()
// long LeaveCri()
// long ReleaseRef()
// long SetExitThread(enable)
