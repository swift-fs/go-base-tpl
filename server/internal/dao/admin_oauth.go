// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminOauthDao is internal type for wrapping internal DAO implements.
type internalAdminOauthDao = *internal.AdminOauthDao

// adminOauthDao is the data access object for table gc_admin_oauth.
// You can define custom methods on it to extend its functionality as you wish.
type adminOauthDao struct {
	internalAdminOauthDao
}

var (
	// AdminOauth is globally public accessible object for table gc_admin_oauth operations.
	AdminOauth = adminOauthDao{
		internal.NewAdminOauthDao(),
	}
)

// Fill with you ideas below.
