
let cpu = {
    Avg: "平均 CPU 使用率",
}

let Memory = `kernel 可支配的物理内存
系统从加电开始到引导完成
firmware/BIOS 要保留一些内存，kernel 本身要占用一些内存
最后剩下的才是 kernel 可支配的内存`

let Swap = `交换分区
交换分区一般是通过磁盘挂载出来的，速度非常慢`

let Buffers = `块设备(block device)所占用的缓存页
包括：直接读写块设备、以及文件系统元数据(metadata)，比如 SuperBlock 所使用的缓存页
它与 Cached 的区别在于， Cached 表示普通文件所占用的缓存页
Buffers 所占的内存同时也在 LRU list 中，被统计在 Active(file) 或 Inactive(file)`

let Cached = `所有的 file-backed pages, 表示普通文件数据所占用的缓存页
Cached 是 Mapped 的超集，不仅包括 mapped，也包括 unmapped 的页面
当一个文件不再与进程关联之后，原来在page cache中的页面并不会立即回收，仍然被计入Cached，还留在LRU中，但是 Mapped 统计值会减小。
ummaped = (Cached – Mapped)

Cached包含tmpfs中的文件，POSIX/SysV shared memory，以及shared anonymous mmap
  注：POSIX/SysV shared memory和shared anonymous mmap在内核中都是基于tmpfs实现的
  
Cached 和 SwapCached 两个统计值互不重叠`

let Active = `包含 active anon 和 active file`
let Inactive = `包含 inactive anon 和 inactive file`
let ActiveAnon = ` anonymous pages（匿名页）

用户进程的内存页分为两种：
  1. 与文件关联的内存页(比如程序文件,数据文件对应的内存页)
  2. 与内存无关的内存页（比如进程的堆栈，用malloc申请的内存）

前者称为 file pages 或 mapped pages,后者称为匿名页。`
let InactiveAnon = `LRU lists 中的 LRU_INACTIVE_ANON  内存

（LRU 是 Kernel 的页面回收算法(Page Frame Reclaiming) 使用的数据结构
Page cache 和所有用户进程的内存（kernel stack和huge pages除外）都在 LRU lists 上）`

let ActiveFile = `LRU lists 中的 LRU_ACTIVE_FILE  内存

（LRU 是 Kernel 的页面回收算法(Page Frame Reclaiming) 使用的数据结构
Page cache 和所有用户进程的内存（kernel stack和huge pages除外）都在 LRU lists 上）`

let InactiveFile = `LRU lists 中的 LRU_INACTIVE_FILE  内存

（LRU 是 Kernel 的页面回收算法(Page Frame Reclaiming) 使用的数据结构
Page cache 和所有用户进程的内存（kernel stack和huge pages除外）都在 LRU lists 上）`

let Unevictable = `LRU list 上是不能 pageout/swapout 的内存页

LRU lists 中的 LRU_UNEVICTABLE  内存

包括 VM_LOCKED 的内存页、SHM_LOCK 的共享内存页（又被统计在 Mlocked 中）、和 ramfs

（LRU 是 Kernel 的页面回收算法(Page Frame Reclaiming) 使用的数据结构
Page cache 和所有用户进程的内存（kernel stack和huge pages除外）都在 LRU lists 上）`

let Mlocked = `被mlock()系统调用锁定的内存大小

被锁定的内存因为不能 pageout/swapout，会从 Active/Inactive LRU list 移到 Unevictable LRU list 上
也就是说，当 Mlocked 增加时，Unevictable 也同步增加，而 Active 或 Inactive 同时减小
当 Mlocked 减小的时候， Unevictable 也同步减小，而 Active 或 Inactive ”同时增加。

 Mlocked 与以下统计项重叠：LRU Unevictable，AnonPages，Shmem，Mapped 等。`
let Dirty = `需要写入磁盘的内存去的大小`
let Writeback = `正在被写回的内存区的大小`
let AnonPages = `anonymous pages（匿名页）
用户进程的内存页分为两种：file-backed pages（与文件对应的内存页），和 anonymous pages（匿名页）

所有 page cache 里的页面(Cached) 都是 file-backed pages，不是 Anonymous Pages。Cached 与 AnoPages 之间没有重叠
  注：shared memory 不属于 AnonPages，而是属于 Cached，因为 shared memory 基于 tmpfs，所以被视为 file-backed、在 page cache 里

mmap private anonymous pages 属于 AnonPages(Anonymous Pages)
而 mmap shared anonymous pages 属于 Cached(file-backed pages)
因为shared anonymous mmap也是基于tmpfs的

Anonymous Pages是与用户进程共存的
一旦进程退出，则Anonymous pages也释放，不像page cache即使文件与进程不关联了还可以缓存

AnonPages统计值中包含了Transparent HugePages (THP)对应的 AnonHugePages`

let Mapped = `设备和文件等映射的大小

包含 page cache(Cached) 中所有的 mapped 页面,  Mapped 是 Cached 的子集
Page cache 中(Cached)包含了文件的缓存页，其中有些文件当前已不在使用，
page cache 仍然可能保留着它们的缓存页面；
而另一些文件正被用户进程关联，比如 shared libraries、可执行程序的文件、mmap 的文件等，这些文件的缓存页就称为 mapped`

let Shmem = `包括 shared memory 和 tmpfs

shared memory 包括：
  SysV shared memory [shmget etc.]
  POSIX shared memory [shm_open etc.]
  shared anonymous mmap [ mmap(…MAP_ANONYMOUS|MAP_SHARED…)]
（因为shared memory在内核中都是基于tmpfs实现的）

这些 shared memory 被视为基于 tmpfs 文件系统的内存页，
由于基于文件系统，所以不算匿名页，
所以不被计入 /proc/meminfo 中的 AnonPages，而是被统计进了：
  1. Cached (i.e. page cache)
  2. Mapped (当shmem被attached时候)

在LRU lists里，被放在：
  1. Inactive(anon) 或 Active(anon)
    注：虽然它们在LRU中被放进了anon list，但是不会被计入 AnonPages。这是shared memory & tmpfs比较拧巴的一个地方，需要特别注意。
  2. unevictable （如果被locked的话）

注意：
当 shmget/shm_open/mmap 创建共享内存时
物理内存尚未分配，要直到真正访问时才分配
这里统计的是已经分配的大小，而不是创建时申请的大小
`

let Slab = `内核数据结构 slab 的大小
Linux 为了提高性能，会对重复使用的对象进行缓冲`

let SReclaimable = `slab 中可回收的部分
调用 kmem_getpages() 时加上 SLAB_RECLAIM_ACCOUNT 标记，
表明是可回收的，计入 SReclaimable，否则计入 SUnreclaim`

let SUnreclaim = `slab 中不可回收的部分`

let KernelStack = `为每个用户线程分配的内核栈所使用的内存
内核栈是常驻内存的，既不包括在 LRU lists 里，也不包括在进程的 RSS/PSS 内存里，所以将其归类到 kernel 消耗的内存中。
内核栈虽然属于线程，但用户态的代码不能访问，
只有通过系统调用 (syscall)、自陷 (trap) 或异常(exception)进入内核态的时候才会用到，也就是说内核栈是给 kernel code 使用的。
在x86系统上Linux的内核栈大小是固定的 8K 或 16K`

let PageTables = `管理内存页的页面的大小
Page Table 用于将内存的虚拟地址翻译成物理地址，随着内存地址分配得越来越多，Page Table 会增大`

let NfsUnstable = `NFS 页面发送到服务器，但尚未写入到稳定存储的内存`

let Bounce = `有些老设备只能访问低端内存，比如 16M 以下的内存
当应用程序发出一个 I/O 请求，DMA 的目的地址却是高端内存时（比如在 16M 以上）
内核将在低端内存中分配一个临时buffer作为跳转，把位于高端内存的缓存数据复制到此处
这种额外的数据拷贝被称为 “bounce buffering”，会降低 I/O 性能
大量分配的 bounce buffers 也会占用额外的内存。`

let WritebackTmp = `FUSE 用于临时写回缓冲区的内存`

let CommitLimit = `可在当前系统上分配的内存总量

只有在启用了 strict overcommit accounting(mode 2 in /proc/sys/vm/overcommit_ratio) 时才会遵守这个限制

这个值是根据过度使用率(vm.overcommit_ratio) 计算出来的
计算公式：
  CommitLimit = (overcommit_ratio * Physical RAM) + Swap

例如：
  在 1GB物理内存和 7GB 交换分区，overcommit_ratio 为 30 的系统上，
  公式产生的 CommitLimit 为 7.3GB`

let CommittedAS = `当前在系统上分配的内存量
提交的内存是进程分配的所有内存的总和，包括未使用部分

一个分配 1GB 内存的进程（使用 malloc（3）或类似接口），
但只 touch 300MB 的内存的话，将显示为仅使用300MB内存，
即使它具有分配给整个1GB的地址空间。
这1GB是由 VM “提交” 的内存，可以由分配应用程序随时使用

在系统上启用严格过度使用（模式2 / proc / sys / vm / overcommit_memory）时，
将不允许超过CommitLimit（上面详述）的分配。 
如果需要保证一旦成功分配内存后进程不会因内存不足而失败，这非常有用。`

let Vmalloc = `vmalloc 内存去的使用状况`

let HardwareCorrupted = `由于内存的硬件故障而删除掉的内存页的总大小
当系统检测到内存的硬件故障时，会把有问题的页面删除掉，不再使用`

let AnonHugePages = `Transparent HugePages (THP),
被包含在 AnonPages 之中,与进程的 RSS/PSS 有重叠，
如果用户进程用到了 THP，进程的 RSS/PSS 也会相应增加`

let HugePages = `Hugepages 状态
Hugepages 是被独立统计的，与其它统计项不重叠
既不计入进程的 RSS/PSS 中，又不计入 LRU Active/Inactive，也不会计入 cache/buffer
如果进程使用了Hugepages，它的RSS/PSS不会增加
Hugepages 不同于 Transparent HugePages (THP)
THP的统计值是 /proc/meminfo 中的 AnonHugePages`

let DirectMap4k = `映射为 4kB 的内存数量

DirectMap 所统计的不是关于内存的使用，而是一个反映 TLB 效率的指标
TLB(Translation Lookaside Buffer) 是位于CPU上的缓存，用于将内存的虚拟地址翻译成物理地址，
由于TLB的大小有限，不能缓存的地址就需要访问内存里的 page table 来进行翻译，速度慢很多
为了尽可能地将地址放进TLB缓存，新的 CPU 硬件支持比 4k 更大的页面从而达到减少地址数量的目的
比如2MB，4MB，甚至1GB的内存页，视不同的硬件而定
`


let DirectMap2M = `映射为 2MB 的内存数量

DirectMap 所统计的不是关于内存的使用，而是一个反映 TLB 效率的指标
TLB(Translation Lookaside Buffer) 是位于CPU上的缓存，用于将内存的虚拟地址翻译成物理地址，
由于TLB的大小有限，不能缓存的地址就需要访问内存里的 page table 来进行翻译，速度慢很多
为了尽可能地将地址放进TLB缓存，新的 CPU 硬件支持比 4k 更大的页面从而达到减少地址数量的目的
比如2MB，4MB，甚至1GB的内存页，视不同的硬件而定`


let DirectMap1G = `映射为 1GB 的内存数量

DirectMap 所统计的不是关于内存的使用，而是一个反映 TLB 效率的指标
TLB(Translation Lookaside Buffer) 是位于CPU上的缓存，用于将内存的虚拟地址翻译成物理地址，
由于TLB的大小有限，不能缓存的地址就需要访问内存里的 page table 来进行翻译，速度慢很多
为了尽可能地将地址放进TLB缓存，新的 CPU 硬件支持比 4k 更大的页面从而达到减少地址数量的目的
比如2MB，4MB，甚至1GB的内存页，视不同的硬件而定`


let mem = {
    'Memory': Memory,

    'Swap': Swap, 
    'Buffers': Buffers, 
    'Cached': Cached, 

    'Active': Active, 
    'Inactive': Inactive, 
    'ActiveAnon': ActiveAnon, 
    'InactiveAnon': InactiveAnon, 
    'ActiveFile': ActiveFile, 
    'InactiveFile': InactiveFile, 
    'Unevictable': Unevictable, 

    'Mlocked': Mlocked, 
    'Dirty': Dirty, 
    'Writeback': Writeback, 
    'AnonPages': AnonPages, 
    'Mapped': Mapped, 
    'Shmem': Shmem, 

    'Slab': Slab, 
    'SReclaimable': SReclaimable, 
    'SUnreclaim': SUnreclaim, 

    'KernelStack': KernelStack, 
    'PageTables': PageTables, 

    'NfsUnstable': NfsUnstable, 
    'Bounce': Bounce, 
    'WritebackTmp': WritebackTmp, 
    'CommitLimit': CommitLimit, 
    'CommittedAS': CommittedAS, 

    'Vmalloc': Vmalloc, 

    'HardwareCorrupted': HardwareCorrupted, 

    'AnonHugePages': AnonHugePages, 

    'HugePages': HugePages, 

    'DirectMap4k': DirectMap4k, 
    'DirectMap2M': DirectMap2M, 
    'DirectMap1G': DirectMap1G,
}

let net = {
    "Sum(rx)": "所有网络接口的接收速率之和",
    "Sum(tx)": "所有网络接口的上传速率之和",
}

let disk = {
    "Sum(r)": "所有磁盘的读取速率之和",
    "Sum(w)": "所有磁盘的写入速率之和",
}

let legend = {
    cpu: cpu,
    mem: mem,
    net: net,
    disk: disk,
}

exports.legend = legend