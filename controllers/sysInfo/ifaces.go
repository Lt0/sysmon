package sysInfo

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/google/gopacket/pcap"
)

type Ifaces struct {
	Controller *beego.Controller

	Devs []pcap.Interface
}

func (p *Ifaces) Do() interface{} {
	p.getAllDev()
	return p.Devs
}

func (p *Ifaces) getAllDev() {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println("getAllDev failed:", err)
		return
	}

	p.Devs = devs
}
