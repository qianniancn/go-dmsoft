package dmsoft

func (com *Dmsoft) DoubleToData(value float64) string {
	ret, _ := com.dm.CallMethod("DoubleToData", value)
	return ret.ToString()
}

func (com *Dmsoft) FindData(hwnd int, addr_range string, data string) string {
	ret, _ := com.dm.CallMethod("FindData", hwnd, addr_range, data)
	return ret.ToString()
}

func (com *Dmsoft) FindDataEx(hwnd int, addr_range string, data string, step int, multi_thread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindDataEx", hwnd, addr_range, data, step, multi_thread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindDouble(hwnd int, addr_range string, double_value_min float64, double_value_max float64) string {
	ret, _ := com.dm.CallMethod("FindDouble", hwnd, addr_range, double_value_min, double_value_max)
	return ret.ToString()
}

func (com *Dmsoft) FindDoubleEx(hwnd int, addr_range string, double_value_min float64, double_value_max float64, step int, multi_thread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindDoubleEx", hwnd, addr_range, double_value_min, double_value_max, step, multi_thread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindFloat(hwnd int, addr_range string, float_value_min float32, float_value_max float32) string {
	ret, _ := com.dm.CallMethod("FindFloat", hwnd, addr_range, float_value_min, float_value_max)
	return ret.ToString()
}

func (com *Dmsoft) FindFloatEx(hwnd int, addr_range string, float_value_min float32, float_value_max float32, step int, multi_thread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindFloatEx", hwnd, addr_range, float_value_min, float_value_max, step, multi_thread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindInt(hwnd int, addr_range string, int_value_min int64, int_value_max int64, itype int) string {
	ret, _ := com.dm.CallMethod("FindInt", hwnd, addr_range, int_value_min, int_value_max, itype)
	return ret.ToString()
}

func (com *Dmsoft) FindIntEx(hwnd int, addr_range string, int_value_min int64, int_value_max int64, itype int, step int, multi_thread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindIntEx", hwnd, addr_range, int_value_min, int_value_max, itype, step, multi_thread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindString(hwnd int, addr_range string, string_value string, itype int) string {
	ret, _ := com.dm.CallMethod("FindString", hwnd, addr_range, string_value, itype)
	return ret.ToString()
}

func (com *Dmsoft) FindStringEx(hwnd int, addr_range string, string_value string, itype int, step int, multi_thread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindStringEx", hwnd, addr_range, string_value, itype, step, multi_thread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FloatToData(value float32) string {
	ret, _ := com.dm.CallMethod("FloatToData", value)
	return ret.ToString()
}

func (com *Dmsoft) FreeProcessMemory(hwnd int) int {
	ret, _ := com.dm.CallMethod("FreeProcessMemory", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) GetCommandLine(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetCommandLine", hwnd)
	return ret.ToString()
}

func (com *Dmsoft) GetModuleBaseAddr(hwnd int, module string) int64 {
	ret, _ := com.dm.CallMethod("GetModuleBaseAddr", hwnd, module)
	return ret.Val
}

func (com *Dmsoft) GetModuleSize(hwnd int, module string) int {
	ret, _ := com.dm.CallMethod("GetModuleSize", hwnd, module)
	return int(ret.Val)
}

func (com *Dmsoft) GetRemoteApiAddress(hwnd int, base_addr int64, fun_name string) int64 {
	ret, _ := com.dm.CallMethod("GetRemoteApiAddress", hwnd, base_addr, fun_name)
	return ret.Val
}

func (com *Dmsoft) Int64ToInt32(value int64) int {
	ret, _ := com.dm.CallMethod("Int64ToInt32", value)
	return int(ret.Val)
}

func (com *Dmsoft) IntToData(value int64, itype int) string {
	ret, _ := com.dm.CallMethod("IntToData", value, itype)
	return ret.ToString()
}

func (com *Dmsoft) OpenProcess(pid int) int {
	ret, _ := com.dm.CallMethod("OpenProcess", pid)
	return int(ret.Val)
}

func (com *Dmsoft) ReadData(hwnd int, addr string, len int) string {
	ret, _ := com.dm.CallMethod("ReadData", hwnd, addr, len)
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddr(hwnd int, addr int64, len int) string {
	ret, _ := com.dm.CallMethod("ReadDataAddr", hwnd, addr, len)
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddrToBin(hwnd int, addr int64, len int) int {
	ret, _ := com.dm.CallMethod("ReadDataAddrToBin", hwnd, addr, len)
	return int(ret.Val)
}

func (com *Dmsoft) ReadDataToBin(hwnd int, addr string, len int) int {
	ret, _ := com.dm.CallMethod("ReadDataToBin", hwnd, addr, len)
	return int(ret.Val)
}

func (com *Dmsoft) ReadDouble(hwnd int, addr string) float64 {
	ret, _ := com.dm.CallMethod("ReadDouble", hwnd, addr)
	return float64(ret.Val)
}

func (com *Dmsoft) ReadDoubleAddr(hwnd int, addr int64) float64 {
	ret, _ := com.dm.CallMethod("ReadDoubleAddr", hwnd, addr)
	return float64(ret.Val)
}

func (com *Dmsoft) ReadFloat(hwnd int, addr string) float32 {
	ret, _ := com.dm.CallMethod("ReadFloat", hwnd, addr)
	return float32(ret.Val)
}

func (com *Dmsoft) ReadFloatAddr(hwnd int, addr int64) float32 {
	ret, _ := com.dm.CallMethod("ReadFloatAddr", hwnd, addr)
	return float32(ret.Val)
}

func (com *Dmsoft) ReadInt(hwnd int, addr string, itype int) int64 {
	ret, _ := com.dm.CallMethod("ReadInt", hwnd, addr, itype)
	return ret.Val
}

func (com *Dmsoft) ReadIntAddr(hwnd int, addr int64, itype int) int64 {
	ret, _ := com.dm.CallMethod("ReadIntAddr", hwnd, addr, itype)
	return ret.Val
}

func (com *Dmsoft) ReadString(hwnd int, addr string, itype int, len int) string {
	ret, _ := com.dm.CallMethod("ReadString", hwnd, addr, itype, len)
	return ret.ToString()
}

func (com *Dmsoft) ReadStringAddr(hwnd int, addr int64, itype int, len int) string {
	ret, _ := com.dm.CallMethod("ReadStringAddr", hwnd, addr, itype, len)
	return ret.ToString()
}

func (com *Dmsoft) SetMemoryFindResultToFile(file string) int {
	ret, _ := com.dm.CallMethod("SetMemoryFindResultToFile", file)
	return int(ret.Val)
}

func (com *Dmsoft) SetMemoryHwndAsProcessId(en int) int {
	ret, _ := com.dm.CallMethod("SetMemoryHwndAsProcessId", en)
	return int(ret.Val)
}

func (com *Dmsoft) StringToData(value string, itype int) string {
	ret, _ := com.dm.CallMethod("StringToData", value, itype)
	return ret.ToString()
}

func (com *Dmsoft) TerminateProcess(pid int) int {
	ret, _ := com.dm.CallMethod("TerminateProcess", pid)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualAllocEx(hwnd int, addr int64, size int, itype int) int64 {
	ret, _ := com.dm.CallMethod("VirtualAllocEx", hwnd, addr, size, itype)
	return ret.Val
}

func (com *Dmsoft) VirtualFreeEx(hwnd int, addr int64) int {
	ret, _ := com.dm.CallMethod("VirtualFreeEx", hwnd, addr)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualProtectEx(hwnd int, addr int64, size int, itype int, old_protect int) int {
	ret, _ := com.dm.CallMethod("VirtualProtectEx", hwnd, addr, size, itype, old_protect)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualQueryEx(hwnd int, addr int64, pmbi int) string {
	ret, _ := com.dm.CallMethod("VirtualQueryEx", hwnd, addr, pmbi)
	return ret.ToString()
}

func (com *Dmsoft) WriteData(hwnd int, addr string, data string) int {
	ret, _ := com.dm.CallMethod("WriteData", hwnd, addr, data)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDataAddr(hwnd int, addr int64, data string) int {
	ret, _ := com.dm.CallMethod("WriteDataAddr", hwnd, addr, data)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDataAddrFromBin(hwnd int, addr int64, data int, len int) int {
	ret, _ := com.dm.CallMethod("WriteDataAddrFromBin", hwnd, addr, data, len)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDataFromBin(hwnd int, addr string, data int, len int) int {
	ret, _ := com.dm.CallMethod("WriteDataFromBin", hwnd, addr, data, len)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDouble(hwnd int, addr string, v float32) int {
	ret, _ := com.dm.CallMethod("WriteDouble", hwnd, addr, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDoubleAddr(hwnd int, addr int64, v float32) int {
	ret, _ := com.dm.CallMethod("WriteDoubleAddr", hwnd, addr, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteFloat(hwnd int, addr string, v float32) int {
	ret, _ := com.dm.CallMethod("WriteFloat", hwnd, addr, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteFloatAddr(hwnd int, addr int64, v float32) int {
	ret, _ := com.dm.CallMethod("WriteFloatAddr", hwnd, addr, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteInt(hwnd int, addr string, itype int, v float32) int {
	ret, _ := com.dm.CallMethod("WriteInt", hwnd, addr, itype, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteIntAddr(hwnd int, addr int64, itype int, v float32) int {
	ret, _ := com.dm.CallMethod("WriteIntAddr", hwnd, addr, itype, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteString(hwnd int, addr string, itype int, v float32) int {
	ret, _ := com.dm.CallMethod("WriteString", hwnd, addr, itype, v)
	return int(ret.Val)
}

func (com *Dmsoft) WriteStringAddr(hwnd int, addr int64, itype int, v float32) int {
	ret, _ := com.dm.CallMethod("WriteStringAddr", hwnd, addr, itype, v)
	return int(ret.Val)
}
