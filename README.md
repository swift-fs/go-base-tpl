# HotGo-V2

## 平台简介
* 基于全新GoFrame2+Vue3+NaiveUI+uniapp开发的全栖框架，为二次开发而生，适合中小型完整应用开发。
* 前端采用Naive-Ui-Admin、Vue、Naive UI、uniapp。

### 使用文档

[安装文档](docs/guide-zh-CN/start-installation.md) · [本地文档](docs/guide-zh-CN/README.md) · [更新历史](docs/guide-zh-CN/start-update-log.md) · [常见问题](docs/guide-zh-CN/start-issue.md)


## 感谢(排名不分先后)
> gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf)
>
> naive-ui [https://www.naiveui.com](https://www.naiveui.com)
>
> naive-ui-admin [https://github.com/jekip/naive-ui-admin](https://github.com/jekip/naive-ui-admin)
>
> websocket [https://github.com/gorilla/websocket](github.com/gorilla/websocket)
> 
> casbin [https://github.com/casbin/casbin](https://github.com/casbin/casbin)
>
> gopay [https://github.com/go-pay/gopay](https://github.com/go-pay/gopay)


## 项目初始配置之前的操作(`gf gen dao`之前)

**全局搜索替换数据库表前缀可跳过1,2两个步骤.**
1. `server/storage/data/hotgo.sql`,全局替换表前缀,手动导入该`sql文件`进入数据库
2. `server/storage/data/generate/**/*.sql`,搜索该目录下的所有`sql文件`进行表前缀替换
3. `server/hack/config.yaml`,搜索`TODO`进行自定义修改
4. `server/manifest/config/config.example.yaml`,同目录下复制出一个`config.yaml`文件,搜索`TODO`进行自定义修改
5. 删除`server/internal/dao`,`server/internal/model/do`,`server/internal/model/entity`
6. 切换到`server`目录,执行`gf gen dao`重新生成ORM相关文件
7. `web/.env.development`,`web/.env.production`修改对应的请求地址和接口前缀

**剩下的过程可按照HotGo官方文档进行操作.**
## 自用功能修改

- 去除手机号登录
- 去除注册相关内容
- 去除演示
- TCP服务可以选择开启和关闭(使用到后台编辑定时任务时需要开启)

## License
[MIT © HotGo-2024](./LICENSE)
  


  

