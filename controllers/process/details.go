package process

import (
	"strconv"
	"fmt"
	"time"

	"github.com/Lt0/sysmon/utils/proc"
	"github.com/astaxie/beego"
)

type DetailsCtrl struct {
	Controller *beego.Controller

	pid	int
}

type details struct {
	TimeStamp 	time.Time
	CoreNum		int				// 设备的内核总数
	Cores		[]proc.CPU		// 每个 CPU 核心的信息
	UpTime		proc.UpTimeInfo	// 系统启动时间和 idle 时间
	Processes 	[]Process		// 所有的线程信息
}

func (p *DetailsCtrl) Do() interface{} {
	var d details

	p.param()

	d.TimeStamp = time.Now()
	d.Cores, d.CoreNum = coresInfo()
	d.UpTime = uptimeInfo()

	p.AllPidInfo(&d)

	return d
}

func (p *DetailsCtrl) param() {
	pid, err := p.Controller.GetInt("pid")
	if err != nil {
		fmt.Println("get param failed:", err)
	}

	fmt.Println("pid:", pid)
	p.pid = pid
}

func (p *DetailsCtrl) AllPidInfo(d *details) {
	for _, v := range(proc.AllThreadPids(strconv.Itoa(p.pid))) {
		info, err := PidInfo(v)
		if err != nil {
			// fmt.Printf("not add %v cause by: %v\n", v, err)
			continue
		}
		d.Processes = append(d.Processes, info)
	}
}
