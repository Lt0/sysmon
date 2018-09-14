package mem

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/Lt0/sysmon/utils/proc"
)

// unit of mem info is KB
type MemInfo struct {
	MemTotal uint64
	MemFree uint64
	MemAvailable uint64
	Buffers uint64
	Cached uint64
	SwapCached uint64
	Active uint64
	Inactive uint64
	ActiveAnon uint64
	InactiveAnon uint64
	ActiveFile uint64
	InactiveFile uint64
	Unevictable uint64
	Mlocked uint64
	SwapTotal uint64
	SwapFree uint64
	Dirty uint64
	Writeback uint64
	AnonPages uint64
	Mapped uint64
	Shmem uint64
	Slab uint64
	SReclaimable uint64
	SUnreclaim uint64
	KernelStack uint64
	PageTables uint64
	NfsUnstable uint64
	Bounce uint64
	WritebackTmp uint64
	CommitLimit uint64
	CommittedAS uint64
	VmallocTotal uint64
	VmallocUsed uint64
	VmallocChunk uint64
	HardwareCorrupted uint64
	AnonHugePages uint64
	CmaTotal uint64
	CmaFree uint64
	HugePagesTotal uint64
	HugePagesFree uint64
	HugePagesRsvd uint64
	HugePagesSurp uint64
	Hugepagesize uint64
	DirectMap4k uint64
	DirectMap2M uint64
	DirectMap1G uint64

	UsedMem int	// 通过 free 获取，以 KB 为单位
}

func All() MemInfo {
	var mi MemInfo
	GetMemInfo(&mi)
	GetUsedMem(&mi)

	return mi
}

func GetMemInfo(mi *MemInfo) {
	m, _ := proc.GetMeminfo()
	mi.MemTotal = m.MemTotal
	mi.MemFree = m.MemFree
	mi.MemAvailable = m.MemAvailable
	mi.Buffers = m.Buffers
	mi.Cached = m.Cached
	mi.SwapCached = m.SwapCached
	mi.Active = m.Active
	mi.Inactive = m.Inactive
	mi.ActiveAnon = m.ActiveAnon
	mi.InactiveAnon = m.InactiveAnon
	mi.ActiveFile = m.ActiveFile
	mi.InactiveFile = m.InactiveFile
	mi.Unevictable = m.Unevictable
	mi.Mlocked = m.Mlocked
	mi.SwapTotal = m.SwapTotal
	mi.SwapFree = m.SwapFree
	mi.Dirty = m.Dirty
	mi.Writeback = m.Writeback
	mi.AnonPages = m.AnonPages
	mi.Mapped = m.Mapped
	mi.Shmem = m.Shmem
	mi.Slab = m.Slab
	mi.SReclaimable = m.SReclaimable
	mi.SUnreclaim = m.SUnreclaim
	mi.KernelStack = m.KernelStack
	mi.PageTables = m.PageTables
	mi.NfsUnstable = m.NfsUnstable
	mi.Bounce = m.Bounce
	mi.WritebackTmp = m.WritebackTmp
	mi.CommitLimit = m.CommitLimit
	mi.CommittedAS = m.CommittedAS
	mi.VmallocTotal = m.VmallocTotal
	mi.VmallocUsed = m.VmallocUsed
	mi.VmallocChunk = m.VmallocChunk
	mi.HardwareCorrupted = m.HardwareCorrupted
	mi.AnonHugePages = m.AnonHugePages
	mi.CmaTotal = m.CmaTotal
	mi.CmaFree = m.CmaFree
	mi.HugePagesTotal = m.HugePagesTotal
	mi.HugePagesFree = m.HugePagesFree
	mi.HugePagesRsvd = m.HugePagesRsvd
	mi.HugePagesSurp = m.HugePagesSurp
	mi.Hugepagesize = m.Hugepagesize
	mi.DirectMap4k = m.DirectMap4k
	mi.DirectMap2M = m.DirectMap2M
	mi.DirectMap1G = m.DirectMap1G	
}

func GetUsedMem(mi *MemInfo) {
	out, err := exec.Command("sh", "-c", `free | awk 'NR==2 {print $3}'`).Output()
	if err != nil {
		fmt.Println("GetMemInfo", err)
		return
	}
	mi.UsedMem, err = strconv.Atoi(string(out[0:len(out)-1]))
	if err != nil {
		fmt.Println("GetUsedMem: strconv.Atoi: ", err)
	}
}