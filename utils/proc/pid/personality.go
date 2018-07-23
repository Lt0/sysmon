package pid

import (
	"io/ioutil"
	"path/filepath"
)

func PersonalityRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "personality"))
	return string(buf)
}