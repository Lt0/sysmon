package process

import (
	"fmt"
	"os/exec"
	// "strings"
	// "strconv"
	
	"github.com/astaxie/beego"
)

type AllProcessCtrl struct {
	Controller *beego.Controller
}

type AllProcess struct {
	Processes 	[]Process
}

func (p *AllProcessCtrl) Do() interface{} {
	var ap AllProcess

	fmt.Println("do AllProcess")
	p.All(&ap)

	return ap
}

func (p *AllProcessCtrl) All(ap *AllProcess) {
	filter := `tail -n +2 | sed 's/^\s*//g' | sed 's/\s\+/ /g'`
	cmd := fmt.Sprintf("ps caxo %s | %s", psOutputFormatStr, filter)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("AllProcessCtrl", err)
		return
	}
	ap.Processes = FormatProcStatus(out)
}
