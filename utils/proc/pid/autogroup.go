package pid

import (
	"io/ioutil"
	"path/filepath"
)

func AutoGroupRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "autogroup"))
	return string(buf)
}