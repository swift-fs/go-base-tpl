// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCodes is the golang structure of table gc_sys_gen_codes for DAO operations like Where/Data.
type SysGenCodes struct {
	g.Meta        `orm:"table:gc_sys_gen_codes, do:true"`
	Id            interface{} // 生成ID
	GenType       interface{} // 生成类型
	GenTemplate   interface{} // 生成模板
	VarName       interface{} // 实体命名
	Options       *gjson.Json // 配置选项
	DbName        interface{} // 数据库名称
	TableName     interface{} // 主表名称
	TableComment  interface{} // 主表注释
	DaoName       interface{} // 主表dao模型
	MasterColumns *gjson.Json // 主表字段
	AddonName     interface{} // 插件名称
	Status        interface{} // 生成状态
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
