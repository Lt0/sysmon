package disk

import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
)

// disk stats
type DiskInfo struct {
	Storage			[]StorageInfo	// 每一个分区存储为一个 storage
	Partitions		[]PartitionStat	// 按分区统计
}

// 存储状态
type StorageInfo struct {
	Fs			string
	Type		string
	Size		string
	Used		string
	Avail		string
	UsePercent	string
	MountPoint	string
}

type PartitionStat struct {
	MajorNum 		int		// device id, 此块设备的主设备号
	MinorNum 		int		// partition id, 此块设备的次设备号
	Name 			string	// device name, 此块设备名字
	
	ReadsCompleted 	uint64	// 成功完成的读请求次数
	ReadsMerged 	uint64	// 读请求的次数
	SectorsRead		uint64	// 读请求的扇区数总和
	ReadTime		uint64	// 读请求花费的时间总和

	WritesCompleted	uint64	// 成功完成的写请求次数
	WritesMerged	uint64	// 写请求合并的次数
	SectorsWrite	uint64	// 写请求的扇区数总和
	WriteTime		uint64	// 写请求花费的时间总和

	QueueIOs		uint64	// 次块设备队列中的IO请求数
	IOTime			uint64	// 块设备队列非空时间总和
	IOWeightedTime	uint64	// 块设备队列非空时间加权读请求的扇区数总和总和

	PhySector		uint64	// 物理 sector 大小, 默认为 512，通过 BlkInfo() 修正
	LogSector		uint64	// 逻辑 sector 大小, 默认为 512，通过 BlkInfo() 修正
	IsMajorHD		bool	// 是否为物理硬盘，默认为 false，通过 BlkInfo() 修正
}

func AllInfo() DiskInfo {
	var di DiskInfo

	Storage(&di)
	Stats(&di)

	BlkInfo(&di)

	return di
}

func Storage(di *DiskInfo) {
	out, err := exec.Command("sh", "-c", `df -Th -x tmpfs -x devtmpfs | tail -n +2 | sort | sed 's/^\s*//g' | sed 's/\s\+/|/g'`).Output()
	if err != nil {
		fmt.Println("Disk Storage", err)
		return
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrs := strings.Split(v, "|")		
		di.Storage = append(di.Storage, StorageInfo{attrs[0], attrs[1], attrs[2], attrs[3], attrs[4], attrs[5], attrs[6]})
	}
}

func Stats(di *DiskInfo){
	out, err := exec.Command("sh", "-c", `cat /proc/diskstats | sed 's/^\s*//g' | sed 's/\s\+/ /g'`).Output()
	if err != nil {
		fmt.Println("Disk Stats", err)
		return
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrs := strings.Split(v, " ")

		// 只获取物理硬盘的数据
		if (attrs[0] != "8"){
			continue;
		}

		MajorNum, _ := strconv.Atoi(attrs[0])
		MinorNum, _ := strconv.Atoi(attrs[1])
		Name := attrs[2]

		ReadsCompleted, _ := strconv.ParseUint(attrs[3], 10, 64)
		ReadsMerged, _ := strconv.ParseUint(attrs[4], 10, 64)
		SectorsRead, _ := strconv.ParseUint(attrs[5], 10, 64)
		ReadTime, _ := strconv.ParseUint(attrs[6], 10, 64)

		WritesCompleted, _ := strconv.ParseUint(attrs[7], 10, 64)
		WritesMerged, _ := strconv.ParseUint(attrs[8], 10, 64)
		SectorsWrite, _ := strconv.ParseUint(attrs[9], 10, 64)
		WriteTime, _ := strconv.ParseUint(attrs[10], 10, 64)
		
		QueueIOs, _ := strconv.ParseUint(attrs[11], 10, 64)
		IOTime, _ := strconv.ParseUint(attrs[12], 10, 64)
		IOWeightedTime, _ := strconv.ParseUint(attrs[13], 10, 64)

		di.Partitions = append(di.Partitions, PartitionStat{MajorNum, MinorNum, Name, 
			ReadsCompleted, ReadsMerged, SectorsRead, ReadTime, 
			WritesCompleted, WritesMerged, SectorsWrite, WriteTime, 
			QueueIOs, IOTime, IOWeightedTime, 512, 512, false})
	}
}

func BlkInfo(di *DiskInfo){
	out, err := exec.Command("sh", "-c", `lsblk  -tdn | sed 's/^\s*//g' | sed 's/\s\+/|/g'`).Output()
	if err != nil {
		fmt.Println("Disk Storage", err)
		return
	}

	for _, v := range(strings.Split(string(out), "\n")) {
		if v == "" {continue}
		attrs := strings.Split(v, "|")
		name := attrs[0]
		for i := 0; i < len(di.Partitions); i++ {
			if (di.Partitions[i].Name == name){
				di.Partitions[i].PhySector, _ = strconv.ParseUint(attrs[4], 10, 64)
				di.Partitions[i].LogSector, _ = strconv.ParseUint(attrs[5], 10, 64)
				di.Partitions[i].IsMajorHD = true
				break
			}
		}
	}
}