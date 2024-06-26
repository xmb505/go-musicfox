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

const SignaturePlaybackMediaMarker string = "rc(Windows.Media.Playback.PlaybackMediaMarker;{c4d22f5c-3c1c-4444-b6b9-778b0422d41a})"

type PlaybackMediaMarker struct {
	ole.IUnknown
}

func (impl *PlaybackMediaMarker) GetTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarker))
	defer itf.Release()
	v := (*iPlaybackMediaMarker)(unsafe.Pointer(itf))
	return v.GetTime()
}

func (impl *PlaybackMediaMarker) GetMediaMarkerType() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarker))
	defer itf.Release()
	v := (*iPlaybackMediaMarker)(unsafe.Pointer(itf))
	return v.GetMediaMarkerType()
}

func (impl *PlaybackMediaMarker) GetText() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarker))
	defer itf.Release()
	v := (*iPlaybackMediaMarker)(unsafe.Pointer(itf))
	return v.GetText()
}

const GUIDiPlaybackMediaMarker string = "c4d22f5c-3c1c-4444-b6b9-778b0422d41a"
const SignatureiPlaybackMediaMarker string = "{c4d22f5c-3c1c-4444-b6b9-778b0422d41a}"

type iPlaybackMediaMarker struct {
	ole.IInspectable
}

type iPlaybackMediaMarkerVtbl struct {
	ole.IInspectableVtbl

	GetTime            uintptr
	GetMediaMarkerType uintptr
	GetText            uintptr
}

func (v *iPlaybackMediaMarker) VTable() *iPlaybackMediaMarkerVtbl {
	return (*iPlaybackMediaMarkerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iPlaybackMediaMarker) GetTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iPlaybackMediaMarker) GetMediaMarkerType() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMediaMarkerType,
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

func (v *iPlaybackMediaMarker) GetText() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetText,
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

const GUIDiPlaybackMediaMarkerFactory string = "8c530a78-e0ae-4e1a-a8c8-e23f982a937b"
const SignatureiPlaybackMediaMarkerFactory string = "{8c530a78-e0ae-4e1a-a8c8-e23f982a937b}"

type iPlaybackMediaMarkerFactory struct {
	ole.IInspectable
}

type iPlaybackMediaMarkerFactoryVtbl struct {
	ole.IInspectableVtbl

	PlaybackMediaMarkerCreateFromTime uintptr
	PlaybackMediaMarkerCreate         uintptr
}

func (v *iPlaybackMediaMarkerFactory) VTable() *iPlaybackMediaMarkerFactoryVtbl {
	return (*iPlaybackMediaMarkerFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}

func PlaybackMediaMarkerCreateFromTime(value foundation.TimeSpan) (*PlaybackMediaMarker, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.PlaybackMediaMarker", ole.NewGUID(GUIDiPlaybackMediaMarkerFactory))
	if err != nil {
		return nil, err
	}
	v := (*iPlaybackMediaMarkerFactory)(unsafe.Pointer(inspectable))

	var out *PlaybackMediaMarker
	hr, _, _ := syscall.SyscallN(
		v.VTable().PlaybackMediaMarkerCreateFromTime,
		0,                             // this is a static func, so there's no this
		uintptr(value.Duration),       // in foundation.TimeSpan
		uintptr(unsafe.Pointer(&out)), // out PlaybackMediaMarker
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func PlaybackMediaMarkerCreate(value foundation.TimeSpan, mediaMarketType string, text string) (*PlaybackMediaMarker, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.PlaybackMediaMarker", ole.NewGUID(GUIDiPlaybackMediaMarkerFactory))
	if err != nil {
		return nil, err
	}
	v := (*iPlaybackMediaMarkerFactory)(unsafe.Pointer(inspectable))

	var out *PlaybackMediaMarker
	mediaMarketTypeHStr, err := ole.NewHString(mediaMarketType)
	if err != nil {
		return nil, err
	}
	textHStr, err := ole.NewHString(text)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().PlaybackMediaMarkerCreate,
		0,                             // this is a static func, so there's no this
		uintptr(value.Duration),       // in foundation.TimeSpan
		uintptr(mediaMarketTypeHStr),  // in string
		uintptr(textHStr),             // in string
		uintptr(unsafe.Pointer(&out)), // out PlaybackMediaMarker
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
