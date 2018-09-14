package sysInfo

import (
	"os"
	"github.com/astaxie/beego"
)

type Hostname struct {
	Controller *beego.Controller
}

func (p *Hostname) Do() interface{} {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}