package pid

import (
	"os"
	"fmt"
	"path/filepath"
)

type FDInfo struct {
	FDs	[]FDItem
}

type FDItem struct {
	Name	string
	File	string
}

func FD(pid string) FDInfo {
	var info FDInfo

	path := filepath.Join(procfs, pid, "fd")
	p, err := os.Open(path)
	if err != nil {
		fmt.Println("can not open", path)
		return info
	}

	defer p.Close()

	var files []string
	files, err = p.Readdirnames(0)
	if err != nil {
		fmt.Printf("Readdirnames %v failed\n", path)
		return info
	}

	for _, v := range(files) {
		var item FDItem
		item.Name = v
		file, err := os.Readlink(filepath.Join(path, v))
		if err != nil {
			item.File = fmt.Sprint(err)
		} else {
			item.File = file
		}
		
		info.FDs = append(info.FDs, item)
	}

	return info
}