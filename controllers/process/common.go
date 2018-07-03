package process

import (
	"strings"
	"strconv"
	"fmt"
	"os"
)

type Process struct {
	Cmd			string	// 进程命令
	User		string
	
	CPU			float64	// 进程使用的 CPU 百分比
	CPUTime		string	// 进程使用的 CPU 时间
	
	MEM			float64 // 进程使用的内存百分比
	SZ			uint64	// 进程核心镜像的物理页面大小。包括 text，data 和 stack 的空间，不包括设备映射
	RSS			uint64	// 常驻内存大小，任务已使用的为交换物理内存，以 KB 为单位
	DRS			uint64	// 进程的数据集大小，专门用于非可执行代码的物理内存量
	TRS			uint64	// 进程的文本段占用的内存
	VSZ			uint64	// 进程的虚拟内存大小，以 KB 为单位。

	Pid			int		// 进程号
	Nlwp		int		// 进程内的线程数

	State		string	// 进程状态：D 不可中断; R 运行;	S 中断;	T 停止;	Z 僵死
	Nice		int		// 进程优先级
}

const psOutputFormatStr = "command,user,%cpu,cputime,%mem,sz,rss,drs,trs,vsz,pid,nlwp,state,nice"

func FormatProcStatus(cmdOutput []byte) []Process {
	var processes []Process

	for _, v := range(strings.Split(string(cmdOutput), "\n")) {
		if v == "" {continue}
		
		attrs := strings.Split(v, " ")

		if len(attrs) < 14 {
			fmt.Println("invalide length of attrs:", len(attrs), v)
			return processes
		}

		cmd := attrs[0]
		user := attrs[1]

		cpu, _ := strconv.ParseFloat(attrs[2], 64)
		
		cpuTime := attrs[3]
		
		mem, _ := strconv.ParseFloat(attrs[4], 64)
		sz, _ := strconv.ParseUint(attrs[5], 10, 64)
		rss, _ := strconv.ParseUint(attrs[6], 10, 64)
		drs, _ := strconv.ParseUint(attrs[7], 10, 64)
		trs, _ := strconv.ParseUint(attrs[8], 10, 64)
		vsz, _ := strconv.ParseUint(attrs[9], 10, 64)

		pid, _ := strconv.Atoi(attrs[10])
		nlwp, _ := strconv.Atoi(attrs[11])

		state := attrs[12]
		nice, _ := strconv.Atoi(attrs[13])

		processes = append(processes, Process{cmd, user, cpu, cpuTime, mem, sz, rss, drs, trs, vsz, pid, nlwp, state, nice})
	}

	return processes
}


func AllPids() []string {
	p, err := os.Open("/proc")
	if err != nil {
		fmt.Println("can not open /proc")
		return nil
	}

	defer p.Close()

	var files []string
	files, err = p.Readdirnames(0)
	if err != nil {
		fmt.Println("Readdirnames /proc failed")
		return nil
	}

	var pids []string
	for _, v := range(files) {
		if pid, _ := strconv.Atoi(v); pid != 0 {
			pids = append(pids, v)
		}
	}
	return pids
}