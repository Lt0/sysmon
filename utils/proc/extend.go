package proc

import (
	"os"
	"fmt"
	"strconv"
)

// return all pids in /proc
func AllPids() []string {
	p, err := os.Open("/proc")
	if err != nil {
		fmt.Println("can not open /proc")
		return nil
	}

	defer p.Close()

	var files []string
	files, err = p.Readdirnames(0)
	if err != nil {
		fmt.Println("Readdirnames /proc failed")
		return nil
	}

	var pids []string
	for _, v := range(files) {
		if pid, _ := strconv.Atoi(v); pid != 0 {
			pids = append(pids, v)
		}
	}
	return pids
}