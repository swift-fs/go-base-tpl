// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleMenu is the golang structure of table gc_admin_role_menu for DAO operations like Where/Data.
type AdminRoleMenu struct {
	g.Meta `orm:"table:gc_admin_role_menu, do:true"`
	RoleId interface{} // 角色ID
	MenuId interface{} // 菜单ID
}
