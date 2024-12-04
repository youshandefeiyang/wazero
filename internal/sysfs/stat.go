package sysfs

import (
	"io/fs"

	experimentalsys "github.com/youshandefeiyang/wazero/experimental/sys"
	"github.com/youshandefeiyang/wazero/sys"
)

func defaultStatFile(f fs.File) (sys.Stat_t, experimentalsys.Errno) {
	if info, err := f.Stat(); err != nil {
		return sys.Stat_t{}, experimentalsys.UnwrapOSError(err)
	} else {
		return sys.NewStat_t(info), 0
	}
}
