package pid

import (
	"io/ioutil"
	"path/filepath"
)

func ProjidMapRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "projid_map"))
	return string(buf)
}