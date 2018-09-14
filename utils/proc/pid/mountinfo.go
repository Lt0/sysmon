package pid

import (
	"io/ioutil"
	"path/filepath"
)

func MountInfoRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "mountinfo"))
	return string(buf)
}