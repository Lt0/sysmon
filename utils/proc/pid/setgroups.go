package pid

import (
	"io/ioutil"
	"path/filepath"
)

func SetGroupRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "setgroups"))
	return string(buf)
}