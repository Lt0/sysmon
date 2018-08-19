// allow: 允许包含进程pid的用户命名空间中的进程使用setgroups（2）系统调用
// deny: 在该用户命名空间中不允许setgroups（2）
// 请注意，如果没有设置 /proc/[pid]/gid_map，则无论该文件的值是什么，都不允许调用setgroups（2）
package pid

import (
	"io/ioutil"
	"path/filepath"
)

func SetGroupRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "setgroups"))
	return string(buf)
}