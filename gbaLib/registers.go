package gbaLib

import (
	"runtime/volatile"
	"unsafe"
)

var (
	Register = struct {
		DispStat,
		KeyPad,
		IE,
		IF,
		KeyCnt *volatile.Register16
	}{
		DispStat: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004))),
		KeyPad:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000130))),
		KeyCnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000132))),
		IE:       (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000200))),
		IF:       (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000202))),
	}
)
