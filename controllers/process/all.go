package process

import (
	"strings"
	"bufio"
	"path/filepath"
	"strconv"
	"fmt"
	"os"
	"os/user"
	"io"

	"github.com/Lt0/sysmon/utils/proc/pid"
	"github.com/astaxie/beego"
)

type AllProcessCtrl struct {
	Controller *beego.Controller
}


type Process struct {
	Pid 		string		// 进程号			stat
	Task		[]string	// 所有线程
	Nlwp		int64		// 线程数			stat
	State		string		// 进程状态			stat
	Priority	int64		// 动态优先级		stat
	Nice		int64		// 静态优先级		stat
	Comm		string		// 可执行文件名		comm, stat
	Cmdline		string		// 可执行文件路径	cmdline
	Uid			string
	User		string
	TaskCPU		string		// 运行在哪个 CPU 上	stat

	StartTime	uint64		// 系统开机后该进程启动的时间，单位为jiffies	stat
	CPU			uint64		// cpu jiffies		stat
	// CpuTime		time.Time

	// SZ			string		// 进程占用的总内存，单位 page
	// RSS			string		// 真正具有数据页的物理内存（包含进程用到的所有共享库占用的全部内存）, byte			stat
	// PSS			uint64		// 进程占用的实际物理内存大小，对于该进程用到的共享库，会根据使用该库的进程数量，按比例显示该进程占用的内存
	// USS			uint64		// 进程独占的实际分配的私有内存。
	// SHR			uint64		// 进程调用的共享库占用的内存。 
	// VSS			string		// 进程可访问的全部地址空间, byte			stat


	// 以下和内存相关的数据从 /proc/$PID/status 文件中获取，单位为 KB，字符串中不带单位
	VmSize		uint64		// 虚拟地址空间大小，
	// VmLck		string		// 已经锁住的物理内存的大小。锁住的物理内存不能交换到硬盘 
	// VmPin		string		// 固定的内存大小
	// VmHWM		string		// RSS 峰值大小（高峰）
	VmRSS		uint64		// 内存部分的大小。 它包含以下三个部分（VmRSS = RssAnon + RssFile + RssShmem）
	// VmData		string		// 私有数据段的大小
	// VmStk		string		// 堆栈段的大小
	// VmExe		string		// 文本段的大小
	// VmLib		string		// 共享库代码的大小
	VmPTE		uint64		// 该进程的所有页表的大小
	// VmPMD		string		// 二级页表大小
	VmSwap		uint64		// 被交换到交换分区的匿名数据大小
}

type AllProcess struct {
	Processes 	[]Process
}

func (p *AllProcessCtrl) Do() interface{} {
	var ap AllProcess

	fmt.Println("do AllProcess")
	p.All(&ap)

	return ap
}

func (p *AllProcessCtrl) All(ap *AllProcess) {
	for _, v := range(AllPids()) {
		info, err := PidInfo(v)
		if err != nil {
			fmt.Printf("not add %v cause: %v\n", v, err)
			continue
		}
		ap.Processes = append(ap.Processes, info)
	}
}

func PidInfo(pidStr string) (Process, error) {
	var info Process

	info.Pid = pidStr

	// Fill Pid, Comm, State, CPU, Priority, Nice, Nlwp, StartTime, VSS, RSS, TaskCPU by /proc/$PID/stat
	stat, _ := pid.Stat(pidStr)
	if len(stat.Comm) > 2 {
		info.Comm = stat.Comm[1:len(stat.Comm)-2]
	} else {
		info.Comm = stat.Comm
	}
	info.State = stat.State
	info.CPU = uint64(stat.UTime) + uint64(stat.STime) + uint64(stat.CUTime) + uint64(stat.CSTime)
	info.Priority = stat.Priority
	info.Nice = stat.Nice
	info.Nlwp = stat.NumThreads
	info.StartTime = stat.StartTime
	info.TaskCPU = "cpu" + strconv.Itoa(stat.Processor)

	
	// Fill Uid, VmPTE, VmSwap from /proc/$PID/status
	statusFile, err := os.Open(filepath.Join("/proc", pidStr, "status"))
	if err != nil {
		// fmt.Printf("open /proc/%s/status failed\n", pidStr)
		return info, fmt.Errorf("open /proc/%s/status failed\n", pidStr)
	}
	defer statusFile.Close()

	statusReader := bufio.NewReader(statusFile)
	for {
		b1, e := statusReader.ReadBytes(':')
		
		if e == io.EOF {
			break;
		}
		t := string(b1[:len(b1)-1])

		contentBytes, _ := statusReader.ReadBytes('\n')
		content := string(contentBytes)
		
		switch t {
		case "Uid":
			ids := strings.Split(content, "\t")
			info.Uid = ids[1]
			ui, _ := user.LookupId(info.Uid)
			info.User = ui.Username
		case "VmSize":
			str := strings.TrimSpace(content)
			info.VmSize, _ = strconv.ParseUint(str[:len(str)-3], 10, 64)
		case "VmRSS":
			str := strings.TrimSpace(content)
			info.VmRSS, _ = strconv.ParseUint(str[:len(str)-3], 10, 64)
		case "VmPTE":
			str := strings.TrimSpace(content)
			info.VmPTE, _ = strconv.ParseUint(str[:len(str)-3], 10, 64)
		case "VmSwap":
			str := strings.TrimSpace(content)
			info.VmSwap, _ = strconv.ParseUint(str[:len(str)-3], 10, 64)
		}
	}

	// Fill theads id
	info.Task = AllThreads(pidStr)

	// Fill cmdline
	info.Cmdline = cmdline(pidStr)

	return info, nil
}
