// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package backgroundtransfer

type BackgroundTransferStatus int32

const SignatureBackgroundTransferStatus string = "enum(Windows.Networking.BackgroundTransfer.BackgroundTransferStatus;i4)"

const (
	BackgroundTransferStatusIdle                            BackgroundTransferStatus = 0
	BackgroundTransferStatusRunning                         BackgroundTransferStatus = 1
	BackgroundTransferStatusPausedByApplication             BackgroundTransferStatus = 2
	BackgroundTransferStatusPausedCostedNetwork             BackgroundTransferStatus = 3
	BackgroundTransferStatusPausedNoNetwork                 BackgroundTransferStatus = 4
	BackgroundTransferStatusCompleted                       BackgroundTransferStatus = 5
	BackgroundTransferStatusCanceled                        BackgroundTransferStatus = 6
	BackgroundTransferStatusError                           BackgroundTransferStatus = 7
	BackgroundTransferStatusPausedRecoverableWebErrorStatus BackgroundTransferStatus = 8
	BackgroundTransferStatusPausedSystemPolicy              BackgroundTransferStatus = 32
)
