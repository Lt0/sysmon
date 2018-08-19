package main

import (
	_ "github.com/Lt0/sysmon/routers"
	"github.com/astaxie/beego"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/Lt0/sysmon/controllers/info/disk"
	"github.com/Lt0/sysmon/controllers/info/net"

	proc "github.com/Lt0/sysmon/utils/proc"
	ppid "github.com/Lt0/sysmon/utils/proc/pid"
	pnet "github.com/Lt0/sysmon/utils/proc/net"
)

func main() {
	initStaticDir()
	initProcfs()

	beego.Run()
}

func initStaticDir(){
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/web"] = "web"
	beego.BConfig.WebConfig.StaticDir["/static"] = "web/dist/static"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func initProcfs() {
	procfs := beego.AppConfig.String("procfs")
	if procfs == "" {
		procfs = "/proc"
	}

	cpu.Ctx.Procfs = procfs
	disk.Ctx.Procfs = procfs
	net.Ctx.Procfs = procfs

	proc.Ctx.Procfs = procfs
	ppid.Ctx.Procfs = procfs
	pnet.Ctx.Procfs = procfs
}