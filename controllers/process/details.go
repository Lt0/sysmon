package process

import (
	"strconv"
	"fmt"
	"time"

	"github.com/Lt0/sysmon/utils/proc/pid"
	"github.com/Lt0/sysmon/utils/proc"
	"github.com/astaxie/beego"
)

type DetailsCtrl struct {
	Controller *beego.Controller

	pid	int
	threads []string
	details detailsInfo
}

type detailsInfo struct {
	TimeStamp 	time.Time
	CoreNum		int				// 设备的内核总数
	Cores		[]proc.CPU		// 每个 CPU 核心的信息
	UpTime		proc.UpTimeInfo	// 系统启动时间和 idle 时间
	Processes 	[]Process		// 所有的线程信息

	// 进程的详细信息
	Limits		pid.LimitsInfo
	Stacks		[]stacksInfo
	Smaps		pid.SmapsInfo
	NumaMaps	pid.NumaMapsInfo
}

type stacksInfo struct {
	Pid		string
	Stack	pid.StackInfo
}

func (p *DetailsCtrl) Do() interface{} {
	p.param()
	p.threads = proc.AllThreadPids(strconv.Itoa(p.pid))
	p.details.TimeStamp = time.Now()
	p.details.Cores, p.details.CoreNum = coresInfo()
	p.details.UpTime = uptimeInfo()

	p.fillProcesses()
	p.fillLimits()
	p.fillStack()
	p.fillSmaps()
	p.fillNumaMaps()

	return p.details
}

func (p *DetailsCtrl) param() {
	pid, err := p.Controller.GetInt("pid")
	if err != nil {
		fmt.Println("get param failed:", err)
	}

	fmt.Println("pid:", pid)
	p.pid = pid
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
	p.details.Limits, err = pid.Limits(strconv.Itoa(p.pid))
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
	info, err := pid.Smaps(strconv.Itoa(p.pid))
	if err != nil {
		fmt.Println("fillSmaps: ", err)
	}
	p.details.Smaps = info
}

func (p *DetailsCtrl) fillNumaMaps() {
	info, err := pid.NumaMaps(strconv.Itoa(p.pid))
	if err != nil {
		fmt.Println("fillNumaMaps: ", err)
	}
	p.details.NumaMaps = info
}