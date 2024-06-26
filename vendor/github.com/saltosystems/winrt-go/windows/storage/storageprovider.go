// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package storage

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureStorageProvider string = "rc(Windows.Storage.StorageProvider;{e705eed4-d478-47d6-ba46-1a8ebe114a20})"

type StorageProvider struct {
	ole.IUnknown
}

func (impl *StorageProvider) GetId() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiStorageProvider))
	defer itf.Release()
	v := (*iStorageProvider)(unsafe.Pointer(itf))
	return v.GetId()
}

func (impl *StorageProvider) GetDisplayName() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiStorageProvider))
	defer itf.Release()
	v := (*iStorageProvider)(unsafe.Pointer(itf))
	return v.GetDisplayName()
}

func (impl *StorageProvider) IsPropertySupportedForPartialFileAsync(propertyCanonicalName string) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiStorageProvider2))
	defer itf.Release()
	v := (*iStorageProvider2)(unsafe.Pointer(itf))
	return v.IsPropertySupportedForPartialFileAsync(propertyCanonicalName)
}

const GUIDiStorageProvider string = "e705eed4-d478-47d6-ba46-1a8ebe114a20"
const SignatureiStorageProvider string = "{e705eed4-d478-47d6-ba46-1a8ebe114a20}"

type iStorageProvider struct {
	ole.IInspectable
}

type iStorageProviderVtbl struct {
	ole.IInspectableVtbl

	GetId          uintptr
	GetDisplayName uintptr
}

func (v *iStorageProvider) VTable() *iStorageProviderVtbl {
	return (*iStorageProviderVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iStorageProvider) GetId() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetId,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *iStorageProvider) GetDisplayName() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDisplayName,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

const GUIDiStorageProvider2 string = "010d1917-3404-414b-9fd7-cd44472eaa39"
const SignatureiStorageProvider2 string = "{010d1917-3404-414b-9fd7-cd44472eaa39}"

type iStorageProvider2 struct {
	ole.IInspectable
}

type iStorageProvider2Vtbl struct {
	ole.IInspectableVtbl

	IsPropertySupportedForPartialFileAsync uintptr
}

func (v *iStorageProvider2) VTable() *iStorageProvider2Vtbl {
	return (*iStorageProvider2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iStorageProvider2) IsPropertySupportedForPartialFileAsync(propertyCanonicalName string) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	propertyCanonicalNameHStr, err := ole.NewHString(propertyCanonicalName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().IsPropertySupportedForPartialFileAsync,
		uintptr(unsafe.Pointer(v)),         // this
		uintptr(propertyCanonicalNameHStr), // in string
		uintptr(unsafe.Pointer(&out)),      // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
