// @APIVersion 1.0.0
// @Title System Monitor API
// @Description System Monitor API
// @Contact lightimehpq@gmail.com
package routers

import (
	"github.com/Lt0/sysmon/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    ns := beego.NewNamespace("/v1",
        beego.NSNamespace("/process",
            beego.NSInclude(
                &controllers.ProcessController{},
            ),
        ),
        beego.NSNamespace("/info",
            beego.NSInclude(
                &controllers.InfoController{},
            ),
        ),
        
    )

    beego.AddNamespace(ns)
}
