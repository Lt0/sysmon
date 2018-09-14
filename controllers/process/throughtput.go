package process

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ThroughputCtrl struct {
	Controller *beego.Controller
}

func (p *ThroughputCtrl)Do() interface{} {
	fmt.Println("Do Throughtput")
	return "Do Throughput"
}