## gin-swagger-demo

### 介绍
gin 结合 swagger 简单封装，方便开箱即用，API服务及接口文档

### 目录
1. admin          -- 管理后台api
2. app            -- 客户端api
3. conf           -- 配置文件
4. docs           -- 文档
5. middleware     -- 中间件
6. models         -- 数据层
7. pkg            -- 一些工具包
8. routers        -- 路由（多模块 如:admin、app）
9. service        -- 方便api与models交互
10. sql           -- 示例sql

### swagger 
1. go install github.com/swaggo/swag/cmd/swag@latest
2. swag init  // 注意，一定要和 main.go 处于同一级目录
3. 初始化命令，在根目录生成一个docs文件夹


### admin user
username: admin 

password: admin