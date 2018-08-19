// This file can be used to adjust the badness heuristic used to select which process gets killed in out-of-memory conditions.
// 
// The badness heuristic assigns a value to each candidate task ranging from 0 (never kill) to 1000 (always kill) to
// determine which process is targeted.  The units are roughly a proportion along that range of allowed  memory  the
// process  may allocate from, based on an estimation of its current memory and swap use.  For example, if a task is
// using all allowed memory, its badness score will be 1000.  If it is using half of its allowed memory,  its  score
// will be 500.
// 
// There  is an additional factor included in the badness score: root processes are given 3% extra memory over other
// tasks.
// 
// The amount of "allowed" memory depends on the context in which the OOM-killer was called.  If it is  due  to  the
// memory  assigned  to  the allocating task's cpuset being exhausted, the allowed memory represents the set of mems
// assigned to that cpuset (see cpuset(7)).  If it is due to a mempolicy's node(s) being exhausted, the allowed mem‐
// ory  represents  the  set  of mempolicy nodes.  If it is due to a memory limit (or swap limit) being reached, the
// allowed memory is that configured limit.  Finally, if it is due to the entire system being  out  of  memory,  the
// allowed memory represents all allocatable resources.
// 
// The  value  of  oom_score_adj  is  added  to the badness score before it is used to determine which task to kill.
// Acceptable values range from -1000 (OOM_SCORE_ADJ_MIN) to +1000 (OOM_SCORE_ADJ_MAX).  This allows user  space  to
// control  the preference for OOM-killing, ranging from always preferring a certain task or completely disabling it
// from OOM-killing.  The lowest possible value, -1000, is equivalent to disabling  OOM-killing  entirely  for  that
// task, since it will always report a badness score of 0.
// 
// Consequently, it is very simple for user space to define the amount of memory to consider for each task.  Setting
// a oom_score_adj value of +500, for example, is roughly equivalent to allowing the remainder of tasks sharing  the
// same system, cpuset, mempolicy, or memory controller resources to use at least 50% more memory.  A value of -500,
// on the other hand, would be roughly equivalent to discounting 50% of the task's allowed memory from being consid‐
// ered as scoring against the task.
// 
// For  backward  compatibility  with  previous  kernels,  /proc/[pid]/oom_adj can still be used to tune the badness
// score.  Its value is scaled linearly with oom_score_adj.
// 
// Writing to /proc/[pid]/oom_score_adj or /proc/[pid]/oom_adj will change the other with its scaled value.

// 本文件可用于调整进程在 out-of-memory 的情况下被杀死的权重
// 该文件的有效值为 0(任何情况下都不 kill 该进程)~1000(一旦内存不足立即 kill 该进程)，
package pid

import (
	"strconv"
	"path/filepath"
	"io/ioutil"
)

func OOMScoreAdj(pid string) int {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "oom_score_adj"))
	score, _ := strconv.Atoi(string(buf))
	return score
}