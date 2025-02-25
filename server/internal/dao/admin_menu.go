// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminMenuDao is internal type for wrapping internal DAO implements.
type internalAdminMenuDao = *internal.AdminMenuDao

// adminMenuDao is the data access object for table gc_admin_menu.
// You can define custom methods on it to extend its functionality as you wish.
type adminMenuDao struct {
	internalAdminMenuDao
}

var (
	// AdminMenu is globally public accessible object for table gc_admin_menu operations.
	AdminMenu = adminMenuDao{
		internal.NewAdminMenuDao(),
	}
)

// Fill with you ideas below.
