package pid

import (
	"io/ioutil"
	"path/filepath"
)

func MountStatsRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "mountstats"))
	return string(buf)
}