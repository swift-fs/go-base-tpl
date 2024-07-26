// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package crons

import (
	"context"
	"fmt"
	"hotgo/internal/library/cron"

	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	cron.Register(Schedule)
}

// Schedule 任务调度
var Schedule = &cSchedule{name: "schedule"}

type cSchedule struct {
	name string
}

func (c *cSchedule) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cSchedule) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	fmt.Println("cron schedule Execute:", gtime.Now())
	return
}
