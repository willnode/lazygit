//go:build !linux && !darwin && !freebsd && !dragonfly && !netbsd && !openbsd && !redox && !solaris && !zos
// +build !linux,!darwin,!freebsd,!dragonfly,!netbsd,!openbsd,!redox,!solaris,!zos

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
