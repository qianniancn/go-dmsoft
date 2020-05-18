// 图色

package Dmsoft

import ole "github.com/go-ole/go-ole"

func (com *Dmsoft) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.dm.CallMethod("AppendPicAddr", picInfo, addr, size)
	return ret.ToString()
}

func (com *Dmsoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *Dmsoft) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.dm.CallMethod("CaptureGif", x1, y1, x2, y2, file, delay, time)
	return int(ret.Val)
}

func (com *Dmsoft) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("CaptureJpg", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *Dmsoft) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("CapturePng", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *Dmsoft) CapturePre(file string) int {
	ret, _ := com.dm.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *Dmsoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.dm.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

func (com *Dmsoft) EnableFindPicMultithread(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFindPicMultithread", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.dm.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

func (com *Dmsoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorBlock(x1, y1, x2, y2 int, color string, sim float32, count, width, height int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColorBlock", x1, y1, x2, y2, color, sim, count, width, height, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float32, count, width, height int) string {
	ret, _ := com.dm.CallMethod("FindColorBlockEx", x1, y1, x2, y2, color, sim, count, width, height)
	return ret.ToString()
}

func (com *Dmsoft) FindColorE(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorE", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindMulColor(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("FindMulColor", x1, y1, x2, y2, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColorE(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorE", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicE(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicE", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicExS", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

// func (com *Dmsoft)FindPicMem(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  int{}
// func (com *Dmsoft)FindPicMemE(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}
// func (com *Dmsoft)FindPicMemEx(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}

func (com *Dmsoft) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicS", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

// func (com *Dmsoft)FindShape(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) int{}
// func (com *Dmsoft)FindShapeE(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
// func (com *Dmsoft)FindShapeEx(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
// func (com *Dmsoft)FreePic(pic_name) int{}

func (com *Dmsoft) GetAveHSV(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveHSV", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *Dmsoft) GetAveRGB(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveRGB", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *Dmsoft) GetColor(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// func (com *Dmsoft)GetColorBGR(x,y)string{}

func (com *Dmsoft) GetColorHSV(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColorHSV", x, y)
	return ret.ToString()
}

// func (com *Dmsoft)GetColorNum(x1, y1, x2, y2, color, sim) int{}

func (com *Dmsoft) GetPicSize(picName string) string {
	ret, _ := com.dm.CallMethod("GetPicSize", picName)
	return ret.ToString()
}

// func (com *Dmsoft)GetScreenData(x1, y1, x2, y2) int{}
// func (com *Dmsoft)GetScreenDataBmp(x1,y1,x2,y2,data,size) int{}
// func (com *Dmsoft)ImageToBmp(pic_name,bmp_name) int{}

func (com *Dmsoft) IsDisplayDead(x1, y1, x2, y2, time int) int {
	ret, _ := com.dm.CallMethod("IsDisplayDead", x1, y1, x2, y2, time)
	return int(ret.Val)
}

// func (com *Dmsoft)LoadPic(pic_name) int{}
// func (com *Dmsoft)LoadPicByte(addr,size,pic_name) int{}
// func (com *Dmsoft)LoadPic(pic_name) int{}

func (com *Dmsoft) MatchPicName(picName string) string {
	ret, _ := com.dm.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *Dmsoft) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", mode, info)
	return int(ret.Val)
}

func (com *Dmsoft) SetPicPwd(pwd string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", pwd)
	return int(ret.Val)
}
