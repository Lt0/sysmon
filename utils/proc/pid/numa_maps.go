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

// containe all info of /proc/[pid]/numa_maps
type NumaMapsInfo struct {
	Mappings	[]Mapping
}

type Mapping struct {
	Addr		string	// 内存范围的起始地址
	Policy		string	// 对该范围内存生效的内存策略，请注意，有效策略不一定是该内存范围的进程安装的策略。 具体而言，如果进程为该范围安装了“default”策略，则该范围的有效策略将是进程策略，其可能是也可能不是“default”。
	Nodes		[]NodePages	// 各个 NUMA 节点上分配的内存页数。
	File		string	// 内存范围所对应的文件。 如果文件被映射为私有，则写访问可能已在此内存范围中生成COW（写时复制）页。 这些页面显示为匿名页面。
	MapType		string	// 该字段可能为空，以下是可能的类型：Heap: 由堆使用的内存; Stack: 由栈使用的内存; Huge: Huge 内存范围，显示的是大页内存而不是常规页。
	Anon		int		// 该范围内的匿名内存页数
	Dirty		int		// 需要回写磁盘的页数
	Mapped		int		// 映射的内存页面总数（如果和 dirty 以及 anon 的页数不同的话）。
	MapMax		int		// 扫描期间遇到的最大mapcount（映射单个页面的进程数）。 这可以用作在给定存储器范围内发生的共享程度的指示符。
	SwapCache	int		// 在交换设备上具有关联条目的页数。
	Active		int		// 活动列表中的页数。仅当与此范围内的页数不同时，才会显示此字段。 这意味着存储器范围中存在一些非活动页面，这些页面很快就会被交换器从内存中删除。
	WriteBack	int		// 当前正在回写到磁盘的页数

	KernelPageSizeKB string	// 内核页大小
}

// 指定的 NUMA 节点上分配的内存页数。
type NodePages struct {
	Node	string	// numa 节点名
	NrPages	string	// nr_pages 只包含当前进程映射的内存。页面迁移和内存回收的临时内存可能不会映射在这里
}

func NumaMaps(pid string) (NumaMapsInfo, error) {
	var info NumaMapsInfo

	f, err := os.Open(filepath.Join(Ctx.Procfs, pid, "numa_maps"))
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

		s := string(b)
		item, err := parseNumaMapLine(s)
		if err != nil {
			fmt.Printf("NumaMaps: parsing %v failed: %v\n", s, err)
			continue
		}
		info.Mappings = append(info.Mappings, item)
	}

	return info, nil
}

func parseNumaMapLine(s string) (Mapping, error) {
	var mapping Mapping
	
	vs := strings.Fields(s)
	if len(vs) < 2 {
		return mapping, fmt.Errorf("parseNumaMapLine: invalid line: %s\n", s)
	}

	mapping.Addr = vs[0]
	mapping.Policy = vs[1]
	for _, v := range(vs[2:]) {
		v = strings.TrimSpace(v)
		kvpaire := strings.Split(v, "=")

		// kvpaire == 1 means kvpaire[0] is MapType
		if len(kvpaire) == 1 {
			mapping.MapType = kvpaire[0]
			continue
		}
		
		switch kvpaire[0] {
		case "file":
			mapping.File = kvpaire[1]
		case "anon":
			mapping.Anon, _ = strconv.Atoi(kvpaire[1])
		case "dirty":
			mapping.Dirty, _ = strconv.Atoi(kvpaire[1])
		case "mapped":
			mapping.Mapped, _ = strconv.Atoi(kvpaire[1])
		case "mapmax":
			mapping.MapMax, _ = strconv.Atoi(kvpaire[1])
		case "swapcache":
			mapping.SwapCache, _ = strconv.Atoi(kvpaire[1])
		case "active":
			mapping.Active, _ = strconv.Atoi(kvpaire[1])
		case "writeback":
			mapping.WriteBack, _ = strconv.Atoi(kvpaire[1])
		case "kernelpagesize_kB":
			mapping.KernelPageSizeKB = kvpaire[1]
		default:
			node, err := parseNumaMapNode(kvpaire)
			if err == nil {
				mapping.Nodes = append(mapping.Nodes, node)
			}
		}
	}

	return mapping, nil
}

func parseNumaMapNode(vs []string) (NodePages, error) {
	var node NodePages
	if vs[0][0] != 'N' {
		return node, fmt.Errorf("parseNumaMapNode: invalid vs: %v\n", vs)
	}
	
	node.Node = vs[0]
	node.NrPages = vs[1]
	return node, nil
}