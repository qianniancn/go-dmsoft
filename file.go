// 文件操作可以通过go来实现
//
//	// long ([a-z]+)\((.*)\)$
//	func (com *Dmsoft) $1($2) int {\n\tret, _ := com.dm.CallMethod("$1", $2)\n\treturn int(ret.Val)\n}
package dmsoft

func (com *Dmsoft) CopyFile(src_file string, dst_file string, over int) int {
	ret, _ := com.dm.CallMethod("CopyFile", src_file, dst_file, over)
	return int(ret.Val)
}
func (com *Dmsoft) CreateFolder(folder string) int {
	ret, _ := com.dm.CallMethod("CreateFolder", folder)
	return int(ret.Val)
}
func (com *Dmsoft) DecodeFile(file string, pwd string) int {
	ret, _ := com.dm.CallMethod("DecodeFile", file, pwd)
	return int(ret.Val)
}
func (com *Dmsoft) DeleteFile(file string) int {
	ret, _ := com.dm.CallMethod("DeleteFile", file)
	return int(ret.Val)
}
func (com *Dmsoft) DeleteFolder(folder string) int {
	ret, _ := com.dm.CallMethod("DeleteFolder", folder)
	return int(ret.Val)
}
func (com *Dmsoft) DeleteIni(section string, key string, file string) int {
	ret, _ := com.dm.CallMethod("DeleteIni", section, key, file)
	return int(ret.Val)
}
func (com *Dmsoft) DeleteIniPwd(section string, key string, file string, pwd string) int {
	ret, _ := com.dm.CallMethod("DeleteIniPwd", section, key, file, pwd)
	return int(ret.Val)
}
func (com *Dmsoft) DownloadFile(url string, save_file string, timeout string) int {
	ret, _ := com.dm.CallMethod("DownloadFile", url, save_file, timeout)
	return int(ret.Val)
}
func (com *Dmsoft) EncodeFile(file string, pwd string) int {
	ret, _ := com.dm.CallMethod("EncodeFile", file, pwd)
	return int(ret.Val)
}
func (com *Dmsoft) EnumIniKey(section string, file string) string {
	ret, _ := com.dm.CallMethod("EnumIniKey", section, file)
	return ret.ToString()
}

func (com *Dmsoft) EnumIniKeyPwd(section string, file string, pwd string) string {
	ret, _ := com.dm.CallMethod("EnumIniKeyPwd", section, file, pwd)
	return ret.ToString()
}
func (com *Dmsoft) EnumIniSection(file string) string {
	ret, _ := com.dm.CallMethod("EnumIniSection", file)
	return ret.ToString()
}

func (com *Dmsoft) EnumIniSectionPwd(file string, pwd string) string {
	ret, _ := com.dm.CallMethod("EnumIniSectionPwd", file, pwd)
	return ret.ToString()
}
func (com *Dmsoft) GetFileLength(file string) int {
	ret, _ := com.dm.CallMethod("GetFileLength", file)
	return int(ret.Val)
}

func (com *Dmsoft) GetRealPath(path string) string {
	ret, _ := com.dm.CallMethod("GetRealPath", path)
	return ret.ToString()
}
func (com *Dmsoft) IsFileExist(file string) int {
	ret, _ := com.dm.CallMethod("IsFileExist", file)
	return int(ret.Val)
}

func (com *Dmsoft) IsFolderExist(folder string) int {
	ret, _ := com.dm.CallMethod("IsFolderExist", folder)
	return int(ret.Val)
}
func (com *Dmsoft) MoveFile(src_file string, dst_file string) int {
	ret, _ := com.dm.CallMethod("MoveFile", src_file, dst_file)
	return int(ret.Val)
}

func (com *Dmsoft) ReadFile(file string) string {
	ret, _ := com.dm.CallMethod("ReadFile", file)
	return ret.ToString()
}
func (com *Dmsoft) ReadIni(section string, key string, file string) string {
	ret, _ := com.dm.CallMethod("ReadIni", section, key, file)
	return ret.ToString()
}

func (com *Dmsoft) ReadIniPwd(section string, key string, file string, pwd string) string {
	ret, _ := com.dm.CallMethod("ReadIniPwd", section, key, file, pwd)
	return ret.ToString()
}
func (com *Dmsoft) SelectDirectory() string {
	ret, _ := com.dm.CallMethod("SelectDirectory")
	return ret.ToString()
}
func (com *Dmsoft) SelectFile() string {
	ret, _ := com.dm.CallMethod("SelectFile")
	return ret.ToString()
}
func (com *Dmsoft) WriteFile(file string, content string) int {
	ret, _ := com.dm.CallMethod("WriteFile", file, content)
	return int(ret.Val)
}
func (com *Dmsoft) WriteIni(section string, key string, value string, file string) int {
	ret, _ := com.dm.CallMethod("WriteIni", section, key, value, file)
	return int(ret.Val)
}

func (com *Dmsoft) WriteIniPwd(section string, key string, value string, file string, pwd string) int {
	ret, _ := com.dm.CallMethod("WriteIniPwd", section, key, value, file, pwd)
	return int(ret.Val)
}
