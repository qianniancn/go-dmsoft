// 鼠键

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *Dmsoft) EnableMouseAccuracy(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseAccuracy", enable)
	return int(ret.Val)
}

func (com *Dmsoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCursorShape() string {
	ret, _ := com.dm.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *Dmsoft) GetCursorShapeEx(types int) string {
	ret, _ := com.dm.CallMethod("GetCursorShapeEx", types)
	return ret.ToString()
}

func (com *Dmsoft) GetCursorSpot() string {
	ret, _ := com.dm.CallMethod("GetCursorSpot")
	return ret.ToString()
}

func (com *Dmsoft) GetKeyState(vkCode int) int {
	ret, _ := com.dm.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) GetMouseSpeed() int {
	ret, _ := com.dm.CallMethod("GetMouseSpeed")
	return int(ret.Val)
}

func (com *Dmsoft) KeyDown(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *Dmsoft) KeyPress(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

// KeyPressStr指定的字符串序列，依次按顺序按下其中的字符
func (com *Dmsoft) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.dm.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *Dmsoft) KeyUp(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *Dmsoft) LeftClick() int {
	ret, _ := com.dm.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *Dmsoft) LeftDoubleClick() int {
	ret, _ := com.dm.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *Dmsoft) LeftDown() int {
	ret, _ := com.dm.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *Dmsoft) LeftUp() int {
	ret, _ := com.dm.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleClick() int {
	ret, _ := com.dm.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleDown() int {
	ret, _ := com.dm.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleUp() int {
	ret, _ := com.dm.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *Dmsoft) MoveR(rx, ry int) int {
	ret, _ := com.dm.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *Dmsoft) MoveTo(x, y int) int {
	ret, _ := com.dm.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *Dmsoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.dm.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

func (com *Dmsoft) RightClick() int {
	ret, _ := com.dm.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *Dmsoft) RightDown() int {
	ret, _ := com.dm.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *Dmsoft) RightUp() int {
	ret, _ := com.dm.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *Dmsoft) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseSpeed(speed int) int {
	ret, _ := com.dm.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *Dmsoft) SetSimMode(mode int) int {
	ret, _ := com.dm.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

func (com *Dmsoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.dm.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

func (com *Dmsoft) WheelDown() int {
	ret, _ := com.dm.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *Dmsoft) WheelUp() int {
	ret, _ := com.dm.CallMethod("WheelUp")
	return int(ret.Val)
}
