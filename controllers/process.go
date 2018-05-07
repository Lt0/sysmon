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
	handler := &process.AllProcess{Controller: &c.Controller}
	c.Data["json"] = handler.Do()
	c.ServeJSON()
}
