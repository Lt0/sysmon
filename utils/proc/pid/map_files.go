package pid

import (
	"os"
	"fmt"
	"path/filepath"
)

type MapFilesInfo struct {
	Mappings	[]MapFilesItem
}

type MapFilesItem struct {
	AddrRange	string
	File		string
}

func MapFiles(pid string) MapFilesInfo {
	var info MapFilesInfo

	path := filepath.Join(Ctx.Procfs, pid, "map_files")
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
		var item MapFilesItem
		item.AddrRange = v
		file, err := os.Readlink(filepath.Join(path, v))
		if err != nil {
			item.File = fmt.Sprint(err)
		} else {
			item.File = file
		}
		
		info.Mappings = append(info.Mappings, item)
	}

	return info
}