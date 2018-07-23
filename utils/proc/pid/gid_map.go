// http://man7.org/linux/man-pages/man7/user_namespaces.7.html
package pid

import (
	"io/ioutil"
	"path/filepath"
)

func GidMapRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "gid_map"))
	return string(buf)
}