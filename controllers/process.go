package controllers

import (
	"github.com/Lt0/sysmon/controllers/process"
	"github.com/astaxie/beego"
)

type ProcessController struct {
	beego.Controller
}

// @Title All Processes Status
// @Description Get All Processes Status
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /all [get]
func (c *ProcessController) AllProcess() {
	handler := &process.AllProcessCtrl{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title Process Details
// @Description Get Process Details
// @Param	pid	query	string	true	"which process to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /details [get]
func (c *ProcessController) Details() {
	handler := &process.DetailsCtrl{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}