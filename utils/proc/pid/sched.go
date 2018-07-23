// 注下面的时间或时刻都是从rq->clock中获得的，而这个值是由update_rq_clock底层cpu来更新的。并且很多信息是需要内核配置CONFIG_SCHEDSTATS才有。
// 大多数字段的计算在sched.c及sched_fair.c里，在这两个文件里搜索相应的字段就能得到相应的计算方法。
// 
// cpu_test (28733, #threads: 1)
// ---------------------------------------------------------
// se.exec_start       :    2781299327.397282  //此进程最近被调度到的开始执行时刻（这个值是每次update_curr都进行更新）
// se.vruntime        :       3144603.079903  //虚拟运行时间
// se.sum_exec_runtime:       2843625.998498  //累计运行的物理时间时间
// se.wait_start       :             0.000000  //最近一次当前进程被入队的时刻
// se.sleep_start      :             0.000000  //此进程最近一次被从队列里取出，并被置S状态的时刻
// se.block_start      :             0.000000  //此进程最近一次被从队列里取出，并被置D状态的时刻
// se.sleep_max      :             0.000000  //最长处于S状态时间
// se.block_max      :             0.000000  //最长处于D状态时间
// se.exec_max       :             1.004266  //最长单次执行时间
// se.slice_max       :           998.456300  //曾经获得时间片的最长时间
// se.wait_max       :             0.455235  //最长在就绪队列里的等待时间
// se.wait_sum       :            15.615407  //累计在就绪队列里的等待时间
// se.wait_count      :                 3147  //累计等待次数
// se.iowait_sum      :           215.825267  //io等待时间
// se.iowait_count     :                   67 //io等待次数  io_schedule调用次数
// sched_info.bkl_count:                    0  //此进程大内核锁调用次数
// se.nr_migrations    :                    0 //需要迁移当前进程到其他cpu时累加此字段
// se.nr_migrations_cold:                    0
// se.nr_failed_migrations_affine:           194  //进程设置了cpu亲和，进程迁移时检查失败的次数
// se.nr_failed_migrations_running:           0 
// se.nr_failed_migrations_hot:               0  //当前进程因为是cache hot导致迁移失败的次数
// se.nr_forced_migrations :                  0  //在当前进程cache hot下，由于负载均衡尝试多次失败，强行进行迁移的次数
// se.nr_wakeups         :                 0  //被唤醒的累计次数（从不可运行到可运行）
// se.nr_wakeups_sync     :                0  //同步唤醒次数，即a唤醒b，a立刻睡眠，b被唤醒的次数
// se.nr_wakeups_migrate  :                 0 //被唤醒得到调度的当前cpu，不是之前睡眠的cpu的次数
// se.nr_wakeups_local     :                0 //被本地唤醒的次数（唤醒后在当前cpu上执行）
// se.nr_wakeups_remote   :                0 //非本地唤醒累计次数
// se.nr_wakeups_affine    :                0 //考虑了任务的cache亲和性的唤醒次数
// se.nr_wakeups_affine_attempts:            0
// se.nr_wakeups_passive  :                    0
// se.nr_wakeups_idle     :                    0
// avg_atom            :           903.886204 //本进程平均耗时sum_exec_runtime/ nr_switches
// avg_per_cpu           :             0.000001
// nr_switches            :                 3146 //主动切换和被动切换的累计次数
// nr_voluntary_switches   :                    0 //主动切换次数（由于prev->state为不可运行状态引起的切换）
// nr_involuntary_switches  :                 3146 //被动切换次数
// se.load.weight          :                 1024  //该se的load
// policy                 :                    0  //调度策略 normal
// prio                   :                  120  //优先级(nice=0)
// clock-delta             :                   51
package pid

import (
	"strings"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type SchedInfo struct {
	Items	[]SchedItem
	Other	string
}

type SchedItem struct {
	Name	string
	Val		string
}

func Sched(pid string) (SchedInfo, error) {
	var info SchedInfo

	buf, err := ioutil.ReadFile(filepath.Join(procfs, pid, "sched"))
	if err != nil {
		return info, err
	}
	lines := strings.Split(string(buf), "\n")
	if len(lines) < 3 {
		return info, fmt.Errorf("Invalid Sched content: %v\n", string(buf))
	}

	for _, v := range(lines[2:]) {
		if strings.Contains(v, ":") {
			vs := strings.Split(v, ":")
			if len(vs) == 2 {
				var item SchedItem
				item.Name = strings.TrimSpace(vs[0])
				item.Val = strings.TrimSpace(vs[1])
				info.Items = append(info.Items, item)
			} else {
				fmt.Println("Sched: Invalid entry:", v)
			}
		} else {
			info.Other += (v + "\n")
		}
	}

	return info, nil
}

func SchedRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(procfs, pid, "sched"))
	return string(buf)
}

