package pid

import (
	"io/ioutil"
	"path/filepath"
)

func CGroupRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "cgroup"))
	return string(buf)
}