// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleDao is the data access object for table gc_admin_role.
type AdminRoleDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminRoleColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleColumns defines and stores column names for table gc_admin_role.
type AdminRoleColumns struct {
	Id         string // 角色ID
	Name       string // 角色名称
	Key        string // 角色权限字符串
	DataScope  string // 数据范围
	CustomDept string // 自定义部门权限
	Pid        string // 上级角色ID
	Level      string // 关系树等级
	Tree       string // 关系树
	Remark     string // 备注
	Sort       string // 排序
	Status     string // 角色状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// adminRoleColumns holds the columns for table gc_admin_role.
var adminRoleColumns = AdminRoleColumns{
	Id:         "id",
	Name:       "name",
	Key:        "key",
	DataScope:  "data_scope",
	CustomDept: "custom_dept",
	Pid:        "pid",
	Level:      "level",
	Tree:       "tree",
	Remark:     "remark",
	Sort:       "sort",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewAdminRoleDao creates and returns a new DAO object for table data access.
func NewAdminRoleDao() *AdminRoleDao {
	return &AdminRoleDao{
		group:   "default",
		table:   "gc_admin_role",
		columns: adminRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleDao) Columns() AdminRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
