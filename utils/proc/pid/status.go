package pid

import (
	"strconv"
	"strings"
	"io"
	"bufio"
	"fmt"
	"path/filepath"
	"os"
)

type StatusInfo struct{
	Name					string		// Command run by this process.
	State					string		// Current  state of the process.  One of "R (running)", "S (sleeping)", "D (disk sleep)", "T (stopped)", "T (tracing stop)", "Z (zombie)", or "X (dead)".
	Tgid					uint64		// Thread group ID (i.e., Process ID).
	Ngid					uint64
	Pid						uint64		// Thread ID (see gettid(2)).
	PPid					uint64		// PID of parent process.
	TracerPid				[]uint64	// PID of process tracing this process (0 if not being traced).
	Uid						[]uint64	// Real, effective, saved set, and filesystem UIDs (GIDs).
	Gid						[]uint64	// Real, effective, saved set, and filesystem UIDs (GIDs).
	FDSize					uint64		// Number of file descriptor slots currently allocated.
	Groups					[]uint64	// Supplementary group list.
	NStgid					uint64		// 
	NSpid					uint64		// 
	NSpgid					uint64		// 
	NSsid					uint64		// 
	VmPeak					uint64		// Peak virtual memory size.
	VmSize					uint64		// Virtual memory size.
	VmLck					uint64		// Locked memory size (see mlock(3)).
	VmPin					uint64		// Pinned memory size (since Linux 3.2). These are pages that can't be moved because something needs to directly access physical memory.
	VmHWM					uint64		// Peak resident set size ("high water mark").
	VmRSS					uint64		// Resident set size.
	VmData					uint64		// Size of data
	VmStk					uint64		// Size of stack
	VmExe					uint64		// Size of text segments.
	VmLib					uint64		// Shared library code size.
	VmPMD					uint64		// Size of second-level page tables (since Linux 4.0).
	VmPTE					uint64		// Page table entries size (since Linux 2.6.10).
	VmSwap					uint64		// Swapped-out virtual memory size by anonymous private pages; shmem swap usage is not included (since Linux 2.6.34).
	HugetlbPages			uint64		// 
	Threads					uint64		// Number of threads in process containing this thread.
	SigQ					[]uint64	// This field contains two slash-separated numbers that relate to queued signals for the real user ID of  this process.   The  first of these is the number of currently queued signals for this real user ID, and the second is the resource limit on the number of queued signals for this process (see the description of RLIMIT_SIGPENDING  in getrlimit(2)).
	SigPnd					string		// Number of signals pending for thread and for process as a whole (see pthreads(7) and signal(7)).
	ShdPnd					string		// Number of signals pending for process as a whole (see pthreads(7) and signal(7)).
	SigBlk					string		// Masks indicating signals being blocked (see signal(7)).
	SigIgn					string		// Masks indicating signals being ignored (see signal(7)).
	SigCgt					string		// Masks indicating signals being caught (see signal(7)).
	CapInh					string		// Masks of capabilities enabled in inheritable (see capabil‐ ities(7)).
	CapPrm					string		// Masks of capabilities enabled in permitted (see capabil‐ ities(7)).
	CapEff					string		// Masks of capabilities enabled in effective sets (see capabil‐ ities(7)).
	CapBnd					string		// Capability Bounding set (since Linux 2.6.26, see capabilities(7)).
	CapAmb					string		// Ambient capability set (since Linux 4.3, see capabilities(7)).
	Seccomp					uint64		// Seccomp mode of the process (since Linux 3.8, see seccomp(2)).  0 means SECCOMP_MODE_DISABLED;  1  means SECCOMP_MODE_STRICT;  2  means SECCOMP_MODE_FILTER.  This field is provided only if the kernel was built with the CONFIG_SECCOMP kernel configuration option enabled.
	CpusAllowed				string		// Mask of CPUs on which this process may run (since Linux 2.6.24, see cpuset(7)).
	CpusAllowedList			string		// Same as previous, but in "list format" (since Linux 2.6.26, see cpuset(7)).
	MemsAllowed				string		// Mask of memory nodes allowed to this process (since Linux 2.6.24, see cpuset(7)).
	MemsAllowedList			string		// Same as previous, but in "list format" (since Linux 2.6.26, see cpuset(7)).
	VoluntaryCtxtSwitches	uint64		// Number of voluntary context switches (since Linux 2.6.23).
	NonVoluntarCtxtSwitches	uint64		// Number of involuntary context switches (since Linux 2.6.23).
}

func Status(pid string) (StatusInfo, error) {
	var status StatusInfo
	f, err := os.Open(filepath.Join(procfs, pid, "status"))
	if err != nil {
		return status, fmt.Errorf("Pid Status: open %v/%v/status failed: %v\n", procfs, pid, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// get type bytes
		tb, err := r.ReadBytes(':')
		if err == io.EOF {
			break
		}

		// get type string
		ts := string(tb[:len(tb)-1])

		// get type bytes
		vb, _ := r.ReadBytes('\n')

		// get type string
		vs := strings.TrimSpace(string(vb))
		vs = strings.TrimSpace(vs)

		switch ts {
		case "Name":
			status.Name = vs
		case "State":
			status.State = vs
		case "Tgid":
			status.Tgid, _ = strconv.ParseUint(vs, 10, 64)
		case "Ngid":
			status.Ngid, _ = strconv.ParseUint(vs, 10, 64)
		case "Pid":
			status.Pid, _ = strconv.ParseUint(vs, 10, 64)
		case "PPid":
			status.PPid, _ = strconv.ParseUint(vs, 10, 64)
		case "TracerPid":
			ss := strings.Split(vs, "\t")
			for _, v := range(ss) {
				vv, _ := strconv.ParseUint(v, 10, 64)
				status.TracerPid = append(status.TracerPid, vv)
			}
		case "Uid":
			ss := strings.Split(vs, "\t")
			for _, v := range(ss) {
				vv, _ := strconv.ParseUint(v, 10, 64)
				status.Uid = append(status.Uid, vv)
			}
		case "Gid":
			ss := strings.Split(vs, "\t")
			for _, v := range(ss) {
				vv, _ := strconv.ParseUint(v, 10, 64)
				status.Gid = append(status.Gid, vv)
			}
		case "FDSize":
			status.FDSize, _ = strconv.ParseUint(vs, 10, 64)
		case "Groups":
			ss := strings.Split(vs, "\t")
			for _, v := range(ss) {
				vv, _ := strconv.ParseUint(v, 10, 64)
				status.Groups = append(status.Groups, vv)
			}
		case "NStgid":
			status.NStgid, _ = strconv.ParseUint(vs, 10, 64)
		case "NSpid":
			status.NSpid, _ = strconv.ParseUint(vs, 10, 64)
		case "NSpgid":
			status.NSpgid, _ = strconv.ParseUint(vs, 10, 64)
		case "NSsid":
			status.NSsid, _ = strconv.ParseUint(vs, 10, 64)
		case "VmPeak":
			status.VmPeak, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmSize":
			status.VmSize, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmLck":
			status.VmLck, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmPin":
			status.VmPin, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmHWM":
			status.VmHWM, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmRSS":
			status.VmRSS, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmData":
			status.VmData, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmStk":
			status.VmStk, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmExe":
			status.VmExe, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmLib":
			status.VmLib, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmPMD":
			status.VmPMD, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmPTE":
			status.VmPTE, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "VmSwap":
			status.VmSwap, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "HugetlbPages":
			status.HugetlbPages, _ = strconv.ParseUint(vs[:len(vs)-3], 10, 64)
		case "Threads":
			status.Threads, _ = strconv.ParseUint(vs, 10, 64)
		case "SigQ":
			ss := strings.Split(vs, "/")
			for _, v := range(ss) {
				vv, _ := strconv.ParseUint(v, 10, 64)
				status.SigQ = append(status.SigQ, vv)
			}
		case "SigPnd":
			status.SigPnd = vs
		case "ShdPnd":
			status.ShdPnd = vs
		case "SigBlk":
			status.SigBlk = vs
		case "SigIgn":
			status.SigIgn = vs
		case "SigCgt":
			status.SigCgt = vs
		case "CapInh":
			status.CapInh = vs
		case "CapPrm":
			status.CapPrm = vs
		case "CapEff":
			status.CapEff = vs
		case "CapBnd":
			status.CapBnd = vs
		case "CapAmb":
			status.CapAmb = vs
		case "Seccomp":
			status.Seccomp, _ = strconv.ParseUint(vs, 10, 64)
		case "Cpus_allowed":
			status.CpusAllowed = vs
		case "Cpus_allowed_list":
			status.CpusAllowedList = vs
		case "Mems_allowed":
			status.MemsAllowed = vs
		case "Mems_allowed_list":
			status.MemsAllowedList = vs
		case "voluntary_ctxt_switches":
			status.VoluntaryCtxtSwitches, _ = strconv.ParseUint(vs, 10, 64)
		case "nonvoluntary_ctxt_switches":
			status.NonVoluntarCtxtSwitches, _ = strconv.ParseUint(vs, 10, 64)
		default:
			fmt.Println("unknow pattern: ", ts)
		}
	}

	return status, nil
}