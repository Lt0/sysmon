package pid

import (
	"io/ioutil"
	"path/filepath"
)

func CGroupRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "cgroup"))
	return string(buf)
}