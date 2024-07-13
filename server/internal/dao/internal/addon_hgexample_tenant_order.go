// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonHgexampleTenantOrderDao is the data access object for table gc_addon_hgexample_tenant_order.
type AddonHgexampleTenantOrderDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns AddonHgexampleTenantOrderColumns // columns contains all the column names of Table for convenient usage.
}

// AddonHgexampleTenantOrderColumns defines and stores column names for table gc_addon_hgexample_tenant_order.
type AddonHgexampleTenantOrderColumns struct {
	Id          string // 主键
	TenantId    string // 租户ID
	MerchantId  string // 商户ID
	UserId      string // 用户ID
	ProductName string // 购买产品
	OrderSn     string // 订单号
	Money       string // 充值金额
	Remark      string // 备注
	Status      string // 订单状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// addonHgexampleTenantOrderColumns holds the columns for table gc_addon_hgexample_tenant_order.
var addonHgexampleTenantOrderColumns = AddonHgexampleTenantOrderColumns{
	Id:          "id",
	TenantId:    "tenant_id",
	MerchantId:  "merchant_id",
	UserId:      "user_id",
	ProductName: "product_name",
	OrderSn:     "order_sn",
	Money:       "money",
	Remark:      "remark",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAddonHgexampleTenantOrderDao creates and returns a new DAO object for table data access.
func NewAddonHgexampleTenantOrderDao() *AddonHgexampleTenantOrderDao {
	return &AddonHgexampleTenantOrderDao{
		group:   "default",
		table:   "gc_addon_hgexample_tenant_order",
		columns: addonHgexampleTenantOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonHgexampleTenantOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonHgexampleTenantOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonHgexampleTenantOrderDao) Columns() AddonHgexampleTenantOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonHgexampleTenantOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonHgexampleTenantOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonHgexampleTenantOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
