// kernel/system statistics.  Varies with architecture.
package proc

import (
	"strconv"
	"strings"
	"io"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type StatInfo struct {
	CPUs    	[]CPU    // All CPUs info, CPUs[0] is total num, 1-len(CPUS)-1 is cores; CPUs: sum, cpu0, cpu1...
	PageIn  	uint64   // The number of pages the system paged in (from disk).
	PageOut 	uint64   // The number of pages the system that were paged out (from disk).
	SwapIn  	uint64   // The number of swap pages that have been brought in .
	SwapOut 	uint64   // The number of swap pages that have been brought out.
	Intr    	[]uint64 // This line shows counts of interrupts serviced since boot time, for each of the possible system interrupts.The first column is the total of all interrupts serviced including unnumbered architecture specific interrupts;  each  subsequent  column is the total for that particular numbered interrupt.  Unnumbered interrupts are not shown, only summed into the total.
	//disk_io	[]uint64	// (Linux 2.4 only), Not support here
	Ctxt         uint64 // The number of context switches that the system underwent.
	BootTime     uint64 // boot time, in seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC).
	Processes    uint64 // Number of processes in runnable state.  (Linux 2.5.45 onward.)
	ProcsRunning uint64 // Number of processes in runnable state.  (Linux 2.5.45 onward.)
	ProcsBlocked uint64 // Number of processes blocked waiting for I/O to complete.  (Linux 2.5.45 onward.)
}

type CPU struct {
	Name		string	// 0. CPU name: cpu, cpu0, cpu1
	User       	uint64 	// 1. Time spent in user mode.
	Nice       	uint64 	// 2. Time spent in user mode with low priority (nice).
	System     	uint64 	// 3. Time spent in system mode.
	Idle       	uint64 	// 4. Time spent in the idle task. This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file.
	Iowait     	uint64 	// 5. Time waiting for I/O to complete.(since Linux 2.5.41)
	Irq        	uint64 	// 6. Time servicing interrupts.(since Linux 2.6.0-test4)
	Softirq    	uint64 	// 7. Time servicing softirqs.(since Linux 2.6.0-test4)
	Steal      	uint64 	// 8. Stolen time, which is the time spent in other operating systems when running in a virtualized environment(since Linux 2.6.11)
	Guest      	uint64 	// 9. Time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.(since Linux 2.6.24)
	Guest_nice 	uint64 	// 10. Time spent running a niced guest (virtual CPU for guest operating systems under the  control  of the Linux kernel).(since Linux 2.6.33
}

func Stat() (StatInfo, error) {
	var stat StatInfo
	f, err := os.Open(filepath.Join(procfs, "stat"))
	if err != nil {
		return stat, fmt.Errorf("Stat Open %v/stat failed: %v\n", procfs, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var cpuLineNum int
	for {
		// read line bytes
		lb, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		// line string
		s := string(lb)
		vs := strings.Split(s,  " ")
		vs[len(vs)-1] = strings.TrimSpace(vs[len(vs)-1])
		if s[:3] == "cpu" {
			stat.CPUs = append(stat.CPUs, parseCPULine(vs))			
			cpuLineNum++
		} else if s[:4] == "page" {
			stat.PageIn, stat.PageOut = parsePageLine(vs)
		} else if s[:4] == "swap" {
			stat.SwapIn, stat.SwapOut = parseSwapLine(vs)
		} else if s[:4] == "intr" {
			stat.Intr = parseIntrLine(vs)
		} else if s[:4] == "ctxt" {
			stat.Ctxt, _ = strconv.ParseUint(vs[1], 10, 64)
		} else if s[:5] == "btime" {
			stat.BootTime, _ = strconv.ParseUint(vs[1], 10, 64)
		} else if s[:len("processes")] == "processes" {
			stat.Processes, _ = strconv.ParseUint(vs[1], 10, 64)
		} else if s[:len("procs_running")] == "procs_running" {
			stat.ProcsRunning, _ = strconv.ParseUint(vs[1], 10, 64)
		} else if s[:len("procs_blocked")] == "procs_blocked" {
			stat.ProcsBlocked, _ = strconv.ParseUint(vs[1], 10, 64)
		}
	}

	return stat, nil
}

func parseCPULine(vs []string) CPU {
	var c CPU
	for i := 0; i < len(vs); i++ {
		switch i {
		case 0:
			c.Name = vs[i]
		case 1:
			c.User, _ = strconv.ParseUint(vs[i], 10, 64)
		case 2:
			c.Nice, _ = strconv.ParseUint(vs[i], 10, 64)
		case 3:
			c.System, _ = strconv.ParseUint(vs[i], 10, 64)
		case 4:
			c.Idle , _ = strconv.ParseUint(vs[i], 10, 64)
		case 5:    
			c.Iowait, _ = strconv.ParseUint(vs[i], 10, 64)
		case 6:
			c.Irq, _ = strconv.ParseUint(vs[i], 10, 64)
		case 7:
			c.Softirq, _ = strconv.ParseUint(vs[i], 10, 64)
		case 8:
			c.Steal, _ = strconv.ParseUint(vs[i], 10, 64)
		case 9:
			c.Guest, _ = strconv.ParseUint(vs[i], 10, 64)
		case 10:
			c.Guest_nice, _ = strconv.ParseUint(vs[i], 10, 64)
		default:
		}
	}
	return c
}

func parsePageLine(vs []string) (in uint64, out uint64) {
	in, _ = strconv.ParseUint(vs[1], 10, 64)
	out, _ = strconv.ParseUint(vs[2], 10, 64)
	return in, out
}

func parseSwapLine(vs []string) (in uint64, out uint64) {
	in, _ = strconv.ParseUint(vs[1], 10, 64)
	out, _ = strconv.ParseUint(vs[2], 10, 64)
	return in, out
}

func parseIntrLine(vs []string) []uint64 {
	var intr []uint64
	for i := 1; i < len(vs); i++ {
		num, _ := strconv.ParseUint(vs[i], 10, 64)
		intr = append(intr, num)
	}

	return intr
}
