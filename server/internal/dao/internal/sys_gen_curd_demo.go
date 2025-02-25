// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCurdDemoDao is the data access object for table gc_sys_gen_curd_demo.
type SysGenCurdDemoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysGenCurdDemoColumns // columns contains all the column names of Table for convenient usage.
}

// SysGenCurdDemoColumns defines and stores column names for table gc_sys_gen_curd_demo.
type SysGenCurdDemoColumns struct {
	Id          string // ID
	CategoryId  string // 分类ID
	Title       string // 标题
	Description string // 描述
	Content     string // 内容
	Image       string // 单图
	Attachfile  string // 附件
	CityId      string // 所在城市
	Switch      string // 显示开关
	Sort        string // 排序
	Status      string // 状态
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// sysGenCurdDemoColumns holds the columns for table gc_sys_gen_curd_demo.
var sysGenCurdDemoColumns = SysGenCurdDemoColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Title:       "title",
	Description: "description",
	Content:     "content",
	Image:       "image",
	Attachfile:  "attachfile",
	CityId:      "city_id",
	Switch:      "switch",
	Sort:        "sort",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewSysGenCurdDemoDao creates and returns a new DAO object for table data access.
func NewSysGenCurdDemoDao() *SysGenCurdDemoDao {
	return &SysGenCurdDemoDao{
		group:   "default",
		table:   "gc_sys_gen_curd_demo",
		columns: sysGenCurdDemoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysGenCurdDemoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysGenCurdDemoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysGenCurdDemoDao) Columns() SysGenCurdDemoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysGenCurdDemoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysGenCurdDemoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysGenCurdDemoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
