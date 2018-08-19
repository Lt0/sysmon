package pid

import (
	"io/ioutil"
	"path/filepath"
)

func SessionIDRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "sessionid"))
	return string(buf)
}