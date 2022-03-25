// 答题

package dmsoft


//可以把上次FaqPost的发送取消,接着下一次FaqPost
func (com *DmSoft) FaqCancel() int {
	ret, _ := com.dm.CallMethod("FaqCancel")
	return int(ret.Val)
}


//截取指定范围内的动画或者图像,并返回此句柄.
func (com *DmSoft) FaqCapture(x1, y1, x2, y2, quality, delay, time int) int {
	ret, _ := com.dm.CallMethod("FaqCapture",x1, y1, x2, y2, quality, delay, time)
	return int(ret.Val)
}

//截取指定图片中的图像,并返回此句柄.
func (com *DmSoft) FaqCaptureFromFile(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("FaqCaptureFromFile", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

//从给定的字符串(也可以算是文字类型的问题),获取此句柄.
func (com *DmSoft) FaqCaptureString(str string) int {
	ret, _ := com.dm.CallMethod("FaqCaptureString", str)
	return int(ret.Val)
}

//获取由FaqPost发送后，由服务器返回的答案.
func (com *DmSoft) FaqFetch() string {
	ret, _ := com.dm.CallMethod("FaqFetch")
	return ret.ToString()
}

//获取句柄所对应的数据包的大小,单位是字节
func (com *DmSoft) FaqGetSize(handle int) int {
	ret, _ := com.dm.CallMethod("FaqGetSize", handle)
	return int(ret.Val)
}

//用于判断当前对象是否有发送过答题(FaqPost)
func (com *DmSoft) FaqIsPosted() int {
	ret, _ := com.dm.CallMethod("FaqIsPosted")
	return int(ret.Val)
}

//发送指定的图像句柄到指定的服务器,并立即返回(异步操作).
func (com *DmSoft) FaqPost(server string, handle, request_type, time_out int) int {
	ret, _ := com.dm.CallMethod("FaqPost", server, handle, request_type, time_out)
	return int(ret.Val)
}

//发送指定的图像句柄到指定的服务器,并等待返回结果(同步等待).
func (com *DmSoft) FaqSend(server string, handle, request_type, time_out int) string {
	ret, _ := com.dm.CallMethod("FaqSend", server, handle, request_type, time_out)
	return ret.ToString()
}
