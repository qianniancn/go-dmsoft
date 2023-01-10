//go:build windows
// +build windows

// export GOARCH=386

package dmsoft

import (
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	dmReg32      = syscall.NewLazyDLL("DmReg.dll")
	_SetDllPathA = dmReg32.NewProc("SetDllPathA")
	_SetDllPathW = dmReg32.NewProc("SetDllPathW")
)

type Dmsoft struct {
	dm       *ole.IDispatch
	IUnknown *ole.IUnknown
}

func NewDmsoft() *Dmsoft {
	var com Dmsoft
	var err error
	ole.CoInitialize(0)
	com.IUnknown, err = oleutil.CreateObject("dm.dmsoft")
	if err != nil {
		panic(err)
	}
	com.dm, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	return &com
}

// Release 释放
func (com *Dmsoft) Release() {
	com.IUnknown.Release()
	com.dm.Release()
	ole.CoUninitialize()
}

// SetDllPathA Ascii
func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathA.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

// SetDllPathW Unicode
func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathW.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}
