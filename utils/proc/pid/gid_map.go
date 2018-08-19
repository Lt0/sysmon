// http://man7.org/linux/man-pages/man7/user_namespaces.7.html
// 除了把 uid 改为 gid，所有地方和 uid_map 的相同
package pid

import (
	"io/ioutil"
	"path/filepath"
)

func GidMapRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "gid_map"))
	return string(buf)
}