package pid

import (
	"io/ioutil"
	"path/filepath"
)

func LoginUidRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "loginuid"))
	return string(buf)
}