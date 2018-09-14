package sysInfo

import (
	"fmt"
	"os/exec"

	"github.com/Lt0/sysmon/utils/proc"
)

type Memory struct {
	PhySize		uint64
	SWapSize	uint64
}

type HWInfo struct {
	Mem		Memory
	Lshw	string
}

// GHW is global hardware info
var GHW HWInfo
var initedGlobalHW bool

func getHWInfo() HWInfo {
	if initedGlobalHW {
		return GHW
	}

	// GHW.Lshw = lshwStr()

	initedGlobalHW = true

	mi, _ := proc.GetMeminfo()
	GHW.Mem.PhySize = mi.MemTotal
	GHW.Mem.SWapSize = mi.SwapTotal

	return GHW
}

func lshwStr() string {
	cmd := fmt.Sprintf("lshw")
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("lshw:", err)
	}
	// fmt.Println("out:", out)
	return string(out)
}