// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysServeLicenseDao is the data access object for table gc_sys_serve_license.
type SysServeLicenseDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SysServeLicenseColumns // columns contains all the column names of Table for convenient usage.
}

// SysServeLicenseColumns defines and stores column names for table gc_sys_serve_license.
type SysServeLicenseColumns struct {
	Id           string // 许可ID
	Group        string // 分组
	Name         string // 许可名称
	Appid        string // 应用ID
	SecretKey    string // 应用秘钥
	RemoteAddr   string // 最后连接地址
	OnlineLimit  string // 在线限制
	LoginTimes   string // 登录次数
	LastLoginAt  string // 最后登录时间
	LastActiveAt string // 最后心跳
	Routes       string // 路由表，空使用默认分组路由
	AllowedIps   string // IP白名单
	EndAt        string // 授权有效期
	Remark       string // 备注
	Status       string // 状态
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
}

// sysServeLicenseColumns holds the columns for table gc_sys_serve_license.
var sysServeLicenseColumns = SysServeLicenseColumns{
	Id:           "id",
	Group:        "group",
	Name:         "name",
	Appid:        "appid",
	SecretKey:    "secret_key",
	RemoteAddr:   "remote_addr",
	OnlineLimit:  "online_limit",
	LoginTimes:   "login_times",
	LastLoginAt:  "last_login_at",
	LastActiveAt: "last_active_at",
	Routes:       "routes",
	AllowedIps:   "allowed_ips",
	EndAt:        "end_at",
	Remark:       "remark",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSysServeLicenseDao creates and returns a new DAO object for table data access.
func NewSysServeLicenseDao() *SysServeLicenseDao {
	return &SysServeLicenseDao{
		group:   "default",
		table:   "gc_sys_serve_license",
		columns: sysServeLicenseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysServeLicenseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysServeLicenseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysServeLicenseDao) Columns() SysServeLicenseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysServeLicenseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysServeLicenseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysServeLicenseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
