package process

import (
	"path/filepath"
	"strconv"
	"fmt"
	"os"
	"bufio"
)

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

func AllThreads(pid string) []string {
	t, err := os.Open(filepath.Join("/proc", pid, "task"))
	if err != nil {
		fmt.Println("can not open /proc")
		return nil
	}

	defer t.Close()

	var tasks []string
	tasks, err = t.Readdirnames(0)
	if err != nil {
		fmt.Println("Readdirnames /proc failed")
		return nil
	}
	return tasks
}

func cmdline(pid string) string {
	cmdlineFile, err := os.Open(filepath.Join("/proc", pid, "cmdline"))
	if err != nil {
		fmt.Printf("cmdline: open /proc/%s/cmdline failed\n", pid)
		return ""
	}
	defer cmdlineFile.Close()

	cmdLineReader := bufio.NewReader(cmdlineFile)
	cmdline, _ := cmdLineReader.ReadBytes('\n')
	if len(cmdline) > 1 {
		return string(cmdline[:len(cmdline)-1])
	}
	
	return ""
}