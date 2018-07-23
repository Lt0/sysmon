package pid

import (
	"io/ioutil"
	"path/filepath"
)

func LoginUidRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "loginuid"))
	return string(buf)
}