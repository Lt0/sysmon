// This file displays the current score that the kernel gives to this process for the purpose of selecting a process
// for the OOM-killer.  A higher score means that the process is more likely to be selected by the OOM-killer.   The
// basis for this score is the amount of memory used by the process, with increases (+) or decreases (-) for factors
// including:
// 
// * whether the process creates a lot of children using fork(2) (+);
// 
// * whether the process has been running a long time, or has used a lot of CPU time (-);
// 
// * whether the process has a low nice value (i.e., > 0) (+);
// 
// * whether the process is privileged (-); and
// 
// * whether the process is making direct hardware access (-).
// 
// The oom_score also reflects the adjustment specified by the oom_score_adj or oom_adj setting for the process.


// 此文件显示内核为此进程提供的当前分数，以便为OOM杀手选择进程。 得分越高意味着OOM杀手更有可能选择该过程。 此分数的取决于进程运行过程使用的内存量，决策因素如下：
// 1. 该进程是否使用fork（2）（+）创建了很多子进程;
// 2. 进程是否已运行很长时间，或者是否占用了大量CPU时间（ - ）;
// 3. 该进程是否具有低的 nice 值（即> 0）（+）;
// 4. 流程是否具有特权（ - ）; 和
// 5. 该进程是否进行直接硬件访问（ - ）。
// oom_score还反映了进程的oom_score_adj或oom_adj设置指定的调整。

package pid

import (
	"strconv"
	"path/filepath"
	"io/ioutil"
)

func OOMScore(pid string) int {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "oom_score"))
	score, _ := strconv.Atoi(string(buf))
	return score
}