package controllers

import (
	// "fmt"

	"github.com/Lt0/sysmon/controllers/sysInfo"
	"github.com/astaxie/beego"
)

type SysInfoController struct {
	beego.Controller
}

// @Title Static System Global Info
// @Description Get all static system global info
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /all [get]
func (c *SysInfoController) GetSysInfo() {
	// fmt.Println("GetSysInfo")
	handler := &sysInfo.AllInfo{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title Hostname
// @Description Get hostname reported by the kernel.
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /hostname [get]
func (c *SysInfoController) GetHostname() {
	handler := &sysInfo.Hostname{}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title Network Interfaces
// @Description Get all network interfaces
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /ifaces [get]
func (c *SysInfoController) GetIfaces() {
	handler := &sysInfo.Ifaces{}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}