// 文字识别

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

// AddDict 给指定的字库中添加一条字库信息
func (com *Dmsoft) AddDict(index int, dictInfo string) int {
	ret, _ := com.dm.CallMethod("AddDict", index, dictInfo)
	return int(ret.Val)
}

// ClearDict 清空指定的字库
func (com *Dmsoft) ClearDict(index int) int {
	ret, _ := com.dm.CallMethod("ClearDict", index)
	return int(ret.Val)
}

func (com *Dmsoft) EnableShareDict(enable int) int {
	ret, _ := com.dm.CallMethod("EnableShareDict", enable)
	return int(ret.Val)
}

func (com *Dmsoft) FetchWord(x1, y1, x2, y2 int, color, word string) string {
	ret, _ := com.dm.CallMethod("FetchWord", x1, y1, x2, y2, color, word)
	return ret.ToString()
}

func (com *Dmsoft) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrE(x1, y1, x2, y2 int, str string, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrFast", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrFastE(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrFastS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrWithFont(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, fontName string, fontSize, flag int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrWithFont", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrWithFontE(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("FindStrWithFontE", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *Dmsoft) FindStrWithFontEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("FindStrWithFontEx", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *Dmsoft) GetDict(index, fontIndex int) string {
	ret, _ := com.dm.CallMethod("GetDict", index, fontIndex)
	return ret.ToString()
}

func (com *Dmsoft) GetDictCount(index int) int {
	ret, _ := com.dm.CallMethod("GetDictCount", index)
	return int(ret.Val)
}

func (com *Dmsoft) GetDictInfo(str, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("GetDictInfo", str, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *Dmsoft) GetNowDict() int {
	ret, _ := com.dm.CallMethod("GetNowDict")
	return int(ret.Val)
}

func (com *Dmsoft) GetResultCount(str string) int {
	ret, _ := com.dm.CallMethod("GetResultCount", str)
	return int(ret.Val)
}

func (com *Dmsoft) GetResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("GetResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultCount(str string) int {
	ret, _ := com.dm.CallMethod("GetWordResultCount", str)
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("GetWordResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultStr(str string, index int) string {
	ret, _ := com.dm.CallMethod("GetWordResultCount", str, index)
	return ret.ToString()
}

func (com *Dmsoft) GetWords(x1, y1, x2, y2 int, color string, sim float32) string {
	ret, _ := com.dm.CallMethod("GetWords", x1, y1, x2, y2, color, sim)
	return ret.ToString()
}

func (com *Dmsoft) GetWordsNoDict(x1, y1, x2, y2 int, color string) string {
	ret, _ := com.dm.CallMethod("GetWordsNoDict", x1, y1, x2, y2, color)
	return ret.ToString()
}

func (com *Dmsoft) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrEx", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) OcrExOne(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrExOne", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) OcrInFile(x1, y1, x2, y2 int, picName, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrInFile", x1, y1, x2, y2, picName, colorFormat, sim)
	return ret.ToString()
}

func (com *Dmsoft) SaveDict(index int, file string) int {
	ret, _ := com.dm.CallMethod("SaveDict", index, file)
	return int(ret.Val)
}

func (com *Dmsoft) SetColGapNoDict(colGap int) int {
	ret, _ := com.dm.CallMethod("SetColGapNoDict", colGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetDict(index int, file string) int {
	ret, _ := com.dm.CallMethod("SetDict", index, file)
	return int(ret.Val)
}

func (com *Dmsoft) SetDictMem(index, addr, size int) int {
	ret, _ := com.dm.CallMethod("SetDictMem", index, addr, size)
	return int(ret.Val)
}

func (com *Dmsoft) SetDictPwd(pwd string) int {
	ret, _ := com.dm.CallMethod("SetDictPwd", pwd)
	return int(ret.Val)
}

func (com *Dmsoft) SetExactOcr(exactOcr int) int {
	ret, _ := com.dm.CallMethod("SetExactOcr", exactOcr)
	return int(ret.Val)
}

func (com *Dmsoft) SetMinColGap(minColGap int) int {
	ret, _ := com.dm.CallMethod("SetMinColGap", minColGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetMinRowGap(minRowGap int) int {
	ret, _ := com.dm.CallMethod("SetMinRowGap", minRowGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetRowGapNoDict(RowGap int) int {
	ret, _ := com.dm.CallMethod("SetRowGapNoDict", RowGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetWordGap(wordGap int) int {
	ret, _ := com.dm.CallMethod("SetWordGap", wordGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetWordGapNoDict(wordGap int) int {
	ret, _ := com.dm.CallMethod("SetWordGapNoDict", wordGap)
	return int(ret.Val)
}

func (com *Dmsoft) SetWordLineHeight(lineHeight int) int {
	ret, _ := com.dm.CallMethod("SetWordLineHeight", lineHeight)
	return int(ret.Val)
}

func (com *Dmsoft) SetWordLineHeightNoDict(lineHeight int) int {
	ret, _ := com.dm.CallMethod("SetWordLineHeightNoDict", lineHeight)
	return int(ret.Val)
}

func (com *Dmsoft) UseDict(index int) int {
	ret, _ := com.dm.CallMethod("UseDict", index)
	return int(ret.Val)
}
