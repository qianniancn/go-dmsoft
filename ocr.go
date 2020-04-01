// 文字识别

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

// AddDict 给指定的字库中添加一条字库信息
func (com *dmsoft) AddDict(index int, dictInfo string) int {
	ret, _ := com.dm.CallMethod("AddDict", index, dictInfo)
	return int(ret.Val)
}

// ClearDict 清空指定的字库
func (com *dmsoft) ClearDict(index int) int {
	ret, _ := com.dm.CallMethod("ClearDict", index)
	return int(ret.Val)
}

func (com *dmsoft) EnableShareDict(enable int) int {
	ret, _ := com.dm.CallMethod("EnableShareDict", enable)
	return int(ret.Val)
}

func (com *dmsoft) FetchWord(x1, y1, x2, y2 int, color, word string) string {
	ret, _ := com.dm.CallMethod("FetchWord", x1, y1, x2, y2, color, word)
	return ret.ToString()
}

func (com *dmsoft) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *dmsoft) FindStrE(x1, y1, x2, y2 int, str string, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrFast", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *dmsoft) FindStrFastE(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrFastExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("FindStrFastExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) FindStrFastS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrFastS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *dmsoft) FindStrS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *dmsoft) FindStrWithFont(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, fontName string, fontSize, flag int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrWithFont", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *dmsoft) FindStrWithFontE(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("FindStrWithFontE", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *dmsoft) FindStrWithFontEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("FindStrWithFontEx", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *dmsoft) GetDict(index, fontIndex int) string {
	ret, _ := com.dm.CallMethod("GetDict", index, fontIndex)
	return ret.ToString()
}

func (com *dmsoft) GetDictCount(index int) int {
	ret, _ := com.dm.CallMethod("GetDictCount", index)
	return int(ret.Val)
}

func (com *dmsoft) GetDictInfo(str, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("GetDictInfo", str, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *dmsoft) GetNowDict() int {
	ret, _ := com.dm.CallMethod("GetNowDict")
	return int(ret.Val)
}

func (com *dmsoft) GetResultCount(str string) int {
	ret, _ := com.dm.CallMethod("GetResultCount", str)
	return int(ret.Val)
}

func (com *dmsoft) GetResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("GetResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *dmsoft) GetWordResultCount(str string) int {
	ret, _ := com.dm.CallMethod("GetWordResultCount", str)
	return int(ret.Val)
}

func (com *dmsoft) GetWordResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("GetWordResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *dmsoft) GetWordResultStr(str string, index int) string {
	ret, _ := com.dm.CallMethod("GetWordResultCount", str, index)
	return ret.ToString()
}

func (com *dmsoft) GetWords(x1, y1, x2, y2 int, color string, sim float32) string {
	ret, _ := com.dm.CallMethod("GetWords", x1, y1, x2, y2, color, sim)
	return ret.ToString()
}

func (com *dmsoft) GetWordsNoDict(x1, y1, x2, y2 int, color string) string {
	ret, _ := com.dm.CallMethod("GetWordsNoDict", x1, y1, x2, y2, color)
	return ret.ToString()
}

func (com *dmsoft) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrEx", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) OcrExOne(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrExOne", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) OcrInFile(x1, y1, x2, y2 int, picName, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrInFile", x1, y1, x2, y2, picName, colorFormat, sim)
	return ret.ToString()
}

func (com *dmsoft) SaveDict(index int, file string) int {
	ret, _ := com.dm.CallMethod("SaveDict", index, file)
	return int(ret.Val)
}

func (com *dmsoft) SetColGapNoDict(colGap int) int {
	ret, _ := com.dm.CallMethod("SetColGapNoDict", colGap)
	return int(ret.Val)
}

func (com *dmsoft) SetDict(index int, file string) int {
	ret, _ := com.dm.CallMethod("SetDict", index, file)
	return int(ret.Val)
}

func (com *dmsoft) SetDictMem(index, addr, size int) int {
	ret, _ := com.dm.CallMethod("SetDictMem", index, addr, size)
	return int(ret.Val)
}

func (com *dmsoft) SetDictPwd(pwd string) int {
	ret, _ := com.dm.CallMethod("SetDictPwd", pwd)
	return int(ret.Val)
}

func (com *dmsoft) SetExactOcr(exactOcr int) int {
	ret, _ := com.dm.CallMethod("SetExactOcr", exactOcr)
	return int(ret.Val)
}

func (com *dmsoft) SetMinColGap(minColGap int) int {
	ret, _ := com.dm.CallMethod("SetMinColGap", minColGap)
	return int(ret.Val)
}

func (com *dmsoft) SetMinRowGap(minRowGap int) int {
	ret, _ := com.dm.CallMethod("SetMinRowGap", minRowGap)
	return int(ret.Val)
}

func (com *dmsoft) SetRowGapNoDict(RowGap int) int {
	ret, _ := com.dm.CallMethod("SetRowGapNoDict", RowGap)
	return int(ret.Val)
}

func (com *dmsoft) SetWordGap(wordGap int) int {
	ret, _ := com.dm.CallMethod("SetWordGap", wordGap)
	return int(ret.Val)
}

func (com *dmsoft) SetWordGapNoDict(wordGap int) int {
	ret, _ := com.dm.CallMethod("SetWordGapNoDict", wordGap)
	return int(ret.Val)
}

func (com *dmsoft) SetWordLineHeight(lineHeight int) int {
	ret, _ := com.dm.CallMethod("SetWordLineHeight", lineHeight)
	return int(ret.Val)
}

func (com *dmsoft) SetWordLineHeightNoDict(lineHeight int) int {
	ret, _ := com.dm.CallMethod("SetWordLineHeightNoDict", lineHeight)
	return int(ret.Val)
}

func (com *dmsoft) UseDict(index int) int {
	ret, _ := com.dm.CallMethod("UseDict", index)
	return int(ret.Val)
}
