// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLoginLogDao is the data access object for table gc_sys_login_log.
type SysLoginLogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysLoginLogColumns defines and stores column names for table gc_sys_login_log.
type SysLoginLogColumns struct {
	Id         string // 日志ID
	ReqId      string // 请求ID
	MemberId   string // 用户ID
	Username   string // 用户名
	Response   string // 响应数据
	LoginAt    string // 登录时间
	LoginIp    string // 登录IP
	ProvinceId string // 省编码
	CityId     string // 市编码
	UserAgent  string // UA信息
	ErrMsg     string // 错误提示
	Status     string // 状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// sysLoginLogColumns holds the columns for table gc_sys_login_log.
var sysLoginLogColumns = SysLoginLogColumns{
	Id:         "id",
	ReqId:      "req_id",
	MemberId:   "member_id",
	Username:   "username",
	Response:   "response",
	LoginAt:    "login_at",
	LoginIp:    "login_ip",
	ProvinceId: "province_id",
	CityId:     "city_id",
	UserAgent:  "user_agent",
	ErrMsg:     "err_msg",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysLoginLogDao creates and returns a new DAO object for table data access.
func NewSysLoginLogDao() *SysLoginLogDao {
	return &SysLoginLogDao{
		group:   "default",
		table:   "gc_sys_login_log",
		columns: sysLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysLoginLogDao) Columns() SysLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
