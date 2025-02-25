// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminCashDao is the data access object for table gc_admin_cash.
type AdminCashDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminCashColumns // columns contains all the column names of Table for convenient usage.
}

// AdminCashColumns defines and stores column names for table gc_admin_cash.
type AdminCashColumns struct {
	Id        string // ID
	MemberId  string // 管理员ID
	Money     string // 提现金额
	Fee       string // 手续费
	LastMoney string // 最终到账金额
	Ip        string // 申请人IP
	Status    string // 状态码
	Msg       string // 处理结果
	HandleAt  string // 处理时间
	CreatedAt string // 申请时间
}

// adminCashColumns holds the columns for table gc_admin_cash.
var adminCashColumns = AdminCashColumns{
	Id:        "id",
	MemberId:  "member_id",
	Money:     "money",
	Fee:       "fee",
	LastMoney: "last_money",
	Ip:        "ip",
	Status:    "status",
	Msg:       "msg",
	HandleAt:  "handle_at",
	CreatedAt: "created_at",
}

// NewAdminCashDao creates and returns a new DAO object for table data access.
func NewAdminCashDao() *AdminCashDao {
	return &AdminCashDao{
		group:   "default",
		table:   "gc_admin_cash",
		columns: adminCashColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminCashDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminCashDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminCashDao) Columns() AdminCashColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminCashDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminCashDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminCashDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
