package pid

import (
	"io/ioutil"
	"path/filepath"
)

func MountsRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "mounts"))
	return string(buf)
}