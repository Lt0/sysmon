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