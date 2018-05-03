package cpu

import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
)

type CpuInfo struct {
	CoreNum int
	OnlineCore string
	Arch string
	CpuOpMode string
	ByteOrder string
	ThreadPerCore int
	CorePerSocket int
	Socket int
	NumaNodeNum int
	VendorId string
	CpuFamily string
	ModelName string
	Freq string
	HypervisorVendor string
	Virtualization string
	L1dCache string
	L1iCache string
	L2Cache string
	L3Cache string
	NumaNodes []NumaNode
	Cores []Core
}

type Core struct {
	Name string
	User int
	Nice int
	System int
	Idle int
	Iowait int
	Irq int
	Softirq int
	Steal int
	Guest int
	GuestNide int
}

type NumaNode struct {
	Name string
	Cpus string
}

func GetCpuInfo() CpuInfo {
	var ci CpuInfo
	GetCoresInfo(&ci)
	GetGlobalInfo(&ci)
	GetModelName(&ci)
	return ci
}

func GetCoresInfo(ci *CpuInfo){
	cmd := fmt.Sprintf("/bin/cat /proc/stat | /bin/grep \"^cpu[0-9]\"")
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("GetCoresInfo", err)
	}
	
	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrStr := strings.Split(v, " ")
		var a []int
		for i := 1; i < len(attrStr); i++ {
			attr, _ := strconv.Atoi(attrStr[i])
			a = append(a, attr)
		}
		ci.Cores = append(ci.Cores, Core{Name:attrStr[0], User:a[0], Nice:a[1], System:a[2], Idle:a[3], Iowait:a[4], Irq:a[5], Softirq:a[6]})
		if len(a) >= 8 {ci.Cores[len(ci.Cores)-1].Steal = a[7]}
		if len(a) >= 9 {ci.Cores[len(ci.Cores)-1].Guest = a[8]}
		if len(a) >= 10 {ci.Cores[len(ci.Cores)-1].GuestNide = a[9]}
	}

	ci.CoreNum = len(ci.Cores)
}

func GetGlobalInfo(ci *CpuInfo){
	out, err := exec.Command("sh", "-c", `lscpu | sed 's/:\s*/:/g'`).Output()
	if err != nil {
		fmt.Println("GetStaticInfo", err)
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrStr := strings.Split(v, ":")
		switch attrStr[0] {
		case "Architecture":
			ci.Arch = attrStr[1]
		case "CPU op-mode(s)":
			ci.CpuOpMode = attrStr[1]
		case "Byte Order":
			ci.ByteOrder = attrStr[1]
		case "On-line CPU(s) list":
			ci.OnlineCore = attrStr[1]
		case "Thread(s) per core":
			ci.ThreadPerCore, _ = strconv.Atoi(attrStr[1])
		case "Core(s) per socket":
			ci.CorePerSocket, _ = strconv.Atoi(attrStr[1])
		case "Socket(s)":
			ci.Socket, _ = strconv.Atoi(attrStr[1])
		case "NUMA node(s)":
			ci.NumaNodeNum, _ = strconv.Atoi(attrStr[1])
		case "Vendor ID":
			ci.VendorId = attrStr[1]
		case "CPU family":
			ci.CpuFamily = attrStr[1]
		case "CPU MHz":
			ci.Freq = attrStr[1]
		case "Hypervisor vendor":
			ci.HypervisorVendor = attrStr[1]
		case "Virtualization type":
			ci.Virtualization = attrStr[1]
		case "L1d cache":
			ci.L1dCache = attrStr[1]
		case "L1i cache":
			ci.L1iCache = attrStr[1]
		case "L2 cache":
			ci.L2Cache = attrStr[1]
		case "L3 cache":
			ci.L3Cache = attrStr[1]
		case "NUMA node0 CPU(s)": fallthrough
		case "NUMA node1 CPU(s)": fallthrough
		case "NUMA node2 CPU(s)": fallthrough
		case "NUMA node3 CPU(s)": 
			ci.NumaNodes = append(ci.NumaNodes, NumaNode{Name:attrStr[0], Cpus:attrStr[1]})
		default:
		} // end of switch 
	} // end of for
}

func GetModelName(ci *CpuInfo) {
	out, err := exec.Command("sh", "-c", `cat /proc/cpuinfo  | grep "model name" | head -n1 | sed 's/^.*: //'`).Output()
	if err != nil {
		fmt.Println("GetStaticInfo", err)
	}
	ci.ModelName = string(out)
}