package pid

import (
	"io/ioutil"
	"path/filepath"
)

func SchedRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "sched"))
	return string(buf)
}