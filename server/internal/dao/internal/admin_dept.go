// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDeptDao is the data access object for table gc_admin_dept.
type AdminDeptDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminDeptColumns // columns contains all the column names of Table for convenient usage.
}

// AdminDeptColumns defines and stores column names for table gc_admin_dept.
type AdminDeptColumns struct {
	Id        string // 部门ID
	Pid       string // 父部门ID
	Name      string // 部门名称
	Code      string // 部门编码
	Type      string // 部门类型
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	Level     string // 关系树等级
	Tree      string // 关系树
	Sort      string // 排序
	Status    string // 部门状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// adminDeptColumns holds the columns for table gc_admin_dept.
var adminDeptColumns = AdminDeptColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Code:      "code",
	Type:      "type",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Level:     "level",
	Tree:      "tree",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminDeptDao creates and returns a new DAO object for table data access.
func NewAdminDeptDao() *AdminDeptDao {
	return &AdminDeptDao{
		group:   "default",
		table:   "gc_admin_dept",
		columns: adminDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminDeptDao) Columns() AdminDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
