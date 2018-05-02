package info

import (
	"fmt"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/astaxie/beego"
)

type AllInfo struct {
	Controller *beego.Controller

	allInfo SysInfo
}

type SysInfo struct {
	Cpu cpu.CpuInfo
	// MemInfo struct
	// NetInfo struct
	// StorageInfo struct

	// HWInfo struct
	// SWInfo struct
}

func (p *AllInfo) Do() interface{} {
	fmt.Println("do all info")

	return nil
}
