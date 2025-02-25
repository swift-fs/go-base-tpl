// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalPayLogDao is internal type for wrapping internal DAO implements.
type internalPayLogDao = *internal.PayLogDao

// payLogDao is the data access object for table gc_pay_log.
// You can define custom methods on it to extend its functionality as you wish.
type payLogDao struct {
	internalPayLogDao
}

var (
	// PayLog is globally public accessible object for table gc_pay_log operations.
	PayLog = payLogDao{
		internal.NewPayLogDao(),
	}
)

// Fill with you ideas below.
