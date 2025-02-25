// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure of table gc_addon_hgexample_table for DAO operations like Where/Data.
type AddonHgexampleTable struct {
	g.Meta      `orm:"table:gc_addon_hgexample_table, do:true"`
	Id          interface{} // ID
	Pid         interface{} // 上级ID
	Level       interface{} // 树等级
	Tree        interface{} // 关系树
	CategoryId  interface{} // 分类ID
	Flag        *gjson.Json // 标签
	Title       interface{} // 标题
	Description interface{} // 描述
	Content     interface{} // 内容
	Image       interface{} // 单图
	Images      *gjson.Json // 多图
	Attachfile  interface{} // 附件
	Attachfiles *gjson.Json // 多附件
	Map         *gjson.Json // 动态键值对
	Star        interface{} // 推荐星
	Price       interface{} // 价格
	Views       interface{} // 浏览次数
	ActivityAt  *gtime.Time // 活动时间
	StartAt     *gtime.Time // 开启时间
	EndAt       *gtime.Time // 结束时间
	Switch      interface{} // 开关
	Sort        interface{} // 排序
	Avatar      interface{} // 头像
	Sex         interface{} // 性别
	Qq          interface{} // qq
	Email       interface{} // 邮箱
	Mobile      interface{} // 手机号码
	Hobby       *gjson.Json // 爱好
	Channel     interface{} // 渠道
	CityId      interface{} // 所在城市
	Remark      interface{} // 备注
	Status      interface{} // 状态
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
