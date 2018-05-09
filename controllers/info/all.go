package info

import (
	// "fmt"
	"time"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/Lt0/sysmon/controllers/info/mem"
	"github.com/Lt0/sysmon/controllers/info/net"
	"github.com/Lt0/sysmon/controllers/info/disk"
	"github.com/astaxie/beego"
)

type AllInfo struct {
	Controller *beego.Controller

	allInfo SysInfo
}

type SysInfo struct {
	TimeStamp time.Time
	Cpu cpu.CpuInfo
	Mem mem.MemInfo
	Net net.NetInfo
	Disk disk.DiskInfo

	// HWInfo struct
	// SWInfo struct
}

func (p *AllInfo) Do() interface{} {
	var si SysInfo

	// fmt.Println("do all info")
	
	si.TimeStamp = time.Now();
	// fmt.Println("TimeStamp:", si.TimeStamp);

	si.Cpu = cpu.GetCpuInfo()
	si.Mem = mem.GetMemInfo()
	si.Net = net.GetNetInfo()
	si.Disk = disk.AllInfo()
	// fmt.Println("si.Cpu: ", si.Cpu)
	// fmt.Println("si.Mem: ", si.Mem)
	// fmt.Println("si.Net: ", si.Net)
	// fmt.Println("si.Disk: ", si.Disk)

	return si
}
