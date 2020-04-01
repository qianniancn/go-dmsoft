// 汇编

package dmsoft

func (com *dmsoft) AsmAdd(asmIns string) int {
	ret, _ := com.dm.CallMethod("AsmAdd", asmIns)
	return int(ret.Val)
}

func (com *dmsoft) AsmCall(hwnd, mode int) int {
	ret, _ := com.dm.CallMethod("AsmCall", hwnd, mode)
	return int(ret.Val)
}

func (com *dmsoft) AsmCallEx(hwnd, mode int, baseAddr string) int64 {
	ret, _ := com.dm.CallMethod("AsmCallEx", hwnd, mode, baseAddr)
	return ret.Val
}

func (com *dmsoft) AsmClear() int {
	ret, _ := com.dm.CallMethod("AsmClear")
	return int(ret.Val)
}

func (com *dmsoft) AsmSetTimeout(timeOut, param int) int {
	ret, _ := com.dm.CallMethod("AsmSetTimeout", timeOut, param)
	return int(ret.Val)
}

func (com *dmsoft) Assemble(timeOut int64, is64bit int) string {
	ret, _ := com.dm.CallMethod("Assemble", timeOut, is64bit)
	return ret.ToString()
}

func (com *dmsoft) DisAssemble(asmCode string, baseAddr int64, is64bit int) string {
	ret, _ := com.dm.CallMethod("DisAssemble", asmCode, baseAddr, is64bit)
	return ret.ToString()
}
