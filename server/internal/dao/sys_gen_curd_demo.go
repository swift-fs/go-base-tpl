// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysGenCurdDemoDao is internal type for wrapping internal DAO implements.
type internalSysGenCurdDemoDao = *internal.SysGenCurdDemoDao

// sysGenCurdDemoDao is the data access object for table gc_sys_gen_curd_demo.
// You can define custom methods on it to extend its functionality as you wish.
type sysGenCurdDemoDao struct {
	internalSysGenCurdDemoDao
}

var (
	// SysGenCurdDemo is globally public accessible object for table gc_sys_gen_curd_demo operations.
	SysGenCurdDemo = sysGenCurdDemoDao{
		internal.NewSysGenCurdDemoDao(),
	}
)

// Fill with you ideas below.
