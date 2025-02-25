// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminNoticeDao is the data access object for table gc_admin_notice.
type AdminNoticeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AdminNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// AdminNoticeColumns defines and stores column names for table gc_admin_notice.
type AdminNoticeColumns struct {
	Id        string // 公告ID
	Title     string // 公告标题
	Type      string // 公告类型
	Tag       string // 标签
	Content   string // 公告内容
	Receiver  string // 接收者
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 公告状态
	CreatedBy string // 发送人
	UpdatedBy string // 修改人
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// adminNoticeColumns holds the columns for table gc_admin_notice.
var adminNoticeColumns = AdminNoticeColumns{
	Id:        "id",
	Title:     "title",
	Type:      "type",
	Tag:       "tag",
	Content:   "content",
	Receiver:  "receiver",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAdminNoticeDao creates and returns a new DAO object for table data access.
func NewAdminNoticeDao() *AdminNoticeDao {
	return &AdminNoticeDao{
		group:   "default",
		table:   "gc_admin_notice",
		columns: adminNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminNoticeDao) Columns() AdminNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
