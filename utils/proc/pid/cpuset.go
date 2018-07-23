package pid

import (
	"io/ioutil"
	"path/filepath"
)

func CPUSetRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "cpuset"))
	return string(buf)
}