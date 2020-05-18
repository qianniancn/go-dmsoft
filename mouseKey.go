// 鼠键

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *DmSoft) EnableMouseAccuracy(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseAccuracy", enable)
	return int(ret.Val)
}

func (com *DmSoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DmSoft) GetCursorShape() string {
	ret, _ := com.dm.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *DmSoft) GetCursorShapeEx(types int) string {
	ret, _ := com.dm.CallMethod("GetCursorShapeEx", types)
	return ret.ToString()
}

func (com *DmSoft) GetCursorSpot() string {
	ret, _ := com.dm.CallMethod("GetCursorSpot")
	return ret.ToString()
}

func (com *DmSoft) GetKeyState(vkCode int) int {
	ret, _ := com.dm.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *DmSoft) GetMouseSpeed() int {
	ret, _ := com.dm.CallMethod("GetMouseSpeed")
	return int(ret.Val)
}

func (com *DmSoft) KeyDown(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *DmSoft) KeyDownChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *DmSoft) KeyPress(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *DmSoft) KeyPressChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

// KeyPressStr指定的字符串序列，依次按顺序按下其中的字符
func (com *DmSoft) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.dm.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *DmSoft) KeyUp(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *DmSoft) KeyUpChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *DmSoft) LeftClick() int {
	ret, _ := com.dm.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *DmSoft) LeftDoubleClick() int {
	ret, _ := com.dm.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *DmSoft) LeftDown() int {
	ret, _ := com.dm.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *DmSoft) LeftUp() int {
	ret, _ := com.dm.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *DmSoft) MiddleClick() int {
	ret, _ := com.dm.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *DmSoft) MiddleDown() int {
	ret, _ := com.dm.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *DmSoft) MiddleUp() int {
	ret, _ := com.dm.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *DmSoft) MoveR(rx, ry int) int {
	ret, _ := com.dm.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *DmSoft) MoveTo(x, y int) int {
	ret, _ := com.dm.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *DmSoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.dm.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

func (com *DmSoft) RightClick() int {
	ret, _ := com.dm.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *DmSoft) RightDown() int {
	ret, _ := com.dm.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *DmSoft) RightUp() int {
	ret, _ := com.dm.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *DmSoft) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *DmSoft) SetMouseDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *DmSoft) SetMouseSpeed(speed int) int {
	ret, _ := com.dm.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *DmSoft) SetSimMode(mode int) int {
	ret, _ := com.dm.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

func (com *DmSoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.dm.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

func (com *DmSoft) WheelDown() int {
	ret, _ := com.dm.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *DmSoft) WheelUp() int {
	ret, _ := com.dm.CallMethod("WheelUp")
	return int(ret.Val)
}
