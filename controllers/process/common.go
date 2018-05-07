package process

type Process struct {
	Cmd			string	// 进程命令
	User		string
	
	CPU			float64	// 进程使用的 CPU 百分比
	CPUTime		string	// 进程使用的 CPU 时间
	
	MEM			float64 // 进程使用的内存百分比
	SZ			uint64	// 进程核心镜像的物理页面大小。包括 text，data 和 stack 的空间，不包括设备映射
	RSS			uint64	// 常驻内存大小，任务已使用的为交换物理内存，以 KB 为单位
	DRS			uint64	// 进程的数据集大小，专门用于非可执行代码的物理内存量
	TRS			uint64	// 进程的文本段占用的内存
	VSZ			uint64	// 进程的虚拟内存大小，以 KB 为单位。

	Pid			int		// 进程号
	Nlwp		int		// 进程内的线程数

	State		string	// 进程状态：D 不可中断; R 运行;	S 中断;	T 停止;	Z 僵死
	Nice		int		// 进程优先级
}