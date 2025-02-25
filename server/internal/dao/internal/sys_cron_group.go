// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronGroupDao is the data access object for table gc_sys_cron_group.
type SysCronGroupDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysCronGroupColumns // columns contains all the column names of Table for convenient usage.
}

// SysCronGroupColumns defines and stores column names for table gc_sys_cron_group.
type SysCronGroupColumns struct {
	Id        string // 任务分组ID
	Pid       string // 父类任务分组ID
	Name      string // 分组名称
	IsDefault string // 是否默认
	Sort      string // 排序
	Remark    string // 备注
	Status    string // 分组状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysCronGroupColumns holds the columns for table gc_sys_cron_group.
var sysCronGroupColumns = SysCronGroupColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	IsDefault: "is_default",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysCronGroupDao creates and returns a new DAO object for table data access.
func NewSysCronGroupDao() *SysCronGroupDao {
	return &SysCronGroupDao{
		group:   "default",
		table:   "gc_sys_cron_group",
		columns: sysCronGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysCronGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysCronGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysCronGroupDao) Columns() SysCronGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysCronGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysCronGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysCronGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
