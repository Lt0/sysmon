package net

import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
)

type Context struct {
	Procfs string
}

var Ctx Context

type NetInfo struct {
	NicNum int // 网卡数量
	Nics []NicInfo // 每个网卡的状态信息
}

type NicInfo struct {
	Name string

	RBytes int
	RPackets int
	RErrs int
	RDrop int
	RFifo int
	RFrame int
	RComporessed int
	RMulticast int

	TBytes int
	TPackets int
	TErrs int
	TDrop int
	TFifo int
	TFrame int
	TComporessed int
	TMulticast int
}

func GetNetInfo() NetInfo {
	var ni NetInfo

	cmd := fmt.Sprintf("cat %s/net/dev | tail -n +3 | sed 's/^\\s*//g' | sed 's/\\s\\+/|/g'", Ctx.Procfs)
	out, err := exec.Command("sh", "-c", cmd).Output()
	//out, err := exec.Command("sh", "-c", `cat /proc/net/dev | tail -n +3 | sed 's/^\s*//g' | sed 's/\s\+/|/g'`).Output()
	if err != nil {
		fmt.Println("GetCoresInfo", err)
	}
	
	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		NicStrs := strings.Split(v, ":")
		infos := strings.Split(NicStrs[1], "|")

		if len(infos) < 17 {
			fmt.Println("infos is too short, could not get NIC info")
			return ni
		}
		RBytes, _ := strconv.Atoi(infos[1])
		RPackets, _ := strconv.Atoi(infos[2])
		RErrs, _ := strconv.Atoi(infos[3])
		RDrop, _ := strconv.Atoi(infos[4])
		RFifo, _ := strconv.Atoi(infos[5])
		RFrame, _ := strconv.Atoi(infos[6])
		RComporessed, _ := strconv.Atoi(infos[7])
		RMulticast, _ := strconv.Atoi(infos[8])

		TBytes, _ := strconv.Atoi(infos[9])
		TPackets, _ := strconv.Atoi(infos[10])
		TErrs, _ := strconv.Atoi(infos[11])
		TDrop, _ := strconv.Atoi(infos[12])
		TFifo, _ := strconv.Atoi(infos[13])
		TFrame, _ := strconv.Atoi(infos[14])
		TComporessed, _ := strconv.Atoi(infos[15])
		TMulticast, _ := strconv.Atoi(infos[16])
		ni.Nics = append(ni.Nics, NicInfo{NicStrs[0], RBytes, RPackets, RErrs, RDrop, RFifo, RFrame, RComporessed, RMulticast, TBytes, TPackets, TErrs, TDrop, TFifo, TFrame, TComporessed, TMulticast})
	}

	return ni
}