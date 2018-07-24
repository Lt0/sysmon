package proc

import (
	"strconv"
	"strings"
	"path/filepath"
	"io/ioutil"
)

type BuddyInfoInfo struct {
	Zones	[]BuddyInfoZone
}

type BuddyInfoZone	struct {
	Node	string
	Type	string
	Pages	[]uint64	// Linux 通过 buddy 算法，把所有空闲的内存，以2的幂次方的形式，分成 11 个块链表，
						// 分别对应为 1、2、4、8、16、32、64、128、256、512、1024 个页块的链表（比如最后一个链表中记录的是所有可用的连续 1024 页的内存）。
}

func BuddyInfo() (BuddyInfoInfo, error) {
	var info BuddyInfoInfo

	buf, err := ioutil.ReadFile(filepath.Join(procfs, "buddyinfo"))
	if err != nil {
		return info, err
	}

	lines := strings.Split(string(buf), "\n")
	for _, s := range(lines) {
		var zone BuddyInfoZone
		vs := strings.Fields(s)
		if len(vs) < 15 {
			continue
		}

		zone.Node = vs[0] + " " + vs[1][:len(vs[1])-1]
		zone.Type = vs[3]
		zone.Pages = parseUint64List(vs[4:])
		info.Zones = append(info.Zones, zone)
	}

	return info, nil
}

func parseUint64List(strs []string) []uint64 {
	var nums []uint64
	for _, v := range(strs) {
		num, _ := strconv.ParseUint(strings.TrimSpace(v), 10, 64)
		nums = append(nums, num)
	}
	return nums
}