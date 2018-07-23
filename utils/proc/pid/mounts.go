package pid

import (
	"io/ioutil"
	"path/filepath"
)

func MountsRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "mounts"))
	return string(buf)
}