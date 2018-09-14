package process

import (
	"strconv"
	"fmt"
	"os/user"

	"github.com/Lt0/sysmon/utils/proc/pid"
	"github.com/Lt0/sysmon/utils/proc"
)

type Process struct {
	Pid 		string		// 进程号			stat
	Task		[]string	// 所有线程
	Threads		uint64		// 线程数			status
	State		string		// 进程状态			status
	Priority	int64		// 动态优先级		stat
	Nice		int64		// 静态优先级		stat
	Comm		string		// 可执行文件名		comm, stat
	Cmdline		string		// 可执行文件路径	cmdline
	Uid			uint64
	User		string
	TaskCPU		int			// 运行在哪个 CPU 上	stat

	StartTime	uint64		// 系统开机后该进程启动的时间，单位为jiffies	stat
	UsedCPU		uint64		// 该进程所使用的 CPU，单位为 jiffies		stat
	// CpuTime		time.Time

	// SZ			string		// 进程占用的总内存，单位 page
	// RSS			string		// 真正具有数据页的物理内存（包含进程用到的所有共享库占用的全部内存）, byte			stat
	// PSS			uint64		// 进程占用的实际物理内存大小，对于该进程用到的共享库，会根据使用该库的进程数量，按比例显示该进程占用的内存
	// USS			uint64		// 进程独占的实际分配的私有内存。
	// SHR			uint64		// 进程调用的共享库占用的内存。 
	// VSS			string		// 进程可访问的全部地址空间, byte			stat


	// 以下和内存相关的数据从 /proc/$PID/status 和 /proc/$PID/statm 文件中获取，单位为 KB，字符串中不带单位
	VmSize		uint64		// 虚拟地址空间大小，
	VmRSS		uint64		// 内存部分的大小。 它包含以下三个部分（VmRSS = RssAnon + RssFile + RssShmem）
	VmPTE		uint64		// 该进程的所有页表的大小
	VmSwap		uint64		// 被交换到交换分区的匿名数据大小
	VmShare		uint64		// 共享部分的内存

	IOReadBytes		uint64		// 从磁盘读取的字节数, IOReadBytes
	IOWriteBytes	uint64		// 写入磁盘的字节数, IOWriteBytes
}

func PidInfo(pidStr string) (Process, error) {
	var info Process

	info.Pid = pidStr

	// Fill Pid, Comm, State, CPU, Priority, Nice, Nlwp, StartTime, VSS, RSS, TaskCPU by /proc/$PID/stat
	stat, _ := pid.Stat(pidStr)
	if len(stat.Comm) > 2 {
		info.Comm = stat.Comm[1:len(stat.Comm)-1]
	} else {
		info.Comm = stat.Comm
	}
	info.UsedCPU = uint64(stat.UTime) + uint64(stat.STime) + uint64(stat.CUTime) + uint64(stat.CSTime)
	info.Priority = stat.Priority
	info.Nice = stat.Nice
	info.StartTime = stat.StartTime
	info.TaskCPU = stat.Processor


	status, err := pid.Status(pidStr)
	if err != nil {
		return info, fmt.Errorf("run pid.Status failed: %v\n", err)
	}
	info.Uid = status.Uid[0]
	ui, err := user.LookupId(strconv.FormatUint(info.Uid, 10))
	if err != nil {
		return info, err
	}
	info.User = ui.Username

	info.State = status.State
	info.Threads = status.Threads
	info.VmSize = status.VmSize
	info.VmRSS = status.VmRSS
	info.VmPTE = status.VmPTE
	info.VmSwap = status.VmSwap

	// Fill theads id
	info.Task = pid.Task(pidStr)

	// Fill cmdline
	cmdline, _ := pid.Cmdline(pidStr)
	info.Cmdline = cmdline.Cmdline

	statm, err := pid.Statm(pidStr)
	if err != nil {
		return info, fmt.Errorf("run pid.Statm failed: %v\n", err)
	}
	info.VmShare = statm.Share

	ioInfo, err := pid.IO(pidStr)
	if err != nil {
		fmt.Println("io failed: err:", err)
		return info, fmt.Errorf("run pid.IO failed: %v\n", err)
	}
	info.IOReadBytes = ioInfo.ReadBytes
	info.IOWriteBytes = ioInfo.WriteBytes

	return info, nil
}

func coresInfo() ([]proc.CPU, int) {
	totalStat, _ := proc.Stat()

	return totalStat.CPUs, len(totalStat.CPUs) - 1
}

func uptimeInfo() (proc.UpTimeInfo) {
	ut, err := proc.UpTime()
	if err != nil {
		fmt.Println("uptimeInfo:", err)
		return ut
	}
	return ut
}
