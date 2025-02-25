// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCodesDao is the data access object for table gc_sys_gen_codes.
type SysGenCodesDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysGenCodesColumns // columns contains all the column names of Table for convenient usage.
}

// SysGenCodesColumns defines and stores column names for table gc_sys_gen_codes.
type SysGenCodesColumns struct {
	Id            string // 生成ID
	GenType       string // 生成类型
	GenTemplate   string // 生成模板
	VarName       string // 实体命名
	Options       string // 配置选项
	DbName        string // 数据库名称
	TableName     string // 主表名称
	TableComment  string // 主表注释
	DaoName       string // 主表dao模型
	MasterColumns string // 主表字段
	AddonName     string // 插件名称
	Status        string // 生成状态
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// sysGenCodesColumns holds the columns for table gc_sys_gen_codes.
var sysGenCodesColumns = SysGenCodesColumns{
	Id:            "id",
	GenType:       "gen_type",
	GenTemplate:   "gen_template",
	VarName:       "var_name",
	Options:       "options",
	DbName:        "db_name",
	TableName:     "table_name",
	TableComment:  "table_comment",
	DaoName:       "dao_name",
	MasterColumns: "master_columns",
	AddonName:     "addon_name",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysGenCodesDao creates and returns a new DAO object for table data access.
func NewSysGenCodesDao() *SysGenCodesDao {
	return &SysGenCodesDao{
		group:   "default",
		table:   "gc_sys_gen_codes",
		columns: sysGenCodesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysGenCodesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysGenCodesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysGenCodesDao) Columns() SysGenCodesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysGenCodesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysGenCodesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysGenCodesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
