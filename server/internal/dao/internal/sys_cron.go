// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronDao is the data access object for table gc_sys_cron.
type SysCronDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysCronColumns // columns contains all the column names of Table for convenient usage.
}

// SysCronColumns defines and stores column names for table gc_sys_cron.
type SysCronColumns struct {
	Id        string // 任务ID
	GroupId   string // 分组ID
	Title     string // 任务标题
	Name      string // 任务方法
	Params    string // 函数参数
	Pattern   string // 表达式
	Policy    string // 策略
	Count     string // 执行次数
	Sort      string // 排序
	Remark    string // 备注
	Status    string // 任务状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysCronColumns holds the columns for table gc_sys_cron.
var sysCronColumns = SysCronColumns{
	Id:        "id",
	GroupId:   "group_id",
	Title:     "title",
	Name:      "name",
	Params:    "params",
	Pattern:   "pattern",
	Policy:    "policy",
	Count:     "count",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysCronDao creates and returns a new DAO object for table data access.
func NewSysCronDao() *SysCronDao {
	return &SysCronDao{
		group:   "default",
		table:   "gc_sys_cron",
		columns: sysCronColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysCronDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysCronDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysCronDao) Columns() SysCronColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysCronDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysCronDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysCronDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
