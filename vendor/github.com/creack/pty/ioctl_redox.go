//go:build redox
// +build redox

package pty

import "golang.org/x/sys/unix"

const (
	TIOCGWINSZ = unix.TIOCGWINSZ
	TIOCSWINSZ = unix.TIOCSWINSZ
)

func ioctlInner(fd, cmd, ptr uintptr) error {
	err := unix.Ioctl(int(fd), uint(cmd), ptr)
	if err != nil {
		return err
	}
	return nil
}