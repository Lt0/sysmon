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
		if len(vs) < 4 {
			return info, fmt.Errorf("Invalid string slice: %v\n", vs)
		}

		var item limitsItem
		item.Limit = vs[0]
		item.SoftLimit = vs[1]
		item.HardLimit = vs[2]
		item.Units = vs[3]
		info.Limits = append(info.Limits, item)
	}

	return info, nil
}