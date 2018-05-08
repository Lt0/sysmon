package process

import (
	"fmt"
	"os/exec"
	
	"github.com/astaxie/beego"
)

type MyProcessCtrl struct {
	Controller *beego.Controller
}

type MyProcess struct {
	Processes 	[]Process
}

func (p *MyProcessCtrl) Do() interface{} {
	var ap MyProcess

	fmt.Println("do MyProcess")
	p.MyProcess(&ap)

	return ap
}

func (p *MyProcessCtrl) MyProcess(ap *MyProcess) {
	filter := `tail -n +2 | sed 's/^\s*//g' | sed 's/\s\+/ /g'`
	cmd := fmt.Sprintf("ps cxo %s | %s", psOutputFormatStr, filter)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Get ActiveProcess", err)
		return
	}

	ap.Processes = FormatProcStatus(out)
}
