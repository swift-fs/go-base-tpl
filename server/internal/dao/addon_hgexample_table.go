// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAddonHgexampleTableDao is internal type for wrapping internal DAO implements.
type internalAddonHgexampleTableDao = *internal.AddonHgexampleTableDao

// addonHgexampleTableDao is the data access object for table gc_addon_hgexample_table.
// You can define custom methods on it to extend its functionality as you wish.
type addonHgexampleTableDao struct {
	internalAddonHgexampleTableDao
}

var (
	// AddonHgexampleTable is globally public accessible object for table gc_addon_hgexample_table operations.
	AddonHgexampleTable = addonHgexampleTableDao{
		internal.NewAddonHgexampleTableDao(),
	}
)

// Fill with you ideas below.
