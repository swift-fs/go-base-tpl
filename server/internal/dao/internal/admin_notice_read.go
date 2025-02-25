// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminNoticeReadDao is the data access object for table gc_admin_notice_read.
type AdminNoticeReadDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminNoticeReadColumns // columns contains all the column names of Table for convenient usage.
}

// AdminNoticeReadColumns defines and stores column names for table gc_admin_notice_read.
type AdminNoticeReadColumns struct {
	Id        string // 记录ID
	NoticeId  string // 公告ID
	MemberId  string // 会员ID
	Clicks    string // 已读次数
	UpdatedAt string // 更新时间
	CreatedAt string // 阅读时间
}

// adminNoticeReadColumns holds the columns for table gc_admin_notice_read.
var adminNoticeReadColumns = AdminNoticeReadColumns{
	Id:        "id",
	NoticeId:  "notice_id",
	MemberId:  "member_id",
	Clicks:    "clicks",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewAdminNoticeReadDao creates and returns a new DAO object for table data access.
func NewAdminNoticeReadDao() *AdminNoticeReadDao {
	return &AdminNoticeReadDao{
		group:   "default",
		table:   "gc_admin_notice_read",
		columns: adminNoticeReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminNoticeReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminNoticeReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminNoticeReadDao) Columns() AdminNoticeReadColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminNoticeReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminNoticeReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminNoticeReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
