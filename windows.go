// 窗口

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *DmSoft) ClientToScreen(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod("ClientToScreen", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DmSoft) EnumProcess(name string) string {
	ret, _ := com.dm.CallMethod("EnumProcess", name)
	return ret.ToString()
}

func (com *DmSoft) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

func (com *DmSoft) EnumWindowByProcess(processName, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindowByProcess", processName, title, className, filter)
	return ret.ToString()
}

func (com *DmSoft) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindowByProcessId", pid, title, className, filter)
	return ret.ToString()
}

func (com *DmSoft) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.dm.CallMethod("EnumWindowSuper", spec1, flag1, type1, spec2, flag2, type2, sort)
	return ret.ToString()
}

func (com *DmSoft) FindWindow(class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

func (com *DmSoft) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

func (com *DmSoft) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowByProcessId", processId, class, title)
	return int(ret.Val)
}

func (com *DmSoft) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

func (com *DmSoft) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.dm.CallMethod("FindWindowSuper", spec1, flag1, type1, spec2, flag2, type2)
	return int(ret.Val)
}

func (com *DmSoft) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.dm.CallMethod("GetClientRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	// 清除对象变量内存
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *DmSoft) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.dm.CallMethod("GetClientSize", hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	pWidth.Clear()
	pHeight.Clear()
	return int(ret.Val)
}

func (com *DmSoft) GetForegroundFocus() int {
	ret, _ := com.dm.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

func (com *DmSoft) GetForegroundWindow() int {
	ret, _ := com.dm.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

func (com *DmSoft) GetMousePointWindow() int {
	ret, _ := com.dm.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *DmSoft) GetPointWindow(x, y int) int {
	ret, _ := com.dm.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *DmSoft) GetProcessInfo(pid int) string {
	ret, _ := com.dm.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

func (com *DmSoft) GetSpecialWindow(flag int) int {
	ret, _ := com.dm.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

func (com *DmSoft) GetWindow(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

func (com *DmSoft) GetWindowClass(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *DmSoft) GetWindowProcessId(hwnd int) int {
	ret, _ := com.dm.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *DmSoft) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *DmSoft) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.dm.CallMethod("GetWindowRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *DmSoft) GetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *DmSoft) GetWindowTitle(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

func (com *DmSoft) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.dm.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *DmSoft) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod("ScreenToClient", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DmSoft) SendPaste(hwnd int) int {
	ret, _ := com.dm.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

func (com *DmSoft) SendString(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod("SendString", hwnd, str)
	return int(ret.Val)
}

func (com *DmSoft) SendString2(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod("SendString2", hwnd, str)
	return int(ret.Val)
}

func (com *DmSoft) SendStringIme(str string) int {
	ret, _ := com.dm.CallMethod("SendStringIme", str)
	return int(ret.Val)
}

func (com *DmSoft) SendStringIme2(hwnd int, str string, mode int) int {
	ret, _ := com.dm.CallMethod("SendStringIme2", hwnd, str, mode)
	return int(ret.Val)
}

func (com *DmSoft) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *DmSoft) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *DmSoft) SetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *DmSoft) SetWindowText(hwnd int, title string) int {
	ret, _ := com.dm.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *DmSoft) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.dm.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
