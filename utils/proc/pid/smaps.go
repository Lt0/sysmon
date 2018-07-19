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
	Swap			uint64
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
			switch i {
			case 0:
				parseItemHdr(s, &item)
			case 1:
				item.Size = parseValue(s)
			case 2:
				item.Rss = parseValue(s)
			case 3:
				item.Pss = parseValue(s)
			case 4:
				item.SharedClean = parseValue(s)
			case 5:
				item.SharedDirty = parseValue(s)
			case 6:
				item.PrivateClean = parseValue(s)
			case 7:
				item.PrivateDirty = parseValue(s)
			case 8:
				item.Referenced = parseValue(s)
			case 9:
				item.Anonymous = parseValue(s)
			case 10:
				item.AnonHugePages = parseValue(s)
			case 11:
				item.Swap = parseValue(s)
			case 12:
				item.KernelPageSize = parseValue(s)
			case 13:
				item.MMUPageSize = parseValue(s)
			case 14:
				item.Locked = parseValue(s)
			case 15:
				item.VmFlags = parseVmFlags(s)
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
	item.StartAddr = addrs[0]
	item.EndAddr = addrs[1]

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
	vs := strings.Fields(s)
	num, err := strconv.ParseUint(vs[1], 10, 64)
	if err != nil {
		fmt.Println("ParseUint failed")
	}
	return num
}

func parseVmFlags(s string) string {
	vs := strings.Split(s, ":")
	return vs[1]
}