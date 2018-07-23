package process

import (
	// "strconv"
	"fmt"
	"time"

	"github.com/Lt0/sysmon/utils/proc/pid"
	"github.com/Lt0/sysmon/utils/proc"
	"github.com/astaxie/beego"
)

type DetailsCtrl struct {
	Controller *beego.Controller

	pid	string
	threads []string
	details detailsInfo
}

type detailsInfo struct {
	TimeStamp 	int64
	CoreNum		int				// 设备的内核总数
	Cores		[]proc.CPU		// 每个 CPU 核心的信息
	UpTime		proc.UpTimeInfo	// 系统启动时间和 idle 时间
	Processes 	[]Process		// 所有的线程信息

	// 进程的详细信息
	Limits		pid.LimitsInfo
	Stacks		[]stacksInfo
	Smaps		pid.SmapsInfo
	NumaMaps	pid.NumaMapsInfo
	CWD			string
	OOMAdj		int
	OOMScore	int
	OOMScoreAdj	int
	FD			pid.FDInfo
	MapFiles	pid.MapFilesInfo

	// raw data, 文件的原始数据
	RDSched				string
	// RDLimits			string
	RDStatus			string
	RDEnviron			string
	RDCPUSet			string
	RDMountInfo			string	
	RDMounts			string	
	RDMountStats		string	
	RDCGroup			string
	RDAutoGroup			string
	RDCoredumpFilter	string
	RDUidMap			string
	RDGidMap			string
	RDLoginUid			string
	RDPersonality		string
	RDProjidMap			string
	RDSessionID			string
	RDSetGroup			string
	RDSyscall			string
}

type stacksInfo struct {
	Pid		string
	Stack	pid.StackInfo
}

func (p *DetailsCtrl) Do() interface{} {
	p.fillTimeStamp()

	p.param()
	p.threads = proc.AllThreadPids(p.pid)
	p.details.Cores, p.details.CoreNum = coresInfo()
	p.details.UpTime = uptimeInfo()

	p.fillProcesses()
	p.fillLimits()
	p.fillStack()
	p.fillSmaps()
	p.fillNumaMaps()
	p.details.CWD = pid.CWD(p.pid)
	p.details.OOMAdj = pid.OOMAdj(p.pid)
	p.details.OOMScore = pid.OOMScore(p.pid)
	p.details.FD = pid.FD(p.pid)
	p.details.MapFiles = pid.MapFiles(p.pid)

	p.details.RDSched = pid.SchedRawData(p.pid)
	// p.details.RDLimits = pid.LimitsRawData(p.pid)
	p.details.RDStatus = pid.StatusRawData(p.pid)
	p.details.RDEnviron = pid.EnvironRawData(p.pid)
	p.details.RDCPUSet = pid.CPUSetRawData(p.pid)
	p.details.RDMountInfo = pid.MountInfoRawData(p.pid)
	p.details.RDMounts = pid.MountsRawData(p.pid)
	p.details.RDMountStats = pid.MountStatsRawData(p.pid)
	p.details.RDCGroup = pid.CGroupRawData(p.pid)
	p.details.RDAutoGroup = pid.AutoGroupRawData(p.pid)
	p.details.RDCoredumpFilter = pid.CoredumpFilterRawData(p.pid)
	p.details.RDUidMap = pid.UidMapRawData(p.pid)
	p.details.RDGidMap = pid.GidMapRawData(p.pid)
	p.details.RDLoginUid = pid.LoginUidRawData(p.pid)
	p.details.RDPersonality = pid.PersonalityRawData(p.pid)
	p.details.RDProjidMap = pid.ProjidMapRawData(p.pid)
	p.details.RDSessionID = pid.SessionIDRawData(p.pid)
	p.details.RDSetGroup = pid.SetGroupRawData(p.pid)
	p.details.RDSyscall = pid.SyscallRawData(p.pid)
	
	return p.details
}

func (p *DetailsCtrl) param() {
	pid := p.Controller.GetString("pid")

	fmt.Println("pid:", pid)
	p.pid = pid
}

func (p *DetailsCtrl) fillTimeStamp() {
	t := time.Now()
	p.details.TimeStamp = t.UnixNano()
}

func (p *DetailsCtrl) fillProcesses() {
	for _, v := range(p.threads) {
		info, err := PidInfo(v)
		if err != nil {
			// fmt.Printf("not add %v cause by: %v\n", v, err)
			continue
		}
		p.details.Processes = append(p.details.Processes, info)
	}
}

func (p *DetailsCtrl) fillLimits() {
	var err error
	p.details.Limits, err = pid.Limits(p.pid)
	if err != nil {
		fmt.Println("fillLimits: ", err)
	}
}

func (p *DetailsCtrl) fillStack() {
	for _, v := range(p.threads) {
		info, err := pid.Stack(v)
		if err != nil {
			// fmt.Printf("not add %v cause by: %v\n", v, err)
			continue
		}

		var sinfo stacksInfo
		sinfo.Pid = v
		sinfo.Stack = info
		p.details.Stacks = append(p.details.Stacks, sinfo)
	}
}

func (p *DetailsCtrl) fillSmaps() {
	info, err := pid.Smaps(p.pid)
	if err != nil {
		fmt.Println("fillSmaps: ", err)
	}
	p.details.Smaps = info
}

func (p *DetailsCtrl) fillNumaMaps() {
	info, err := pid.NumaMaps(p.pid)
	if err != nil {
		fmt.Println("fillNumaMaps: ", err)
	}
	p.details.NumaMaps = info
}