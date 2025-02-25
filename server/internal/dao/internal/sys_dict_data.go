// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDataDao is the data access object for table gc_sys_dict_data.
type SysDictDataDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysDictDataColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictDataColumns defines and stores column names for table gc_sys_dict_data.
type SysDictDataColumns struct {
	Id        string // 字典数据ID
	Label     string // 字典标签
	Value     string // 字典键值
	ValueType string // 键值数据类型：string,int,uint,bool,datetime,date
	Type      string // 字典类型
	ListClass string // 表格回显样式
	IsDefault string // 是否为系统默认
	Sort      string // 字典排序
	Remark    string // 备注
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysDictDataColumns holds the columns for table gc_sys_dict_data.
var sysDictDataColumns = SysDictDataColumns{
	Id:        "id",
	Label:     "label",
	Value:     "value",
	ValueType: "value_type",
	Type:      "type",
	ListClass: "list_class",
	IsDefault: "is_default",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysDictDataDao creates and returns a new DAO object for table data access.
func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		group:   "default",
		table:   "gc_sys_dict_data",
		columns: sysDictDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDictDataDao) Columns() SysDictDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDictDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
