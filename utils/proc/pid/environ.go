package pid

import (
	"io/ioutil"
	"path/filepath"
	"bytes"
)

func EnvironRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "environ"))
	buf = bytes.Replace(buf, []byte{0}, []byte{'\n'}, -1)
	return string(buf)
}