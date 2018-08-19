package pid

import (
	"io/ioutil"
	"path/filepath"
)

func SyscallRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "syscall"))
	return string(buf)
}