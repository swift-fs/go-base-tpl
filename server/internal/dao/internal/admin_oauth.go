// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminOauthDao is the data access object for table gc_admin_oauth.
type AdminOauthDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AdminOauthColumns // columns contains all the column names of Table for convenient usage.
}

// AdminOauthColumns defines and stores column names for table gc_admin_oauth.
type AdminOauthColumns struct {
	Id           string // 主键
	MemberId     string // 用户ID
	Unionid      string // 唯一ID
	OauthClient  string // 授权组别
	OauthOpenid  string // 授权开放ID
	Sex          string // 性别
	Nickname     string // 昵称
	HeadPortrait string // 头像
	Birthday     string // 生日
	Country      string // 国家
	Province     string // 省
	City         string // 市
	Status       string // 状态
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
}

// adminOauthColumns holds the columns for table gc_admin_oauth.
var adminOauthColumns = AdminOauthColumns{
	Id:           "id",
	MemberId:     "member_id",
	Unionid:      "unionid",
	OauthClient:  "oauth_client",
	OauthOpenid:  "oauth_openid",
	Sex:          "sex",
	Nickname:     "nickname",
	HeadPortrait: "head_portrait",
	Birthday:     "birthday",
	Country:      "country",
	Province:     "province",
	City:         "city",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewAdminOauthDao creates and returns a new DAO object for table data access.
func NewAdminOauthDao() *AdminOauthDao {
	return &AdminOauthDao{
		group:   "default",
		table:   "gc_admin_oauth",
		columns: adminOauthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminOauthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminOauthDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminOauthDao) Columns() AdminOauthColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminOauthDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminOauthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminOauthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
