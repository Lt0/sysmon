package mem

import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
)

// unit of mem info is KB
type MemInfo struct {
	MemTotal int
	MemFree int
	MemAvailable int
	Buffers int
	Cached int
	SwapCached int
	Active int
	Inactive int
	ActiveAnon int
	InactiveAnon int
	ActiveFile int
	InactiveFile int
	Unevictable int
	Mlocked int
	SwapTotal int
	SwapFree int
	Dirty int
	Writeback int
	AnonPages int
	Mapped int
	Shmem int
	Slab int
	SReclaimable int
	SUnreclaim int
	KernelStack int
	PageTables int
	NfsUnstable int
	Bounce int
	WritebackTmp int
	CommitLimit int
	CommittedAS int
	VmallocTotal int
	VmallocUsed int
	VmallocChunk int
	HardwareCorrupted int
	AnonHugePages int
	CmaTotal int
	CmaFree int
	HugePagesTotal int
	HugePagesFree int
	HugePagesRsvd int
	HugePagesSurp int
	Hugepagesize int
	DirectMap4k int
	DirectMap2M int
	DirectMap1G int

	UsedMem int	// 通过 free 获取，以 KB 为单位
}

func All() MemInfo {
	var mi MemInfo
	GetMemInfo(&mi)
	GetUsedMem(&mi)

	return mi
}

func GetMemInfo(mi *MemInfo) {
	out, err := exec.Command("sh", "-c", `cat /proc/meminfo | sed 's/:\s*/:/g' | sed 's/ kB$//g'`).Output()
	if err != nil {
		fmt.Println("GetMemInfo", err)
		return
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrStr := strings.Split(v, ":")
		switch attrStr[0] {
		case "MemTotal":
			mi.MemTotal, _ = strconv.Atoi(attrStr[1]) 
		case "MemFree":
			mi.MemFree, _ = strconv.Atoi(attrStr[1]) 
		case "MemAvailable":
			mi.MemAvailable, _ = strconv.Atoi(attrStr[1]) 
		case "Buffers":
			mi.Buffers, _ = strconv.Atoi(attrStr[1]) 
		case "Cached":
			mi.Cached, _ = strconv.Atoi(attrStr[1]) 
		case "SwapCached":
			mi.SwapCached, _ = strconv.Atoi(attrStr[1]) 
		case "Active":
			mi.Active, _ = strconv.Atoi(attrStr[1]) 
		case "Inactive":
			mi.Inactive, _ = strconv.Atoi(attrStr[1]) 
		case "Active(anon)":
			mi.ActiveAnon, _ = strconv.Atoi(attrStr[1]) 
		case "Inactive(anon)":
			mi.InactiveAnon, _ = strconv.Atoi(attrStr[1]) 
		case "Active(file)":
			mi.ActiveFile, _ = strconv.Atoi(attrStr[1]) 
		case "Inactive(file)":
			mi.InactiveFile, _ = strconv.Atoi(attrStr[1]) 
		case "Unevictable":
			mi.Unevictable, _ = strconv.Atoi(attrStr[1]) 
		case "Mlocked":
			mi.Mlocked, _ = strconv.Atoi(attrStr[1]) 
		case "SwapTotal":
			mi.SwapTotal, _ = strconv.Atoi(attrStr[1]) 
		case "SwapFree":
			mi.SwapFree, _ = strconv.Atoi(attrStr[1]) 
		case "Dirty":
			mi.Dirty, _ = strconv.Atoi(attrStr[1]) 
		case "Writeback":
			mi.Writeback, _ = strconv.Atoi(attrStr[1]) 
		case "AnonPages":
			mi.AnonPages, _ = strconv.Atoi(attrStr[1]) 
		case "Mapped":
			mi.Mapped, _ = strconv.Atoi(attrStr[1]) 
		case "Shmem":
			mi.Shmem, _ = strconv.Atoi(attrStr[1]) 
		case "Slab":
			mi.Slab, _ = strconv.Atoi(attrStr[1]) 
		case "SReclaimable":
			mi.SReclaimable, _ = strconv.Atoi(attrStr[1]) 
		case "SUnreclaim":
			mi.SUnreclaim, _ = strconv.Atoi(attrStr[1]) 
		case "KernelStack":
			mi.KernelStack, _ = strconv.Atoi(attrStr[1]) 
		case "PageTables":
			mi.PageTables, _ = strconv.Atoi(attrStr[1]) 
		case "NFS_Unstable":
			mi.NfsUnstable, _ = strconv.Atoi(attrStr[1]) 
		case "Bounce":
			mi.Bounce, _ = strconv.Atoi(attrStr[1]) 
		case "WritebackTmp":
			mi.WritebackTmp, _ = strconv.Atoi(attrStr[1]) 
		case "CommitLimit":
			mi.CommitLimit, _ = strconv.Atoi(attrStr[1]) 
		case "Committed_AS":
			mi.CommittedAS, _ = strconv.Atoi(attrStr[1]) 
		case "VmallocTotal":
			mi.VmallocTotal, _ = strconv.Atoi(attrStr[1]) 
		case "VmallocUsed":
			mi.VmallocUsed, _ = strconv.Atoi(attrStr[1]) 
		case "VmallocChunk":
			mi.VmallocChunk, _ = strconv.Atoi(attrStr[1]) 
		case "HardwareCorrupted":
			mi.HardwareCorrupted, _ = strconv.Atoi(attrStr[1]) 
		case "AnonHugePages":
			mi.AnonHugePages, _ = strconv.Atoi(attrStr[1]) 
		case "CmaTotal":
			mi.CmaTotal, _ = strconv.Atoi(attrStr[1]) 
		case "CmaFree":
			mi.CmaFree, _ = strconv.Atoi(attrStr[1]) 
		case "HugePages_Total":
			mi.HugePagesTotal, _ = strconv.Atoi(attrStr[1]) 
		case "HugePages_Free":
			mi.HugePagesFree, _ = strconv.Atoi(attrStr[1]) 
		case "HugePages_Rsvd":
			mi.HugePagesRsvd, _ = strconv.Atoi(attrStr[1]) 
		case "HugePages_Surp":
			mi.HugePagesSurp, _ = strconv.Atoi(attrStr[1]) 
		case "Hugepagesize":
			mi.Hugepagesize, _ = strconv.Atoi(attrStr[1]) 
		case "DirectMap4k":
			mi.DirectMap4k, _ = strconv.Atoi(attrStr[1]) 
		case "DirectMap2M":
			mi.DirectMap2M, _ = strconv.Atoi(attrStr[1]) 
		default:
		}
	}

	
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