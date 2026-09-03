// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build (linux || aix || zos) && !appengine && !tinygo

package color

import (
	"syscall"
	"unsafe"
)

// isTerminal uses the TIOCGWINSZ ioctl to determine whether this fd refers to
// a "proper" terminal or not
func isTerminal(fd int) bool {
	// _, err := unix.IoctlGetWinsize(int(fd), unix.TIOCGWINSZ)
	var arg [8]byte // = unsafe.Sizeof(WinSize)
	var req uintptr = syscall.TIOCGWINSZ

	// ~> ioctlPtr(SYS_IOCTL, fd, req, &WinSize{})
	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(unsafe.Pointer(&arg)))

	return 0 == e1
}

// type _winsize struct {
// 	Row    uint16
// 	Col    uint16
// 	Xpixel uint16
// 	Ypixel uint16
// }
