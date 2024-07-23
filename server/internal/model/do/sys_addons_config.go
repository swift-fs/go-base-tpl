// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAddonsConfig is the golang structure of table gc_sys_addons_config for DAO operations like Where/Data.
type SysAddonsConfig struct {
	g.Meta       `orm:"table:gc_sys_addons_config, do:true"`
	Id           interface{} // 配置ID
	AddonName    interface{} // 插件名称
	Group        interface{} // 分组
	Name         interface{} // 参数名称
	Type         interface{} // 键值类型:string,int,uint,bool,datetime,date
	Key          interface{} // 参数键名
	Value        interface{} // 参数键值
	DefaultValue interface{} // 默认值
	Sort         interface{} // 排序
	Tip          interface{} // 变量描述
	IsDefault    interface{} // 是否为系统默认
	Status       interface{} // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
