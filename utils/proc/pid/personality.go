// 该只读文件公开进程的执行域，由 personality(2) 设置。 该值以十六进制表示法显示。
// 访问该文件的权限由 ptrace 访问模式 PTRACE_MODE_ATTACH_FSCREDS 检查控制; 见ptrace（2）。
package pid

import (
	"io/ioutil"
	"path/filepath"
)

func PersonalityRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "personality"))
	return string(buf)
}