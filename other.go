// 其他
// 大漠保护盾

package dmsoft

// DmGuard Type取值请查看大漠文档
func (com *DmSoft) DmGuard(enable int, types string) int {
	ret, _ := com.dm.CallMethod("DmGuard", enable, types)
	return int(ret.Val)
}

// DmGuardExtract 释放插件用的驱动. 可以自己拿去签名. 防止有人对我的签名进行检测.
// 强烈推荐使用驱动的用户使用. 仅释放64位系统的驱动.
func (com *DmSoft) DmGuardExtract(types, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardExtract", types, path)
	return int(ret.Val)
}

// 加载用DmGuardExtract释放出的驱动. 建议自己签名后,然后找个自己喜欢的路径加载. 仅支持64位系统的驱动加载.
// 加载成功后,就可以正常调用DmGuard了.
func (com *DmSoft) DmGuardLoadCustom(types, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardLoadCustom", types, path)
	return int(ret.Val)
}

// DmGuardParams DmGuard的加强接口,用于获取一些额外信息.具体信息请查看大漠文档
func (com *DmSoft) DmGuardParams(cmd, subcmd, param string) string {
	ret, _ := com.dm.CallMethod("DmGuardParams", cmd, subcmd, param)
	return ret.ToString()
}

// UnLoadDriver 卸载插件相关的所有驱动. 仅对64位系统的驱动生效
// 注 : 此接口一般不建议使用. 除非你真的知道自己在干什么.
func (com *DmSoft) UnLoadDriver() int {
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
