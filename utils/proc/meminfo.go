package proc

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"
)

// unit of mem info is KB
type Meminfo struct {
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
}

func GetMeminfo() (Meminfo, error) {
	var mi Meminfo

	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return mi, fmt.Errorf("open /proc/meminfo failed\n")
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// type bytes
		tb, e := r.ReadBytes(':')
		if e == io.EOF {
			break;
		}
		//type string
		ts := string(tb[:len(tb)-1])

		vb, _ := r.ReadBytes('\n')
		vs := string(vb)
		vs = strings.TrimSpace(vs);
		if vs[len(vs)-1] == 'B' {
			vs = vs[:len(vs)-3]
		}
		v, _ := strconv.ParseUint(vs, 10, 64) 

		switch ts {
		case "MemTotal":
			mi.MemTotal = v
		case "MemFree":
			mi.MemFree = v
		case "MemAvailable":
			mi.MemAvailable = v
		case "Buffers":
			mi.Buffers = v
		case "Cached":
			mi.Cached = v
		case "SwapCached":
			mi.SwapCached = v
		case "Active":
			mi.Active = v
		case "Inactive":
			mi.Inactive = v
		case "Active(anon)":
			mi.ActiveAnon = v 
		case "Inactive(anon)":
			mi.InactiveAnon = v 
		case "Active(file)":
			mi.ActiveFile = v 
		case "Inactive(file)":
			mi.InactiveFile = v 
		case "Unevictable":
			mi.Unevictable = v 
		case "Mlocked":
			mi.Mlocked = v 
		case "SwapTotal":
			mi.SwapTotal = v 
		case "SwapFree":
			mi.SwapFree = v 
		case "Dirty":
			mi.Dirty = v 
		case "Writeback":
			mi.Writeback = v 
		case "AnonPages":
			mi.AnonPages = v 
		case "Mapped":
			mi.Mapped = v 
		case "Shmem":
			mi.Shmem = v 
		case "Slab":
			mi.Slab = v 
		case "SReclaimable":
			mi.SReclaimable = v 
		case "SUnreclaim":
			mi.SUnreclaim = v 
		case "KernelStack":
			mi.KernelStack = v 
		case "PageTables":
			mi.PageTables = v 
		case "NFS_Unstable":
			mi.NfsUnstable = v 
		case "Bounce":
			mi.Bounce = v 
		case "WritebackTmp":
			mi.WritebackTmp = v 
		case "CommitLimit":
			mi.CommitLimit = v 
		case "Committed_AS":
			mi.CommittedAS = v 
		case "VmallocTotal":
			mi.VmallocTotal = v 
		case "VmallocUsed":
			mi.VmallocUsed = v 
		case "VmallocChunk":
			mi.VmallocChunk = v 
		case "HardwareCorrupted":
			mi.HardwareCorrupted = v 
		case "AnonHugePages":
			mi.AnonHugePages = v 
		case "CmaTotal":
			mi.CmaTotal = v 
		case "CmaFree":
			mi.CmaFree = v 
		case "HugePages_Total":
			mi.HugePagesTotal = v 
		case "HugePages_Free":
			mi.HugePagesFree = v 
		case "HugePages_Rsvd":
			mi.HugePagesRsvd = v 
		case "HugePages_Surp":
			mi.HugePagesSurp = v 
		case "Hugepagesize":
			mi.Hugepagesize = v 
		case "DirectMap4k":
			mi.DirectMap4k = v 
		case "DirectMap2M":
			mi.DirectMap2M = v 
		default:
			fmt.Println("Not record:", ts)
		}
	}

	return mi, nil
}
