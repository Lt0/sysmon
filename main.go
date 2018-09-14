package main

import (
	"net"
	"fmt"

	_ "github.com/Lt0/sysmon/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	// "github.com/google/gopacket/pcap"

	"github.com/Lt0/sysmon/controllers/info/cpu"
	"github.com/Lt0/sysmon/controllers/info/disk"
	infoNet "github.com/Lt0/sysmon/controllers/info/net"

	proc "github.com/Lt0/sysmon/utils/proc"
	ppid "github.com/Lt0/sysmon/utils/proc/pid"
	pnet "github.com/Lt0/sysmon/utils/proc/net"
	"github.com/Lt0/sysmon/utils/capture"
)

func main() {
	setCORS()
	initStaticDir()
	initProcfs()
	initCapture()

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
	infoNet.Ctx.Procfs = procfs

	proc.Ctx.Procfs = procfs
	ppid.Ctx.Procfs = procfs
	pnet.Ctx.Procfs = procfs
}

func setCORS() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{   
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "PATCH"},        
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},     
		AllowCredentials: true,}))
}

func initCapture() {
	fmt.Println("init capture...")

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get ifaces failed:", err)
		return
	}
	
	for _, iface := range ifaces {
		fmt.Println("init capture routine on", iface.Name)
		go capture.CaptureDev(iface)
	}
}