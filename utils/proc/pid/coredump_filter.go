// Controlling which mappings are written to the core dump
// Since kernel 2.6.23, the Linux-specific /proc/[pid]/coredump_filter
// file can be used to control which memory segments are written to the
// core dump file in the event that a core dump is performed for the
// process with the corresponding process ID.
// 
// The value in the file is a bit mask of memory mapping types (see
// mmap(2)).  If a bit is set in the mask, then memory mappings of the
// corresponding type are dumped; otherwise they are not dumped.  The
// bits in this file have the following meanings:
// 
//     bit 0  Dump anonymous private mappings.
//     bit 1  Dump anonymous shared mappings.
//     bit 2  Dump file-backed private mappings.
//     bit 3  Dump file-backed shared mappings.
//     bit 4 (since Linux 2.6.24)
//            Dump ELF headers.
//     bit 5 (since Linux 2.6.28)
//            Dump private huge pages.
//     bit 6 (since Linux 2.6.28)
//            Dump shared huge pages.
//     bit 7 (since Linux 4.4)
//            Dump private DAX pages.
//     bit 8 (since Linux 4.4)
//            Dump shared DAX pages.
// 
// By default, the following bits are set: 0, 1, 4 (if the
// CONFIG_CORE_DUMP_DEFAULT_ELF_HEADERS kernel configuration option is
// enabled), and 5.  This default can be modified at boot time using the
// coredump_filter boot option.
// 
// The value of this file is displayed in hexadecimal.  (The default
// value is thus displayed as 33.)
// 
// Memory-mapped I/O pages such as frame buffer are never dumped, and
// virtual DSO pages are always dumped, regardless of the
// coredump_filter value.
// 
// A child process created via fork(2) inherits its parent's
// coredump_filter value; the coredump_filter value is preserved across
// an execve(2).
// 
// It can be useful to set coredump_filter in the parent shell before
// running a program, for example:
// 
//     $ echo 0x7 > /proc/self/coredump_filter
//     $ ./some_program
// 
// This file is provided only if the kernel was built with the CON‐
// FIG_ELF_CORE configuration option.
package pid

import (
	"io/ioutil"
	"path/filepath"
)

func CoredumpFilterRawData(pid string) string {
	buf, _ := ioutil.ReadFile(filepath.Join(Ctx.Procfs, pid, "coredump_filter"))
	return string(buf)
}