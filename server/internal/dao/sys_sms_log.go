// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysSmsLogDao is internal type for wrapping internal DAO implements.
type internalSysSmsLogDao = *internal.SysSmsLogDao

// sysSmsLogDao is the data access object for table gc_sys_sms_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysSmsLogDao struct {
	internalSysSmsLogDao
}

var (
	// SysSmsLog is globally public accessible object for table gc_sys_sms_log operations.
	SysSmsLog = sysSmsLogDao{
		internal.NewSysSmsLogDao(),
	}
)

// Fill with you ideas below.
