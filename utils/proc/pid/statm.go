package pid

import (
	"strconv"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// containe all info of /proc/[pid]/statm
type StatmInfo struct {
	// 1. total program size (same as VmSize in /proc/[pid]/status)
	Size     uint64 

	// 2. resident set size (same as VmRSS in /proc/[pid]/status)
	Resident uint64 

	// 3. shared pages (i.e., backed by a file)
	Share    uint64 

	// 4. text (code)
	Text     uint64 

	// 5. library (unused in Linux 2.6)
	Lib      uint64 

	// 6. data + stack
	Data     uint64 

	// 7. dirty pages (unused in Linux 2.6)
	Dt       uint64 
}

// get statm info from /proc/[pid]/statm.
func Statm(pid string) (StatmInfo, error) {
	var statm StatmInfo

	f, err := os.Open(filepath.Join(procfs, pid, "statm"))
	if err != nil {
		return statm, fmt.Errorf("Statm: Open %v/%v/statm failed: %v\n", procfs, pid, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for i := 1; ; i++ {
		b, err := r.ReadBytes(' ')
		if err == io.EOF {
			break
		}
		s := strings.TrimSpace(string(b))

		switch i {
		case 1:
			statm.Size, _ = strconv.ParseUint(s, 10, 64)
		case 2:
			statm.Resident, _ = strconv.ParseUint(s, 10, 64)
		case 3:
			statm.Share, _ = strconv.ParseUint(s, 10, 64)
		case 4:
			statm.Text, _ = strconv.ParseUint(s, 10, 64)
		case 5:
			statm.Lib, _ = strconv.ParseUint(s, 10, 64)
		case 6:
			statm.Data, _ = strconv.ParseUint(s, 10, 64)
		case 7:
			statm.Dt, _ = strconv.ParseUint(s, 10, 64)
		default:
			// fmt.Println("unknow value in statm:", s)
		}
	}

	return statm, nil
}
