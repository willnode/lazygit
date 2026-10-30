// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package gocui

import (
	"os"
	"os/signal"

	"github.com/go-errors/errors"
	"golang.org/x/sys/unix"
)

// getTermWindowSize is get terminal window size on linux or unix.
// When gocui run inside the docker contaienr need to check and get the window size.
func (g *Gui) getTermWindowSize() (int, int, error) {
	var termw, termh int

	out, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return 0, 0, err
	}
	defer out.Close()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, unix.SIGWINCH, unix.SIGINT)

	for {
		ws, err := unix.IoctlGetWinsize(int(out.Fd()), unix.TIOCGWINSZ)

		if err == nil {
			// check terminal window size
			termw, termh = int(ws.Col), int(ws.Row)
			if termw > 0 && termh > 0 {
				return termw, termh, nil
			}
		}

		signal := <-signalCh
		switch signal {
		// when the terminal window size is changed
		case unix.SIGWINCH:
			continue
		// ctrl + c to cancel
		case unix.SIGINT:
			return 0, 0, errors.New("stop to get term window size")
		}
	}
}