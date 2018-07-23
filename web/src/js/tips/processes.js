
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
    // limit: "进程的资源限制",
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
    MapMax: "扫描期间遇到的最大mapcount（映射单个页面的进程数）。 这可以用作在给定存储器范围内发生的共享程度的指示符", 
    SwapCache: "在交换设备上具有关联条目的页数", 
    Active: "活动列表中的页数。仅当与此范围内的页数不同时，才会显示此字段。 这意味着存储器范围中存在一些非活动页面，这些页面很快就会被交换器从内存中删除", 
    WriteBack: "当前正在回写到磁盘的页数", 
    KernelPageSizeKB: "内核页大小", 
}

let oomAdj = `在内存不足（OOM）的情况下，系统根据各个进程的该文件的值来权衡要杀死的进程。 
内核使用这个进行进程的 oom_score 值的位移操作：有效值在 -16 到 +15 的范围内，再加上特殊值 -17，它会为此进程完全禁用OOM查杀。 
正分会增加这个过程被 OOM 杀手杀死的可能性; 负分数降低了可能性。
该文件的默认值为0; 新进程继承其父进程的oom_adj设置。 进程必须具有特权（CAP_SYS_RESOURCE）才能更新此文件。
从Linux 2.6.36开始，不推荐使用此文件，而使用/ proc / [pid] / oom_score_adj。
`

let oomScore = `此文件显示内核为此进程提供的当前分数，以便为OOM杀手选择进程。 得分越高意味着OOM杀手更有可能选择该过程。 此分数的取决于进程运行过程使用的内存量，决策因素如下：
1. 该进程是否使用fork（2）（+）创建了很多子进程;
2. 进程是否已运行很长时间，或者是否占用了大量CPU时间（ - ）;
3. 该进程是否具有低的 nice 值（即> 0）（+）;
4. 流程是否具有特权（ - ）; 和
5. 该进程是否进行直接硬件访问（ - ）。
oom_score还反映了进程的oom_score_adj或oom_adj设置指定的调整。`

let oomScoreAdj = `
This file can be used to adjust the badness heuristic used to select which process gets killed in out-of-memory conditions.

The badness heuristic assigns a value to each candidate task ranging from 0 (never kill) to 1000 (always kill) to
determine which process is targeted.  The units are roughly a proportion along that range of allowed  memory  the
process  may allocate from, based on an estimation of its current memory and swap use.  For example, if a task is
using all allowed memory, its badness score will be 1000.  If it is using half of its allowed memory,  its  score
will be 500.

There  is an additional factor included in the badness score: root processes are given 3% extra memory over other
tasks.

The amount of "allowed" memory depends on the context in which the OOM-killer was called.  If it is  due  to  the
memory  assigned  to  the allocating task's cpuset being exhausted, the allowed memory represents the set of mems
assigned to that cpuset (see cpuset(7)).  If it is due to a mempolicy's node(s) being exhausted, the allowed mem‐
ory  represents  the  set  of mempolicy nodes.  If it is due to a memory limit (or swap limit) being reached, the
allowed memory is that configured limit.  Finally, if it is due to the entire system being  out  of  memory,  the
allowed memory represents all allocatable resources.

The  value  of  oom_score_adj  is  added  to the badness score before it is used to determine which task to kill.
Acceptable values range from -1000 (OOM_SCORE_ADJ_MIN) to +1000 (OOM_SCORE_ADJ_MAX).  This allows user  space  to
control  the preference for OOM-killing, ranging from always preferring a certain task or completely disabling it
from OOM-killing.  The lowest possible value, -1000, is equivalent to disabling  OOM-killing  entirely  for  that
task, since it will always report a badness score of 0.

Consequently, it is very simple for user space to define the amount of memory to consider for each task.  Setting
a oom_score_adj value of +500, for example, is roughly equivalent to allowing the remainder of tasks sharing  the
same system, cpuset, mempolicy, or memory controller resources to use at least 50% more memory.  A value of -500,
on the other hand, would be roughly equivalent to discounting 50% of the task's allowed memory from being consid‐
ered as scoring against the task.

For  backward  compatibility  with  previous  kernels,  /proc/[pid]/oom_adj can still be used to tune the badness
score.  Its value is scaled linearly with oom_score_adj.

Writing to /proc/[pid]/oom_score_adj or /proc/[pid]/oom_adj will change the other with its scaled value.`

let other = {
    oomAdj: oomAdj, 
    oomScore: oomScore,
    oomScoreAdj: oomScoreAdj,
}

exports.hdr = hdr;
exports.state = state;
exports.detailsTabs = detailsTabs;
exports.limits = limits;
exports.stack = stack;
exports.smapsHdr = smapsHdr;
exports.numaMapsHdr = numaMapsHdr;
exports.other = other;