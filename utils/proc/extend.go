package proc

import (
	"path/filepath"
	"os"
	"fmt"
	"strconv"
)

// return all pids in procfs
func AllPids() []string {
	return pids(procfs)
}

func AllThreadPids(pid string) []string {
	return pids(filepath.Join(procfs, pid, "task"))
}

func pids(path string) []string {
	p, err := os.Open(path)
	if err != nil {
		fmt.Println("can not open", path)
		return nil
	}

	defer p.Close()

	var files []string
	files, err = p.Readdirnames(0)
	if err != nil {
		fmt.Printf("Readdirnames %v failed\n", path)
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