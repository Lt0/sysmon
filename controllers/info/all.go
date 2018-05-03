package info

import (
	"fmt"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/Lt0/sysmon/controllers/info/mem"
	"github.com/astaxie/beego"
)

type AllInfo struct {
	Controller *beego.Controller

	allInfo SysInfo
}

type SysInfo struct {
	Cpu cpu.CpuInfo
	Mem mem.MemInfo
	// NetInfo struct
	// StorageInfo struct

	// HWInfo struct
	// SWInfo struct
}

func (p *AllInfo) Do() interface{} {
	var si SysInfo

	fmt.Println("do all info")
	si.Cpu = cpu.GetCpuInfo()
	si.Mem = mem.GetMemInfo()
	fmt.Println("si.Cpu: ", si.Cpu)
	fmt.Println("si.Mem: ", si.Mem)

	return nil
}
