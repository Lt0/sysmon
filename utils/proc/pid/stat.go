package pid

import (
	"strings"
	"bufio"
	"fmt"
	"path/filepath"
	"os"
	"io"
	"strconv"
)

type PidStatInfo struct {
	Pid					int			// 1. The process ID.
	Comm				string		// 2. The filename of the executable, in parentheses.  This is visible whether or not the executable is swapped out.
	State				string		// 3. One of the following characters, indicating process state: R/S/D/Z/T/t/W/X/x/K/W/P
	PPid				int			// 4. The PID of the parent of this process.
	PGrp				int			// 5. The process group ID of the process.
	Session				int			// 6. The session ID of the process
	TtyNr				int			// 7. The controlling terminal of the process.  (The minor device number is contained  in  the  combination of bits 31 to 20 and 7 to 0; the major device number is in bits 15 to 8.)
	TPGid				int			// 8. The ID of the foreground process group of the controlling terminal of the process.
	Flags				uint64		// 9. The kernel flags word of the process. 
	Minflt				uint64		// 10. The number of minor faults the process has made which have not required loading a memory page from disk.
	CMinflt				uint64		// 11. The number of minor faults that the process's waited-for children have made
	Majflt				uint64		// 12. The number of major faults the process has made which have required loading a memory page from disk.
	CMajflt				uint64		// 13. The number of major faults that the process's waited-for children have made.
	UTime				uint64		// 14. Amount of time that this process has been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	STime				uint64		// 15. Amount of time that this process has been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	CUTime				int64		// 16. Amount of time that this process's waited-for children have been scheduled in user mode, measured in clock ticks (divide  by  sysconf(_SC_CLK_TCK)).
	CSTime				int64		// 17. Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Priority			int64		// 18. (Explanation for Linux 2.6) For processes running a real-time scheduling policy
	Nice				int64		// 19. The nice value (see setpriority(2)), a value in the range 19 (low priority) to -20 (high priority).
	NumThreads			int64		// 20. Number of threads in this process (since Linux 2.6).
	Itrealvalue			int64		// 21. The time in jiffies before the next SIGALRM is sent to the process due to an interval timer.
	StartTime			uint64		// 22. The time the process started after system boot. In kernels before Linux 2.6, this value was expressed in jiffies. Since Linux 2.6, the value is expressed in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	VSize				uint64		// 23. Virtual memory size in bytes.
	RSS					int64		// 24. Resident Set Size: number of pages the process has in real memory.
	RSSLim				uint64		// 25. Current soft limit in bytes on the rss of the process; see the description of RLIMIT_RSS in getrlimit(2).
	Startcode			uint64		// 26. The address above which program text can run.
	EndCode				uint64		// 27. The address below which program text can run.
	StartStack			uint64		// 28. The address of the start (i.e., bottom) of the stack.
	KStkESP				uint64		// 29. The current value of ESP (stack pointer), as found in the kernel stack page for the process.
	KStkEIP				uint64		// 30. The current EIP (instruction pointer).
	Signal				uint64		// 31. The bitmap of pending signals, displayed as a decimal number. Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.
	Blocked				uint64		// 32. The bitmap of blocked signals, displayed as a decimal number. Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.
	SigIgnore			uint64		// 33. The bitmap of ignored signals, displayed as a decimal number. Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.
	SigCatch			uint64		// 34. The bitmap of caught signals, displayed as a decimal number. Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.
	WChan				uint64		// 35. This  is  the  "channel" in which the process is waiting. It is the address of a location in the kernel where the process is sleeping.  The corresponding symbolic name can be found in /proc/[pid]/wchan.
	NSwap				uint64		// 36. Number of pages swapped (not maintained).
	CNSwap				uint64		// 37. Cumulative nswap for child processes (not maintained).
	ExitSignal			int			// 38. Signal to be sent to parent when we die.
	Processor			int			// 39. CPU number last executed on.
	RTPriority			uint64		// 40. Real-time scheduling priority, a number in the range 1 to 99 for processes scheduled  under  a  real-time policy, or 0, for non-real-time processes (see sched_setscheduler(2)).
	Policy				uint64		// 41. Scheduling policy (see sched_setscheduler(2)).
	DelayacctBlkioTicks	uint64		// 42. Aggregated block I/O delays, measured in clock ticks (centiseconds).
	GuestTime			uint64		// 43. Guest time of the process (time spent running a virtual CPU for a guest operating system), measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	CGuestTime			uint64		// 44. Guest time of the process's children, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	StartData			uint64		// 45. Address above which program initialized and uninitialized (BSS) data are placed.
	EndData				uint64		// 46. Address below which program initialized and uninitialized (BSS) data are placed.
	StartBrk			uint64		// 47. Address above which program heap can be expanded with brk(2).
	ArgStart			uint64		// 48. Address above which program command-line arguments (argv) are placed.
	ArgEnd				uint64		// 49. Address below program command-line arguments (argv) are placed.
	EnvStart			uint64		// 50. Address above which program environment is placed.
	EnvEnd				uint64		// 51. Address below which program environment is placed.
	ExitCode			int			// 52. The thread's exit status in the form reported by waitpid(2).
}

func Stat(pid string) (PidStatInfo, error) {
	var stat PidStatInfo
	f, err := os.Open(filepath.Join(procfs, pid, "stat"))
	if err != nil {
		return stat, fmt.Errorf("GetStat: open %v/%s/pid failed, err: %v\n", procfs, pid, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for i := 1; ; i++ {
		b, err := r.ReadBytes(' ')
		if err == io.EOF {
			break
		}
		
		s := string(b)
		s = strings.TrimSpace(s)

		switch i {
		case 1:
			stat.Pid, _ = strconv.Atoi(s)
		case 2:
			stat.Comm = s
		case 3:
			stat.State = s
		case 4:
			stat.PPid, _ = strconv.Atoi(s)
		case 5:
			stat.PGrp, _ = strconv.Atoi(s)
		case 6:
			stat.Session, _ = strconv.Atoi(s)
		case 7:
			stat.TtyNr, _ = strconv.Atoi(s)
		case 8:
			stat.TPGid, _ = strconv.Atoi(s)
		case 9:
			stat.Flags, _ = strconv.ParseUint(s, 10, 64)
		case 10:
			stat.Minflt, _ = strconv.ParseUint(s, 10, 64)
		case 11:
			stat.CMinflt, _ = strconv.ParseUint(s, 10, 64)
		case 12:
			stat.Majflt, _ = strconv.ParseUint(s, 10, 64)
		case 13:
			stat.CMajflt, _ = strconv.ParseUint(s, 10, 64)
		case 14:
			stat.UTime, _ = strconv.ParseUint(s, 10, 64)
		case 15:
			stat.STime, _ = strconv.ParseUint(s, 10, 64)
		case 16:
			stat.CUTime, _ = strconv.ParseInt(s, 10, 64)
		case 17:
			stat.CSTime, _ = strconv.ParseInt(s, 10, 64)
		case 18:
			stat.Priority, _ = strconv.ParseInt(s, 10, 64)
		case 19:
			stat.Nice, _ = strconv.ParseInt(s, 10, 64)
		case 20:
			stat.NumThreads, _ = strconv.ParseInt(s, 10, 64)
		case 21:
			stat.Itrealvalue, _ = strconv.ParseInt(s, 10, 64)
		case 22:
			stat.StartTime, _ = strconv.ParseUint(s, 10, 64)
		case 23:
			stat.VSize, _ = strconv.ParseUint(s, 10, 64)
		case 24:
			stat.RSS, _ = strconv.ParseInt(s, 10, 64)
		case 25:
			stat.RSSLim, _ = strconv.ParseUint(s, 10, 64)
		case 26:
			stat.Startcode, _ = strconv.ParseUint(s, 10, 64)
		case 27:
			stat.EndCode, _ = strconv.ParseUint(s, 10, 64)
		case 28:
			stat.StartStack, _ = strconv.ParseUint(s, 10, 64)
		case 29:
			stat.KStkESP, _ = strconv.ParseUint(s, 10, 64)
		case 30:
			stat.KStkEIP, _ = strconv.ParseUint(s, 10, 64)
		case 31:
			stat.Signal, _ = strconv.ParseUint(s, 10, 64)
		case 32:
			stat.Blocked, _ = strconv.ParseUint(s, 10, 64)
		case 33:
			stat.SigIgnore, _ = strconv.ParseUint(s, 10, 64)
		case 34:
			stat.SigCatch, _ = strconv.ParseUint(s, 10, 64)
		case 35:
			stat.WChan, _ = strconv.ParseUint(s, 10, 64)
		case 36:
			stat.NSwap, _ = strconv.ParseUint(s, 10, 64)
		case 37:
			stat.CNSwap, _ = strconv.ParseUint(s, 10, 64)
		case 38:
			stat.ExitSignal, _ = strconv.Atoi(s)
		case 39:
			stat.Processor, _ = strconv.Atoi(s)
		case 40:
			stat.RTPriority, _ = strconv.ParseUint(s, 10, 64)
		case 41:
			stat.Policy, _ = strconv.ParseUint(s, 10, 64)
		case 42:
			stat.DelayacctBlkioTicks, _ = strconv.ParseUint(s, 10, 64)
		case 43:
			stat.GuestTime, _ = strconv.ParseUint(s, 10, 64)
		case 44:
			stat.CGuestTime, _ = strconv.ParseUint(s, 10, 64)
		case 45:
			stat.StartData, _ = strconv.ParseUint(s, 10, 64)
		case 46:
			stat.EndData, _ = strconv.ParseUint(s, 10, 64)
		case 47:
			stat.StartBrk, _ = strconv.ParseUint(s, 10, 64)
		case 48:
			stat.ArgStart, _ = strconv.ParseUint(s, 10, 64)
		case 49:
			stat.ArgEnd, _ = strconv.ParseUint(s, 10, 64)
		case 50:
			stat.EnvStart, _ = strconv.ParseUint(s, 10, 64)
		case 51:
			stat.EnvEnd, _ = strconv.ParseUint(s, 10, 64)
		case 52:
			stat.ExitCode, _ = strconv.Atoi(s)
		default:
			fmt.Printf("unknon info: index: %d, value: %v\n", i, s)
		}
	}

	return stat, nil
}