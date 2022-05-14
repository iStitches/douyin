# douyin
## 文件目录
|文件|解释|
|----|----|
|config|配置文件对应的结构体定义|
|controller|业务层|
|dao|操作数据库给业务|
|controller|提供数据|
|forms|字段验证的struct|
|global|定义全局变量|
|initialize|服务初始化|
|logs|日志存储|
|middlewares|中间件|
|models|数据库字段定义|
|Response|统一封装|
|responsestatic|资源文件夹|
|router|路由|
|setting-dev.yaml|配置文件|
|main.go|服务启动文件|

## git commit 提交规范
提交建议规范一点，这里找了一个链接[掘金](https://juejin.cn/post/6917117275380137998)
- feat: 新功能（feature）
- fix: 修补问题
- docs: 更新文档
- refactor: 重构（即不是新增功能，也不是修改bug的代码变动）
- build: 构建编译相关的改动
- chore: bump version to [版本号]
- test: 增加测试
- style: 格式变更（不影响代码运行的变动）
