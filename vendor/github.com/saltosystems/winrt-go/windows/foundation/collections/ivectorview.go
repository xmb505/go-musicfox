// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package collections

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIVectorView string = "bbe1fa4c-b0e3-4583-baef-1f1b2e483e56"
const SignatureIVectorView string = "{bbe1fa4c-b0e3-4583-baef-1f1b2e483e56}"

type IVectorView struct {
	ole.IInspectable
}

type IVectorViewVtbl struct {
	ole.IInspectableVtbl

	GetAt   uintptr
	GetSize uintptr
	IndexOf uintptr
	GetMany uintptr
}

func (v *IVectorView) VTable() *IVectorViewVtbl {
	return (*IVectorViewVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IVectorView) GetAt(index uint32) (unsafe.Pointer, error) {
	var out unsafe.Pointer
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAt,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(index),                // in uint32
		uintptr(unsafe.Pointer(&out)), // out unsafe.Pointer
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IVectorView) GetSize() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *IVectorView) IndexOf(value unsafe.Pointer) (uint32, bool, error) {
	var index uint32
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().IndexOf,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(value),                  // in unsafe.Pointer
		uintptr(unsafe.Pointer(&index)), // out uint32
		uintptr(unsafe.Pointer(&out)),   // out bool
	)

	if hr != 0 {
		return 0, false, ole.NewError(hr)
	}

	return index, out, nil
}

func (v *IVectorView) GetMany(startIndex uint32, itemsSize uint32) ([]unsafe.Pointer, uint32, error) {
	var items []unsafe.Pointer = make([]unsafe.Pointer, itemsSize)
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMany,
		uintptr(unsafe.Pointer(v)),         // this
		uintptr(startIndex),                // in uint32
		uintptr(itemsSize),                 // in uint32
		uintptr(unsafe.Pointer(&items[0])), // out unsafe.Pointer
		uintptr(unsafe.Pointer(&out)),      // out uint32
	)

	if hr != 0 {
		return nil, 0, ole.NewError(hr)
	}

	return items, out, nil
}
