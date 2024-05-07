// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureMediaPlaybackCommandManagerPositionReceivedEventArgs string = "rc(Windows.Media.Playback.MediaPlaybackCommandManagerPositionReceivedEventArgs;{5591a754-d627-4bdd-a90d-86a015b24902})"

type MediaPlaybackCommandManagerPositionReceivedEventArgs struct {
	ole.IUnknown
}

func (impl *MediaPlaybackCommandManagerPositionReceivedEventArgs) GetHandled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPositionReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPositionReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetHandled()
}

func (impl *MediaPlaybackCommandManagerPositionReceivedEventArgs) SetHandled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPositionReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPositionReceivedEventArgs)(unsafe.Pointer(itf))
	return v.SetHandled(value)
}

func (impl *MediaPlaybackCommandManagerPositionReceivedEventArgs) GetPosition() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPositionReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPositionReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetPosition()
}

func (impl *MediaPlaybackCommandManagerPositionReceivedEventArgs) GetDeferral() (*foundation.Deferral, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPositionReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPositionReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetDeferral()
}

const GUIDiMediaPlaybackCommandManagerPositionReceivedEventArgs string = "5591a754-d627-4bdd-a90d-86a015b24902"
const SignatureiMediaPlaybackCommandManagerPositionReceivedEventArgs string = "{5591a754-d627-4bdd-a90d-86a015b24902}"

type iMediaPlaybackCommandManagerPositionReceivedEventArgs struct {
	ole.IInspectable
}

type iMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl struct {
	ole.IInspectableVtbl

	GetHandled  uintptr
	SetHandled  uintptr
	GetPosition uintptr
	GetDeferral uintptr
}

func (v *iMediaPlaybackCommandManagerPositionReceivedEventArgs) VTable() *iMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl {
	return (*iMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackCommandManagerPositionReceivedEventArgs) GetHandled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetHandled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerPositionReceivedEventArgs) SetHandled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetHandled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackCommandManagerPositionReceivedEventArgs) GetPosition() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPosition,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerPositionReceivedEventArgs) GetDeferral() (*foundation.Deferral, error) {
	var out *foundation.Deferral
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDeferral,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.Deferral
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
