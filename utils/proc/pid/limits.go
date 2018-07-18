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
type LimitsInfo struct {
	Limits	[]limitsItem
}

type limitsItem struct {
	Limit		string
	SoftLimit	string
	HardLimit	string
	Units		string
}

func Limits(pid string) (LimitsInfo, error) {
	var info LimitsInfo

	f, err := os.Open(filepath.Join(procfs, pid, "limits"))
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
		fmt.Println("length:", len(l))
		if len(vs) < 4 {
			return info, fmt.Errorf("Invalid string slice: %v\n", vs)
		}

		if vs[0] == "Limit" {
			continue
		}

		var item limitsItem
		item.Limit = strings.TrimSpace(l[:25])
		item.SoftLimit = strings.TrimSpace(l[25:46])
		item.HardLimit = strings.TrimSpace(l[46:67])
		item.Units = strings.TrimSpace(l[67:])
		info.Limits = append(info.Limits, item)
		fmt.Println("item: ", item)
	}

	return info, nil
}