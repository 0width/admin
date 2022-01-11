# 简介

本项目目前正在开发中。。。

- 本项目使用 https://github.com/go-spring/go-spring 作为底层IoC容器， 实现了半自动扫描依赖
- web框架使用 https://github.com/gin-gonic/gin
- 数据库orm使用https://github.com/go-gorm/gorm， 实现了数据库migrate, 并记录ddl语句到文件
- 依赖golang 1.16的embed， 嵌入前端代码， 支持304 NotModified
- 前端使用 https://github.com/PanJiaChen/vue-admin-template 作为基础框架
- 支持jwt认证， 支持前端动态路由， 权限配置



# IoC和半自动扫描

半自动扫描依赖 go:generate, 可以手动执行 go generate， 或者在开发工具中配置 pre run

生成的依赖保存在imports.go中

```
//go:generate go run generate/auto_imports/auto_imports.go
```

## 扫描依赖的基本原理

使用 go/parser 解析项目下所有的 go 文件， 找到有依赖git.xios.club/xios/gc的包， 这些包就是需要自动引入的包， 这些包内都有注册bean的方法

# 自动迁移和DDL语句保存

这部分内容在 business/pogo/entity/system/entity.go

使用gorm/logger的Trace接口， 识别出DDL语句，并写入到指定的文件