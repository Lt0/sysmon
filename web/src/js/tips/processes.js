
let hdr = {
    Comm: "进程命令",
    CPU: "CPU 使用率，每个内核按 100% 计算",
    MEM: "内存使用率",
    CPUTime: "CPU 运行时长，表示当前进程所有消耗的计算资源，交给单个内核来承担的话，所需要的时间长度",
    TaskCPU: "当前运行在哪个内核上",
    RRate: "磁盘读取速率",
    WRate: "磁盘写入速率",
    VmSize: "虚拟地址空间大小，是进程所能访问的全部逻辑地址空间的大小",
    VmRSS: "进程当前实际占用的物理内存，包含 VmRSS = RssAnon + RssFile + RssShmem 三部分",
    VmPTE: "进程内存的所有页表所占的大小",
    VmSwap: "交换到 swap 分区的匿名私有页的大小（不包括 shmem 交换所使用的内存）",
    Pid: "进程 ID",
    Threads: "进程内的线程数",
    State: "进程状态（R: 运行, S: 睡眠, D: 不可中断的睡眠, Z: 僵尸, T: 处于被跟踪状态或停止状态, X: dead",
    Nice: "进程的静态优先级",
    Priority: "进程的动态优先级",
    User: "启动进程的用户",
    Uid: "启动进程的用户的 ID",
    Read: "进程从磁盘读取的数据总和，包括子线程",
    Write: "进程写入磁盘的数据总和，包括子线程",
    Cmdline: "完整的进程命令，包括参数等",
}

// let state = {
//     R: "正在运行",
//     S: "可中断的睡眠",
//     D: "不可中断的睡眠",
//     Z: "僵尸",
//     T: "处于被跟踪状态或停止状态（基于信号）",
//     W: "paging"
// }
let state = {
    "R (running)": "正在运行",
    "S (sleeping)": "可中断的睡眠",
    "D (disk sleep)": "不可中断的睡眠",
    "T (stopped)": "处于被停止状态（基于信号）",
    "Z (zombie)": "僵尸",
    "T (tracing stop)": "处于被跟踪状态（基于信号）",
    "X (dead)": "进程挂了",
}

let detailsTabs = {
    thread: "进程中的所有线程信息",
    stack: "所有线程的内部调用状态",
    smaps: "整个进程的内存映射信息",
    numaMaps: "整个进程的内存在 NUMA 节点之间的映射关系",
    other: "进程相关的其它信息"
}

let limits = {
    Limit: "所限制的资源",
    SoftLimit: "当前系统生效的设置值",
    HardLimit: "系统中所能设定的最大值",
    Units: "衡量单位",
}

let stack = {

}

// check details from man 5 proc(ubuntu 16.04)
// https://lkml.org/lkml/2012/10/22/481
let smapsHdrVmFlags = `VmFlags field represents the kernel flags associated with the particular virtual memory area in two letter encoded manner. The codes are the following: 
    RD: readable;  
    WR: writeable;  
    EX: executable; 
    SH: shared; 
    MR: may read; 
    MW: may write; 
    ME: may execute; 
    MS: may share; 
    GD: stack segment growns down; 
    PF: pure PFN range; 
    DW: disabled write to the mapped file; 
    LO: pages are locked in memory; 
    IO: memory mapped I/O area; 
    SR: sequential read advise provided; 
    RR: random read advise provided; 
    DC: do not copy area on fork; 
    DE: do not expand area on remapping; 
    AC: area is accountable; 
    NR: swap space is not reserved for the area; 
    HT: area uses huge tlb pages; 
    NL: non-linear mapping; 
    AR: architecture specific flag; 
    DD: do not include area into core dump; 
    MM: mixed map area; 
    HG: huge page advise flag; 
    NH: no-huge page advise flag; 
    MG: mergable advise flag; 
`
let smapsHdr = {
    File: "映射的文件，中括号括起来的表示不是真正的文件，[stack:N] 表示线程号为 N 的线程对应的栈在内存中的映射情况", 
    Size: "虚拟内存大小", 
    Rss: "实际占用的物理内存大小，使用的共享库所占用的内存会全部计算在里边, Rss = Shared_Clean + Shared_Dirty + Private_Clean + Private_Dirty ", 
    Pss: "实际占用的物理内存大小，对于该进程用到的共享库，会根据使用该库的进程数量，按比例显示该进程占用的内存", 
    SharedClean: "没有更改过的共享页大小，当发生换页时不用写回磁盘", 
    SharedDirty: "更改过的共享页大小，当发生换页时需要写回磁盘", 
    PrivateClean: "没有更改过的私有页大小，当发生换页时不用写回磁盘", 
    PrivateDirty: "更改过的私有页大小，当发生换页时需要写回磁盘", 
    Referenced: "", 
    Anonymous: "匿名页大小，包含 AnonHugePages",  
    AnonHugePages: "Transparent HugePages (THP)，THP 与 Hugepages 不是一回事，Hugepages 在 /proc/meminfo中是被独立统计的，而 AnonHugePages 是被包含在 AnonPages 之中的，与进程的 RSS/PSS 是有重叠的", 
    SharedHugetlb: "Shared Hugetlb",
    PrivateHugetlb: "Private Hugetlb", 
    Swap: "交换到 swap 分区的大小", 
    SwapPss: "Swap Pss", 
    KernelPageSize: "内核的内存页面大小", 
    MMUPageSize: "体系结构 MMU 一个页面大小", 
    Locked: "被 mlock() 系统调用锁定的内存大小。被锁定的内存因为不能 pageout/swapout，会从 Active/Inactive LRU list 移到 Unevictable LRU list 上。", 
    VmFlags: smapsHdrVmFlags, 
    StartAddr: "映射的起始虚拟地址", 
    EndAddr: "映射的结束虚拟地址", 
    Perm: "虚拟内存的权限，r=读, w=写, x=可执行, s=共享, p=私有", 
    Offset: "偏移量,如果这段内存是从文件里映射过来的，则偏移量为这段内容在文件中的偏移量。如果不是从文件里面映射过来的则为 0", 
    Dev: "映像文件所在的磁盘的主设备号和次设备号", 
    INode: "映像文件的的磁盘节点（inode）", 
}

let numaMapsHdr = {
    Addr: "内存范围的起始地址", 
    Policy: "对该范围内存生效的内存策略，请注意，有效策略不一定是该内存范围的进程安装的策略。 具体而言，如果进程为该范围安装了“default”策略，则该范围的有效策略将是进程策略，其可能是也可能不是“default”。", 
    File: "内存范围所对应的文件。 如果文件被映射为私有，则写访问可能已在此内存范围中生成COW（写时复制）页。 这些页面显示为匿名页面", 
    Nodes: "各个 NUMA 节点上分配的内存页数。", 
    MapType: "内存用途，该字段可能为空，以下是可能的类型：Heap: 由堆使用的内存; Stack: 由栈使用的内存; Huge: Huge 内存范围，显示的是大页内存而不是常规页。", 
    Anon: "该范围内的匿名内存页数", 
    Dirty: "需要回写磁盘的页数", 
    Mapped: "映射的内存页面总数（如果和 dirty 以及 anon 的页数不同的话）", 
    MapMax: "扫描期间遇到的最大 mapcount（映射单个页面的进程数）。 这可以用作在给定存储器范围内发生的共享程度的指示符", 
    SwapCache: "在交换设备上具有关联条目的页数", 
    Active: "活动列表中的页数。仅当与此范围内的页数不同时，才会显示此字段。 这意味着存储器范围中存在一些非活动页面，这些页面很快就会被交换器从内存中删除", 
    WriteBack: "当前正在回写到磁盘的页数", 
    KernelPageSizeKB: "内核页大小", 
}

let oomAdj = `在内存不足（OOM）的情况下，系统根据各个进程的该文件的值来权衡要杀死的进程。 
内核使用这个进行进程的 oom_score 值的位移操作：有效值在 -16 到 +15 的范围内，再加上特殊值 -17，它会为此进程完全禁用OOM查杀。 
正分会增加这个过程被 OOM 杀手杀死的可能性; 负分数降低了可能性。
该文件的默认值为0; 新进程继承其父进程的 oom_adj 设置。 进程必须具有特权（CAP_SYS_RESOURCE）才能更新此文件。
从 Linux 2.6.36 开始，不推荐使用此文件，而使用 /proc/[pid]/oom_score_adj。
`

let oomScore = `此文件显示内核为此进程提供的当前分数，以便为OOM杀手选择进程。 得分越高意味着 OOM-Killer 更有可能选择该过程。 此分数的取决于进程运行过程使用的内存量，决策因素如下：
1. 该进程是否使用fork(2) (+)创建了很多子进程;
2. 进程是否已运行很长时间，或者是否占用了大量CPU时间 (-);
3. 该进程是否具有低的 nice 值（即> 0）(+));
4. 流程是否具有特权 (-); 和
5. 该进程是否进行直接硬件访问 (-)。
oom_score 还反映了进程的 oom_score_adj 或 oom_adj 设置指定的调整。`

let oomScoreAdj = `本文件可用于调整进程在 out-of-memory 的情况下被杀死的权重, 
该文件的有效值为 0(任何情况下都不 kill 该进程)~1000(一旦内存不足立即 kill 该进程)，`

let loginUid = `用户登录 ID
如果一个进程不是由任何登陆的用户启动的（如开机启动的服务进程）
则不会设置该进程的 login uid，其值为 4294967295（无符号整数的 -1 表示）
用户登陆时，pam_loginuid 模块在登录时（通过 tty/DM/ssh 等方式）会将进程的该文件设置为登陆的 uid，并且子进程会保留此值"`

let sessionID = `用户登录 ID
如果一个进程不是由任何登陆的用户启动的（如开机启动的服务进程）
则不会设置该进程的 session id，其值为 4294967295（无符号整数的 -1 表示）`

let setGroup = `allow: 允许包含进程 pid 的用户命名空间中的进程使用 setgroups（2）系统调用
deny: 在该用户命名空间中不允许 setgroups（2）
请注意，如果没有设置 /proc/[pid]/gid_map，则无论该文件的值是什么，都不允许调用setgroups（2）`

let cpuSet= `进程 cpuset 目录相对于 cpuset 根文件系统的路径`

let coredumpFilter = `从 2.6.23 内核开始，该文件用来控制在进程发生 crash 的时候，具体需要 dump 哪些内存的信息。
该文件的值是内存映射类型的比特掩码，如果对应的比特位被设置为 1，则该类型的内存映射的内容将会被 dump 出来。以下是各个比特位对应的类型：
bit 0  Dump anonymous private mappings.
bit 1  Dump anonymous shared mappings.
bit 2  Dump file-backed private mappings.
bit 3  Dump file-backed shared mappings.
bit 4  Dump ELF headers. (since Linux 2.6.24)
bit 5  Dump private huge pages. (since Linux 2.6.28)
bit 6  Dump shared huge pages. (since Linux 2.6.28)
bit 7  Dump private DAX pages. (since Linux 4.4)
bit 8  Dump shared DAX pages.(since Linux 4.4)
       
默认情况下，只会设置 0, 1, 4(内核启用 CONFIG_CORE_DUMP_DEFAULT_ELF_HEADERS)， 5 这几个比特位。其默认值可以通过 coredump_filter 启动选项来修改。`

let personality = `该只读文件公开进程的执行域，由 personality(2) 设置。 该值以十六进制表示法显示。
访问该文件的权限由 ptrace 访问模式 PTRACE_MODE_ATTACH_FSCREDS 检查控制; 见 ptrace（2）。`

let cwd = `进程的当前运行目录`

let schedHdr = `下面的时间或时刻都是从 rq->clock 中获得的，而这个值是由 update_rq_clock 底层 cpu 来更新的。并且很多信息是需要内核配置 CONFIG_SCHEDSTATS 才有。
大多数字段的计算在 sched.c 及 sched_fair.c 里，在这两个文件里搜索相应的字段就能得到相应的计算方法。`

let sched = {
    "se.exec_start": "进程最近一次被调度的时间（起始时间点），单位 ns", 
    "se.vruntime": "虚拟运行时间，cfs 调度用", 
    "se.sum_exec_runtime": "处于运行状态的时长的总时长", 
    "se.statistics.sum_sleep_runtime": "处于 sleep 状态的总时长", 
    "se.avg_overlap": "", 
    "se.avg_wakeup": "当前进程最近一次调用 try_to_wake_up 为止的单次执行时间", 
    "se.avg_running": "平均单次执行时间", 
    "se.statistics.wait_start": "最近一次进入 wait 队列的时刻", 
    "se.statistics.sleep_start": "最近一次被设置为 S 状态的时刻", 
    "se.statistics.block_start": "最近一次被设置为 D 状态的时刻", 
    "se.statistics.sleep_max": "处于 S 状态时间最长的一次的时长", 
    "se.statistics.block_max": "处于 D 状态时间最长的一次的时长", 

    "se.statistics.exec_max": "最长单次执行时间", 
    "se.statistics.slice_max": "曾经获得时间片的最长时间（当 cpu load 过高时开始计算）", 
    "se.statistics.wait_max": "最长在就绪队列里的等待时间", 
    "se.statistics.wait_sum": "累计在就绪队列里的等待时间", 
    "se.statistics.wait_count": "累计等待次数 （被出队的次数） cfs使用，当被选中做下一个待运行进程时（set_next_entity），等待结束", 
    "se.statistics.iowait_sum": "IO 等待总时长", 
    "se.statistics.iowait_count": "IO 等待次数，io_schedule 调用次数", 
    "sched_info.bkl_count": "此进程大内核锁调用次数", 
    "se.nr_migrations": "需要迁移当前进程到其他 cpu 时累加此字段", 

    "se.statistics.nr_migrations_cold": "2.6.32 代码里赋值都是 0", 
    "se.statistics.nr_failed_migrations_affine": "进程设置了 cpu 亲和，进程迁移时检查失败的次数", 
    "se.statistics.nr_failed_migrations_running": "由于进程处于 runing 状态，不运行迁移的次数", 
    "se.statistics.nr_failed_migrations_hot": "当前进程因为是 cache hot 导致迁移失败的次数", 
    "se.statistics.nr_forced_migrations": "在当前进程 cache hot 下，由于负载均衡尝试多次失败，强行进行迁移的次数", 
    "se.statistics.nr_wakeups": "被唤醒的次数", 
    "se.statistics.nr_wakeups_sync": "同步唤醒次数，即 a 唤醒 b，a 立刻睡眠，b 被唤醒的次数 /* waker goes to sleep after wakup */", 
    "se.statistics.nr_wakeups_migrate": "被唤醒得到调度的当前 cpu，不是之前睡眠的 cpu 的次数", 
    "se.statistics.nr_wakeups_local": "被本地唤醒的次数", 
    "se.statistics.nr_wakeups_remote": "远程唤醒累计次数", 
    "se.statistics.nr_wakeups_affine": "考虑了任务的 cache 亲和性的唤醒次数", 
    "se.statistics.nr_wakeups_affine_attempts": "尝试进行考虑了任务的 cache 亲和性的唤醒次数", 
    "se.statistics.nr_wakeups_passive": "2.6.32 代码里赋值都是 0", 
    "se.statistics.nr_wakeups_idle": "2.6.32 代码里赋值都是 0", 
    "avg_atom": "进程平均切换耗时", 
    "avg_per_cpu": "如果本进程曾经被推或者拉到其他 cpu 上，则计算每个 cpu 上的平均耗时", 
    "nr_switches": "主动切换和被动切换的累计次数", 
    "nr_voluntary_switches": "主动切换次数", 
    "nr_involuntary_switches": "被动切换次数", 
    "se.load.weight": "负载均衡相关的权重", 
    "se.avg.load_sum": "", 
    "se.avg.util_sum": "", 
    "se.avg.load_avg": "", 
    "se.avg.util_avg": "", 
    "se.avg.last_update_time": "", 
    "policy": "进程属性，0 为非实时", 
    "prio": "动态优先级", 
    "clock-delta": "连续两次调用 cpu_clock() 之间的时间差，单位为 ns",  // https://stackoverflow.com/questions/15024601/what-is-clock-delta-in-proc-pid-sched
    "mm->numa_scan_seq": "", 
    "numa_pages_migrated": "", 
    "numa_preferred_nid": "", 
    "total_numa_faults": "", 
}

let limitsHdr = `进程的资源限制`
let statusHdr = `进程的状态信息`
let syscallHdr = `当前进程正在执行的系统调用
第一个值是系统调用号，后面跟着 6 个系统调用的参数值（位于寄存器中），最后两个值依次是堆栈指针和指令计数器的值。
如果当前进程虽然阻塞，但阻塞函数并不是系统调用，则系统调用号的值为 -1，后面只有堆栈指针和指令计数器的值。
如果进程没有阻塞，则这个文件只有一个 running 的字符串。`
let environHdr = `进程的环境变量`
let fdHdr = `进程的文件描述符`
let mapFilesHdr = `映射的文件，即进程将使用到的二进制和动态库等，映射到该进程的内存虚拟地址的地址段信息`
let autoGroupHdr = `task group`
let cgroupHdr = `进程的 cgroup 设置`
let uidMapHdr = `uid_map 文件公开了用户 ID 从进程 pid 的用户命名空间到打开 uid_map 的进程的用户命名空间的映射（但是请参阅下面的这一点的限定条件）。 

换句话说，根据读取进程的用户命名空间的用户ID映射，在从特定uid_map文件读取时，位于不同用户命名空间中的进程可能会看到不同的值。

uid_map文件中的每一行指定两个用户命名空间之间的一系列连续用户ID的一对一映射。（首次创建用户命名空间时，此文件为空。）
每行中由被空格分隔的三个数字组成。 前两个数字指定两个用户命名空间中每个用户名称的起始用户 ID。 第三个数字指定映射范围的长度。 
字段解释如下：
    1. 进程pid的用户命名空间中用户标识范围的开始。
    2. 由字段一映射指定的用户ID的用户ID范围的开始。 如何解释字段2取决于打开uid_map的进程和进程pid是否在同一个用户命名空间中，如下所示：
        a. 如果两个进程位于不同的用户名称空间中：字段2是打开uid_map的进程的用户名称空间中一系列用户标识的开头。
        b. 如果两个进程在同一个用户命名空间中：字段2是进程pid的父用户命名空间中用户ID范围的开始。 这种情况启用uid_map的开启者（这里的常见情况是打开/ proc / self / uid_map）来查看用户ID到创建此用户命名空间的进程的用户命名空间的映射。
    3. 在两个用户名称空间之间映射的用户标识范围的长度。`
let gidMapHdr = `同 uid map，只是将 uid 换位 gid`
let projidMapHdr = `no tips`
let mountInfoHdr = `该文件包含挂载点的信息，其每一行内容如下：

36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
(1)(2)(3)   (4)   (5)      (6)      (7)   (8) (9)   (10)         (11)

1. mount ID: moun t的唯一标识符（可以在umount（2）之后重用）。
2. 父 ID: 上层挂载点的 ID（或跟挂载点本身）。
3. major：minor：文件系统上文件的st_dev值（参见stat（2））。
4. root：文件系统中 mount 的根目录。
5. 挂载点：相对于进程根的挂载点。
6. mount选项：per-mount选项。
7. 可选字段：“tag [：value]” 形式的零个或多个字段。
8. 分隔符：标记可选字段的结尾。
9. filesystem type：文件系统的名称，格式为 “type [.subtype]”。
10. mount source：特定于文件系统的信息或 “none”。
11. super 选项：per-superblock。`

let mountsHdr = `当前进程的 mount 命名空间中的所有文件系统的列表。 记录格式同 fstab(5)`
let mountStatsHdr = `该文件导出有关进程的 mount 命名空间中的挂载点的信息（统计信息，配置信息）。 
文件中的行内容说明如下：

device /dev/sda7 mounted on /home with fstype ext3 [statistics]
(       1      )            ( 2 )             (3 ) (4)

1. 设备名称（如果没有相应的设备，则为“nodevice”）。
2. 文件系统中的挂载点。
3. 文件系统类型。
4. 可选的统计信息和配置信息。 目前（如在Linux 2.6.26中），只有NFS文件系统通过此字段导出信息。

此文件只能由进程所有者读取。

详细内容参考 namespaces(7)。`


let other = {
    oomAdj: oomAdj, 
    oomScore: oomScore,
    oomScoreAdj: oomScoreAdj,
    loginUid: loginUid,
    sessionID: sessionID, 
    setGroup: setGroup,
    cpuSet: cpuSet, 
    coredumpFilter: coredumpFilter,
    personality: personality,
    cwd: cwd, 
    schedHdr: schedHdr,
    sched: sched,
    limitsHdr: limitsHdr,  
    statusHdr: statusHdr, 
    syscallHdr: syscallHdr,
    environHdr: environHdr,  
    fdHdr: fdHdr, 
    mapFilesHdr: mapFilesHdr, 
    autoGroupHdr: autoGroupHdr,
    cgroupHdr: cgroupHdr,
    uidMapHdr: uidMapHdr,
    gidMapHdr: gidMapHdr, 
    projidMapHdr: projidMapHdr,
    mountInfoHdr: mountInfoHdr,
    mountsHdr: mountsHdr,
    mountStatsHdr: mountStatsHdr,
}

exports.hdr = hdr;
exports.state = state;
exports.detailsTabs = detailsTabs;
exports.limits = limits;
exports.stack = stack;
exports.smapsHdr = smapsHdr;
exports.numaMapsHdr = numaMapsHdr;
exports.other = other;