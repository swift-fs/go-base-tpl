// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminPostDao is the data access object for table gc_admin_post.
type AdminPostDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminPostColumns // columns contains all the column names of Table for convenient usage.
}

// AdminPostColumns defines and stores column names for table gc_admin_post.
type AdminPostColumns struct {
	Id        string // 岗位ID
	Code      string // 岗位编码
	Name      string // 岗位名称
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// adminPostColumns holds the columns for table gc_admin_post.
var adminPostColumns = AdminPostColumns{
	Id:        "id",
	Code:      "code",
	Name:      "name",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminPostDao creates and returns a new DAO object for table data access.
func NewAdminPostDao() *AdminPostDao {
	return &AdminPostDao{
		group:   "default",
		table:   "gc_admin_post",
		columns: adminPostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminPostDao) Columns() AdminPostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
