// http://man7.org/linux/man-pages/man7/user_namespaces.7.html

// The uid_map file exposes the mapping of user IDs from the user
// namespace of the process pid to the user namespace of the process
// that opened uid_map (but see a qualification to this point below).
// In other words, processes that are in different user namespaces will
// potentially see different values when reading from a particular
// uid_map file, depending on the user ID mappings for the user
// namespaces of the reading processes.

// Each line in the uid_map file specifies a 1-to-1 mapping of a range
// of contiguous user IDs between two user namespaces.  (When a user
// namespace is first created, this file is empty.)  The specification
// in each line takes the form of three numbers delimited by white
// space.  The first two numbers specify the starting user ID in each of
// the two user namespaces.  The third number specifies the length of
// the mapped range.  In detail, the fields are interpreted as follows:

// (1) The start of the range of user IDs in the user namespace of the process pid.
// (2) The start of the range of user IDs to which the user IDs specified by field one map.  How field two is interpreted depends on whether the process that opened uid_map and the process pid are in the same user namespace, as follows:
// 	a) If the two processes are in different user namespaces: field two is the start of a range of user IDs in the user namespace of the process that opened uid_map.
// 	b) If the two processes are in the same user namespace: field two is the start of the range of user IDs in the parent user namespace of the process pid.  This case enables the opener of uid_map (the common case here is opening /proc/self/uid_map) to see the mapping of user IDs into the user namespace of the process that created this user namespace.
// (3) The length of the range of user IDs that is mapped between the two user namespaces.

// uid_map 文件公开了用户 ID 从进程 pid 的用户命名空间到打开 uid_map 的进程的用户命名空间的映射（但是请参阅下面的这一点的限定条件）。 
// 换句话说，根据读取进程的用户命名空间的用户ID映射，在从特定uid_map文件读取时，位于不同用户命名空间中的进程可能会看到不同的值。

// uid_map文件中的每一行指定两个用户命名空间之间的一系列连续用户ID的一对一映射。（首次创建用户命名空间时，此文件为空。）
// 每行中由被空格分隔的三个数字组成。 前两个数字指定两个用户命名空间中每个用户名称的起始用户 ID。 第三个数字指定映射范围的长度。 
// 字段解释如下：
// （1）进程pid的用户命名空间中用户标识范围的开始。
// （2）由字段一映射指定的用户ID的用户ID范围的开始。 如何解释字段2取决于打开uid_map的进程和进程pid是否在同一个用户命名空间中，如下所示：
//     a）如果两个进程位于不同的用户名称空间中：字段2是打开uid_map的进程的用户名称空间中一系列用户标识的开头。
//     b）如果两个进程在同一个用户命名空间中：字段2是进程pid的父用户命名空间中用户ID范围的开始。 这种情况启用uid_map的开启者（这里的常见情况是打开/ proc / self / uid_map）来查看用户ID到创建此用户命名空间的进程的用户命名空间的映射。
// （3）在两个用户名称空间之间映射的用户标识范围的长度。

package pid

import (
	"io/ioutil"
	"path/filepath"
)

func UidMapRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "uid_map"))
	return string(buf)
}