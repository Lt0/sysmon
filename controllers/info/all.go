package info

import (
	"fmt"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/Lt0/sysmon/controllers/info/mem"
	"github.com/Lt0/sysmon/controllers/info/net"
	"github.com/astaxie/beego"
)

type AllInfo struct {
	Controller *beego.Controller

	allInfo SysInfo
}

type SysInfo struct {
	Cpu cpu.CpuInfo
	Mem mem.MemInfo
	Net net.NetInfo
	// StorageInfo struct

	// HWInfo struct
	// SWInfo struct
}

func (p *AllInfo) Do() interface{} {
	var si SysInfo

	fmt.Println("do all info")
	si.Cpu = cpu.GetCpuInfo()
	si.Mem = mem.GetMemInfo()
	si.Net = net.GetNetInfo()
	fmt.Println("si.Cpu: ", si.Cpu)
	fmt.Println("si.Mem: ", si.Mem)
	fmt.Println("si.Net: ", si.Net)

	return si
}
