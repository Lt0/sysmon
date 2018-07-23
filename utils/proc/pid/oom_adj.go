// This file can be used to adjust the score used to select which process should be killed in an out-of-memory (OOM) situation. 
// The kernel uses this value for a bit-shift operation of the process's oom_score value: 
// valid values are in the range -16 to +15, plus the special value -17, which disables OOM-killing altogether for this process. 
// A positive score increases the likelihood of this process being killed by the OOM-killer; a negative score decreases the likelihood.
// The default value for this file is 0; a new process inherits its parent's oom_adj setting. A process must be privileged (CAP_SYS_RESOURCE) to update this file.
// Since Linux 2.6.36, use of this file is deprecated in favor of /proc/[pid]/oom_score_adj.

// 在内存不足（OOM）的情况下，系统根据各个进程的该文件的值来权衡要杀死的进程。 
// 内核使用这个进行进程的 oom_score 值的位移操作：有效值在 -16 到 +15 的范围内，再加上特殊值 -17，它会为此进程完全禁用OOM查杀。 
// 正分会增加这个过程被 OOM 杀手杀死的可能性; 负分数降低了可能性。
// 该文件的默认值为0; 新进程继承其父进程的oom_adj设置。 进程必须具有特权（CAP_SYS_RESOURCE）才能更新此文件。
// 从Linux 2.6.36开始，不推荐使用此文件，而使用/ proc / [pid] / oom_score_adj。
package pid

import (
	"strconv"
	"path/filepath"
	"io/ioutil"
)

func OOMAdj(pid string) int {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "oom_adj"))
	score, _ := strconv.Atoi(string(buf))
	return score
}