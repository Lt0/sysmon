package pid

import (
	"bufio"
	// "fmt"
	"io"
	"io/ioutil"
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

	f, err := os.Open(filepath.Join(Ctx.Procfs, pid, "limits"))
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
		if vs[0] == "Limit" {
			continue
		}

		var item limitsItem
		item.Limit = strings.TrimSpace(l[:25])
		item.SoftLimit = strings.TrimSpace(l[25:46])
		item.HardLimit = strings.TrimSpace(l[46:67])
		item.Units = strings.TrimSpace(l[67:])
		info.Limits = append(info.Limits, item)
	}

	return info, nil
}


func LimitsRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "limits"))
	return string(buf)
}