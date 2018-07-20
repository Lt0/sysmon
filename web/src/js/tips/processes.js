
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
    limit: "进程的资源限制",
    stack: "所有线程的内部调用状态",
    smaps: "整个进程的内存映射信息",
    numaMaps: "整个进程的内存在 NUMA 节点之间的映射关系",
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
let smapsHdrVmFlags = `
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
    VmFlags: "VmFlags field represents the kernel flags associated with the particular virtual memory area in two letter encoded manner. The codes are the following: " + smapsHdrVmFlags, 
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

exports.hdr = hdr;
exports.state = state;
exports.detailsTabs = detailsTabs;
exports.limits = limits;
exports.stack = stack;
exports.smapsHdr = smapsHdr;
exports.numaMapsHdr = numaMapsHdr;