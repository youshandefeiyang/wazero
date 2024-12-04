//go:build tinygo

package sysfs

import (
	"os"

	"github.com/youshandefeiyang/wazero/experimental/sys"
)

func datasync(f *os.File) sys.Errno {
	return sys.ENOSYS
}
