package controllers

import (
	"fmt"

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
	fmt.Println("AllProcess")
	handler := &process.AllProcessCtrl{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title Active Processes Status
// @Description Get Active Processes Status
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /active [get]
func (c *ProcessController) ActiveProcess() {
	fmt.Println("ActiveProcess")
	handler := &process.ActiveProcessCtrl{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}

// @Title My Processes Status
// @Description Get My Processes Status
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /my [get]
func (c *ProcessController) MyProcess() {
	fmt.Println("MyProcess")
	handler := &process.MyProcessCtrl{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}