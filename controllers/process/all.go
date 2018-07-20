package process

import (
	"fmt"
	"time"

	"github.com/Lt0/sysmon/utils/proc"
	"github.com/astaxie/beego"
)

type AllProcessCtrl struct {
	Controller *beego.Controller
}

type AllProcess struct {
	TimeStamp 	int64
	CoreNum		int				// 设备的内核总数
	Cores		[]proc.CPU		// 每个 CPU 核心的信息
	UpTime		proc.UpTimeInfo	// 系统启动时间和 idle 时间
	Processes 	[]Process		// 所有的线程信息
}

func (p *AllProcessCtrl) Do() interface{} {
	var ap AllProcess

	fmt.Println("do AllProcess")
	t := time.Now()
	ap.TimeStamp = t.UnixNano()
	p.AllPidInfo(&ap)
	ap.Cores, ap.CoreNum = coresInfo()
	ap.UpTime = uptimeInfo()

	return ap
}

func (p *AllProcessCtrl) AllPidInfo(ap *AllProcess) {
	for _, v := range(proc.AllPids()) {
		info, err := PidInfo(v)
		if err != nil {
			// fmt.Printf("not add %v cause by: %v\n", v, err)
			continue
		}
		ap.Processes = append(ap.Processes, info)
	}
}
