package pid

import (
	"io/ioutil"
	"path/filepath"
)

func ProjidMapRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "projid_map"))
	return string(buf)
}