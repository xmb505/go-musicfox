// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package storage

type StreamedFileFailureMode int32

const SignatureStreamedFileFailureMode string = "enum(Windows.Storage.StreamedFileFailureMode;i4)"

const (
	StreamedFileFailureModeFailed               StreamedFileFailureMode = 0
	StreamedFileFailureModeCurrentlyUnavailable StreamedFileFailureMode = 1
	StreamedFileFailureModeIncomplete           StreamedFileFailureMode = 2
)
