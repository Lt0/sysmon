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
type SmapsInfo struct {
	Mappings	[]MapsItem
}

type MapsItem struct {
	StartAddr	string	// 起始虚拟地址
	EndAddr		string	// 终点虚拟地址
	Perm		string	// 权限
	Offset		string	// 偏移量
	Dev			string	// 映像文件所在的磁盘的主设备号和次设备号
	INode		string	// 映像文件的的磁盘节点（inode）
	File		string	// 文件/库/堆栈的路径

	Size			uint64
	Rss				uint64
	Pss				uint64
	SharedClean		uint64
	SharedDirty		uint64
	PrivateClean	uint64
	PrivateDirty	uint64
	Referenced		uint64
	Anonymous		uint64
	AnonHugePages	uint64
	SharedHugetlb	uint64
	PrivateHugetlb	uint64
	Swap			uint64
	SwapPss			uint64
	KernelPageSize	uint64
	MMUPageSize		uint64
	Locked			uint64
	VmFlags			string
}

func Smaps(pid string) (SmapsInfo, error) {
	var info SmapsInfo

	f, err := os.Open(filepath.Join(procfs, pid, "smaps"))
	if err != nil {
		return info, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		var item MapsItem
		for i := 0; i < 16; i++ {
			b, err := r.ReadBytes('\n')
			if err == io.EOF {
				goto finish
			}

			s := string(b)
			vs := strings.Fields(s)
			// fmt.Printf("s: %v: vs[1]: %v\n", s, vs[1])
			switch vs[0] {
			case "Size:":
				item.Size = parseValue(vs[1])
			case "Rss:":
				item.Rss = parseValue(vs[1])
			case "Pss:":
				item.Pss = parseValue(vs[1])
			case "Shared_Clean:":
				item.SharedClean = parseValue(vs[1])
			case "Shared_Dirty:":
				item.SharedDirty = parseValue(vs[1])
			case "Private_Clean:":
				item.PrivateClean = parseValue(vs[1])
			case "Private_Dirty:":
				item.PrivateDirty = parseValue(vs[1])
			case "Referenced:":
				item.Referenced = parseValue(vs[1])
			case "Anonymous:":
				item.Anonymous = parseValue(vs[1])
			case "AnonHugePages:":
				item.AnonHugePages = parseValue(vs[1])
			case "Shared_Hugetlb:":
				item.SharedHugetlb = parseValue(vs[1])
			case "Private_Hugetlb:":
				item.PrivateHugetlb = parseValue(vs[1])
			case "Swap:":
				item.Swap = parseValue(vs[1])
			case "SwapPss:":
				item.SwapPss = parseValue(vs[1])
			case "KernelPageSize:":
				item.KernelPageSize = parseValue(vs[1])
			case "MMUPageSize:":
				item.MMUPageSize = parseValue(vs[1])
			case "Locked:": 
				item.Locked = parseValue(vs[1])
			case "VmFlags:":
				item.VmFlags = parseVmFlags(s)
				break
			default:
				parseItemHdr(s, &item)
			}
		}
		info.Mappings = append(info.Mappings, item)
	}

finish:
	return info, nil
}

func parseItemHdr(s string, item *MapsItem) {
	vs := strings.Fields(s)
	if len(vs) < 1 {
		return
	}

	addrs := strings.Split(vs[0], "-")
	if len(addrs) < 1 {
		return
	}
	item.StartAddr = addrs[0]

	if len(addrs) > 1 {
		item.EndAddr = addrs[1]
	}

	if len(vs) > 1 {
		item.Perm = vs[1]
	}
	if len(vs) > 2 {
		item.Offset = vs[2]
	}
	if len(vs) > 3 {
		item.Dev = vs[3]
	}
	if len(vs) > 4 {
		item.INode = vs[4]
	}
	if len(vs) > 5 {
		item.File = vs[5]
	}
	
}

func parseValue(s string) uint64 {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Printf("pid.smaps: parseValue(%s) ParseUint failed: err: %v\n", s, err)
	}
	return num
}

func parseVmFlags(s string) string {
	vs := strings.Split(s, ":")
	return vs[1]
}