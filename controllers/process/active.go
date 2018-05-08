package process

import (
	"fmt"
	"os/exec"
	
	"github.com/astaxie/beego"
)

type ActiveProcessCtrl struct {
	Controller *beego.Controller
}

type ActiveProcess struct {
	Processes 	[]Process
}

func (p *ActiveProcessCtrl) Do() interface{} {
	var ap ActiveProcess

	fmt.Println("do ActiveProcess")
	p.ActiveProcess(&ap)

	return ap
}

func (p *ActiveProcessCtrl) ActiveProcess(ap *ActiveProcess) {
	filter := `tail -n +2 | awk '$13 == "R" {print}'| sed 's/^\s*//g' | sed 's/\s\+/ /g'`
	cmd := fmt.Sprintf("ps caxo %s | %s", psOutputFormatStr, filter)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Get ActiveProcess", err)
		return
	}
	
	ap.Processes = FormatProcStatus(out)
}
