package process

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	
	"github.com/astaxie/beego"
)

type AllProcess struct {
	Controller *beego.Controller
}

type AllProcessStatus struct {
	Processes 	[]Process
}

func (p *AllProcess) Do() interface{} {
	var ap AllProcessStatus

	fmt.Println("do AllProcess")
	p.All(&ap)

	return ap
}

func (p *AllProcess) All(ap *AllProcessStatus) {
	out, err := exec.Command("sh", "-c", `ps caxo command,user,%cpu,cputime,%mem,sz,rss,drs,trs,vsz,pid,nlwp,state,nice | tail -n +2 | sed 's/^\s*//g' | sed 's/\s\+/ /g'`).Output()
	if err != nil {
		fmt.Println("Disk Storage", err)
		return
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrs := strings.Split(v, " ")

		cmd := attrs[0]
		user := attrs[1]

		cpu, _ := strconv.ParseFloat(attrs[2], 64)
		
		cpuTime := attrs[3]
		
		mem, _ := strconv.ParseFloat(attrs[4], 64)
		sz, _ := strconv.ParseUint(attrs[5], 10, 64)
		rss, _ := strconv.ParseUint(attrs[6], 10, 64)
		drs, _ := strconv.ParseUint(attrs[7], 10, 64)
		trs, _ := strconv.ParseUint(attrs[8], 10, 64)
		vsz, _ := strconv.ParseUint(attrs[9], 10, 64)

		pid, _ := strconv.Atoi(attrs[10])
		nlwp, _ := strconv.Atoi(attrs[9])

		state := attrs[12]
		nice, _ := strconv.Atoi(attrs[13])

		ap.Processes = append(ap.Processes, Process{cmd, user, cpu, cpuTime, mem, sz, rss, drs, trs, vsz, pid, nlwp, state, nice})

	}
}
