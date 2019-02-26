// +build linux

package water

import (
	"fmt"
    "unsafe"
    "os"
)

// NewTAP creates a new TAP interface whose name is ifName. If ifName is empty, a
// default name (tap0, tap1, ... ) will be assigned. ifName should not exceed
// 16 bytes. TAP interfaces are not supported on darwin.
// ifName cannot be specified on windows, you will need ifce.Name() to use some cmds.
//
// Deprecated: This function may be removed in the future. Please use New() instead.
func NewTAP(ifName string) (ifce *Interface, err error) {
	fmt.Println("Deprecated: NewTAP(..) may be removed in the future. Please use New() instead.")
	config := Config{DeviceType: TAP}
	config.Name = ifName
	return newTAP(config)
}

// NewTUN creates a new TUN interface whose name is ifName. If ifName is empty, a
// default name (tap0, tap1, ... ) will be assigned. ifName should not exceed
// ifName cannot be specified on windows, you will need ifce.Name() to use some cmds.
//
// Deprecated: This function will be removed in the future. Please use New() instead.
func NewTUN(ifName string) (ifce *Interface, err error) {
	fmt.Println("Deprecated: NewTUN(..) may be removed in the future. Please use New() instead.")
	config := Config{DeviceType: TUN}
	config.Name = ifName
	return newTUN(config)
}

func (ifce *Interface) QueueAttach() (err error) {
    var qreq  ifReq
    ifce.attached = true
    qreq.Flags = cIFF_ATTACH_QUEUE
    if err := ioctl(ifce.File.Fd(), uintptr(TUNSETQUEUE), uintptr(unsafe.Pointer(&qreq))); err != nil {
        os.NewSyscallError("ioctl TUNSETQUEUE attach:", err)
        return err
    }
    return nil
}

func (ifce *Interface) QueueDetach() (err error) {
    var qreq  ifReq
    ifce.attached = false
    qreq.Flags = cIFF_DETACH_QUEUE
    if err := ioctl(ifce.File.Fd(), uintptr(TUNSETQUEUE), uintptr(unsafe.Pointer(&qreq))); err != nil {
        os.NewSyscallError("ioctl TUNSETQUEUE detach:", err)
        return err
    }
    return nil
}

func (ifce *Interface) IsQueueAttached() (a bool) {
    return ifce.attached
}