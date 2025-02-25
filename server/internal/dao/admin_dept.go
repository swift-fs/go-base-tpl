// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminDeptDao is internal type for wrapping internal DAO implements.
type internalAdminDeptDao = *internal.AdminDeptDao

// adminDeptDao is the data access object for table gc_admin_dept.
// You can define custom methods on it to extend its functionality as you wish.
type adminDeptDao struct {
	internalAdminDeptDao
}

var (
	// AdminDept is globally public accessible object for table gc_admin_dept operations.
	AdminDept = adminDeptDao{
		internal.NewAdminDeptDao(),
	}
)

// Fill with you ideas below.
