package pid

import (
	"strconv"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type SchedStatInfo struct {
	SumExecRuntime	uint64	// time spent on the cpu, task->se.sum_exec_runtime, 累计运行的物理时间时间, se.sum_exec_runtime
	WaitSum			uint64	// io 等待时间, se.wait_sum
	Switches		uint64	// 主动切换和被动切换的累计次数, se->nr_switches
}

func SchedStat(pid string) (SchedStatInfo, error) {
	var info SchedStatInfo

	buf, err := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "schedstat"))
	if err != nil {
		return info, err
	}

	vs := strings.Fields(string(buf))
	if len(vs) < 3 {
		return info, fmt.Errorf("invalid schedstat file: %v\n", string(buf))
	}

	info.SumExecRuntime, _ = strconv.ParseUint(vs[0], 10, 64)
	info.WaitSum, _ = strconv.ParseUint(vs[1], 10, 64)
	info.Switches, _ = strconv.ParseUint(strings.TrimSpace(vs[2]), 10, 64)

	return info, nil
}