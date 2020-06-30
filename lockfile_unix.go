// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris aix

package lockfile

import (
	"errors"
	"os"
	"syscall"
)

func isRunning(pid int) (bool, error) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return false, err
	}

	if err := proc.Signal(syscall.Signal(0)); err != nil {
		// syscall.EPERM will be returned if the process exists, but we don't
		// have permissions to send signals to it. All other errors will mean
		// that the process doesn't exist.
		if !errors.Is(err, syscall.EPERM) {
			return false, nil
		}
	}

	return true, nil
}
