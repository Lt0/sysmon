package sysInfo

import (
	"fmt"
	"os"
	"github.com/astaxie/beego"
)

type Hostname struct {
	Controller *beego.Controller
}

func (p *Hostname) Do() interface{} {
	fmt.Println("hostname:", os.Hostname)
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}