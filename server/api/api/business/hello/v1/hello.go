package v1

import "github.com/gogf/gf/v2/frame/g"

type HelloReq struct {
	g.Meta `path:"/hello" method:"get" tags:"Hello" summary:"Hello" dc:"Hello"`
	Name   string `json:"name" dc:"Name"`
}

type HelloRes struct {
	g.Meta `mime:"application/json"`
	Msg    string `json:"msg" dc:"Msg"`
}
