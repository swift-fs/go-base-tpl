// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysEmsLogDao is the data access object for table gc_sys_ems_log.
type SysEmsLogDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysEmsLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysEmsLogColumns defines and stores column names for table gc_sys_ems_log.
type SysEmsLogColumns struct {
	Id        string // 主键
	Event     string // 事件
	Email     string // 邮箱地址，多个用;隔开
	Code      string // 验证码
	Times     string // 验证次数
	Content   string // 邮件内容
	Ip        string // ip地址
	Status    string // 状态(1未验证,2已验证)
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysEmsLogColumns holds the columns for table gc_sys_ems_log.
var sysEmsLogColumns = SysEmsLogColumns{
	Id:        "id",
	Event:     "event",
	Email:     "email",
	Code:      "code",
	Times:     "times",
	Content:   "content",
	Ip:        "ip",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysEmsLogDao creates and returns a new DAO object for table data access.
func NewSysEmsLogDao() *SysEmsLogDao {
	return &SysEmsLogDao{
		group:   "default",
		table:   "gc_sys_ems_log",
		columns: sysEmsLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysEmsLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysEmsLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysEmsLogDao) Columns() SysEmsLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysEmsLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysEmsLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysEmsLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
