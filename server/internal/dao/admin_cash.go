// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminCashDao is internal type for wrapping internal DAO implements.
type internalAdminCashDao = *internal.AdminCashDao

// adminCashDao is the data access object for table gc_admin_cash.
// You can define custom methods on it to extend its functionality as you wish.
type adminCashDao struct {
	internalAdminCashDao
}

var (
	// AdminCash is globally public accessible object for table gc_admin_cash operations.
	AdminCash = adminCashDao{
		internal.NewAdminCashDao(),
	}
)

// Fill with you ideas below.
