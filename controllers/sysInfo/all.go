package sysInfo

import (
	"github.com/astaxie/beego"
)

type AllInfo struct {
	Controller *beego.Controller

	allSysInfo SysInfo
}

type SysInfo struct {
	HW HWInfo
	SW SWInfo
}

func (p *AllInfo) Do() interface{} {
	var si SysInfo
	si.HW = getHWInfo()
	return si
}
