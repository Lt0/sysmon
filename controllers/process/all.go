package process

import (
	"bufio"
	"path/filepath"
	"strconv"
	// "unicode"
	"fmt"
	"os/exec"
	// "strings"
	// "strconv"
	// "path/filepath"
	"os"
	"io"
	// "time"
	
	"github.com/astaxie/beego"
)

type AllProcessCtrl struct {
	Controller *beego.Controller
}

type AllProcess struct {
	Processes 	[]Process
}

func (p *AllProcessCtrl) Do() interface{} {
	var ap AllProcess

	fmt.Println("do AllProcess")
	p.All(&ap)

	AllPidInfo(AllPids())

	return ap
}

func (p *AllProcessCtrl) All(ap *AllProcess) {
	filter := `tail -n +2 | sed 's/^\s*//g' | sed 's/\s\+/ /g'`
	cmd := fmt.Sprintf("ps caxo %s | %s", psOutputFormatStr, filter)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("AllProcessCtrl", err)
		return
	}
	ap.Processes = FormatProcStatus(out)
}

type PidStat struct {
	Pid 		string		// 进程号			stat
	// Tasks		[]pidInfo	// 所有线程信息
	Nlwp		string		// 线程数			stat
	State		string		// 进程状态			stat
	Priority	string		// 动态优先级		stat
	Nice		string		// 静态优先级		stat
	Comm		string		// 可执行文件名		comm, stat
	Cmdline		string		// 可执行文件路径	cmdline
	// User		string
	TaskCPU		string		// 运行在哪个 CPU 上	stat

	StartTime	string		// 系统开机后该进程启动的时间，单位为jiffies	stat
	CPU			uint64		// cpu jiffies		stat
	// CpuTime		time.Time

	RSS			string		// 真正具有数据页的物理内存（包含进程用到的所有共享库占用的全部内存）			stat
	// PSS			uint64		// 进程占用的实际物理内存大小，对于该进程用到的共享库，会根据使用该库的进程数量，按比例显示该进程占用的内存
	// USS			uint64		// 进程独占的实际分配的私有内存。
	// SHR			uint64		// 进程调用的共享库占用的内存。 
	VSS			string		// 进程可访问的全部地址空间			stat
}

func AllPidInfo(pids []string) {
	var all []PidStat
	for _, v := range(pids) {
		all = append(all, PidInfo(v))
	}
	fmt.Println("all:", all)
}

func PidInfo(pid string) PidStat {
	var info PidStat

	// fill Pid, Comm, State, CPU, Priority, Nice, Nlwp, StartTime, VSS, RSS, TaskCPU by /proc/$PID/stat
	f, err := os.Open(filepath.Join("/proc", pid, "stat"))
	if err != nil {
		fmt.Printf("open /proc/%s/stat failed\n", pid)
		return info
	}
	defer f.Close()

	statReader := bufio.NewReader(f)
	for i := 0; ; i++ {
		b, e := statReader.ReadBytes(' ')
		if e == io.EOF {
			break;
		}
		s := string(b)

		switch i {
		case 0:
			info.Pid = s
		case 1:
			info.Comm = s[1:len(s)-2]
		case 2:
			info.State = s
		case 13, 14, 15, 16:
			n, _ := strconv.ParseUint(s, 10, 64)
			info.CPU += n
		case 17:
			info.Priority = s
		case 18:
			info.Nice = s
		case 19:
			info.Nlwp = s
		case 21:
			info.StartTime = s
		case 22:
			info.VSS = s
		case 23:
			info.RSS = s
		case 38:
			info.TaskCPU = s
		default:
		}
	}


	// fill cmdline by /proc/$PID/cmdline
	cmdlineFile, err := os.Open(filepath.Join("/proc", pid, "cmdline"))
	if err != nil {
		fmt.Printf("open /proc/%s/cmdline failed\n", pid)
		return info
	}
	defer cmdlineFile.Close()

	cmdLineReader := bufio.NewReader(cmdlineFile)
	cmdline, _ := cmdLineReader.ReadBytes('\n')
	info.Cmdline = string(cmdline)


	
	return info
}