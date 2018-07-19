package pid

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// containe all info of /proc/[pid]/statm
type StackInfo struct {
	Stacks	[]stackItem
}

type stackItem struct {
	Address	string
	Name	string
}

func Stack(pid string) (StackInfo, error) {
	var info StackInfo

	f, err := os.Open(filepath.Join(procfs, pid, "stack"))
	if err != nil {
		return info, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		l := string(b)
		vs := strings.Fields(l)

		if len(vs) < 2 {
			return info, fmt.Errorf("Invalid string slice: %v\n", vs)
		}

		var item stackItem
		item.Address = vs[0]
		item.Address = item.Address[2:len(item.Address)-2]
		item.Name = vs[1]
		info.Stacks = append(info.Stacks, item)
	}

	return info, nil
}