package controllers

import (
	// "fmt"

	"github.com/Lt0/sysmon/controllers/info"
	"github.com/Lt0/sysmon/controllers/info/disk"
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}

// @Title System Info
// @Description Get all system info
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /all [get]
func (c *InfoController) GetSysInfo() {
	// fmt.Println("GetSysInfo")
	handler := &info.AllInfo{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title System Disk Info
// @Description Get system disk info
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /disk [get]
func (c *InfoController) GetSysDiskInfo() {
	// fmt.Println("GetSysInfo")
	//handler := &disk.AllInfo{Controller: &c.Controller}
	c.Data["json"] = disk.AllInfo()
	c.ServeJSON()
}