package pid

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"path/filepath"
)

type IOInfo struct {
	RChar				uint64	// characters read. The number of bytes which this task has caused to be read from storage. This is simply the sum of bytes which this process passed to read(2) and similar system calls. It includes things such as terminal I/O and is unaffected by whether or not actual physical disk I/O was required (the read might have been satisfied from pagecache).
	WChar				uint64	// characters written. The number of bytes which this task has caused, or shall cause to be written to disk. Similar caveats apply here as with rchar.
	SysCR				uint64	// read syscalls. Attempt to count the number of read I/O operations—that is, system calls such as read(2) and pread(2).
	SysCW				uint64	// Attempt to count the number of write I/O operations—that is, system calls such as write(2) and pwrite(2).
	ReadBytes			uint64	// Attempt to count the number of bytes which this process really did cause to be fetched from the storage layer.  This is accurate for block-backed filesystems.
	WriteBytes			uint64	// Attempt to count the number of bytes which this process caused to be sent to the storage layer.
	CancelledWriteBytes	uint64	// The big inaccuracy here is truncate.
	// If a process writes 1MB to a file and then deletes the file, it will
	// in  fact  perform  no  writeout.  But it will have been accounted as having caused 1MB of write.  In other
	// words: this field represents the number of bytes which this process caused to not  happen,  by  truncating
	// pagecache.   A  task  can cause "negative" I/O too.  If this task truncates some dirty pagecache, some I/O
	// which another task has been accounted for (in its write_bytes) will not be happening.
}

func IO(pid string) (IOInfo, error) {
	var info IOInfo
	f, err := os.Open(filepath.Join("/proc", pid, "io"))
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
		vs := strings.Split(l, ":")
		if len(vs) < 2 {
			return info, fmt.Errorf("Invalid string slice: %v\n", vs)
		}

		num, err := strconv.ParseUint(strings.TrimSpace(vs[1]), 10, 64)
		if err != nil {
			return info, err
		}

		switch vs[0] {
		case "rchar":
			info.RChar = num
		case "wchar":
			info.WChar = num
		case "syscr":
			info.SysCR = num
		case "syscw":
			info.SysCW = num
		case "read_bytes":
			info.ReadBytes = num
		case "write_bytes":
			info.WriteBytes = num
		case "cancelled_write_bytes":
			info.CancelledWriteBytes = num
		default:
			fmt.Println("Pid IO(): unknow pattern:", vs[0])
		}
	}

	return info, nil
}