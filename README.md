# generate

#### 介绍
go代码生成

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1. 生成结构体
```go
package main

import (
	"fmt"
	"github.com/yunboom/generate"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/datebase/driver"
)

func main() {
	const MysqlDSN = "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	generator := generate.New(generate.NewConfig(
		generate.WithModelPath("../gen"), //model 代码输出路径
	))
	generator.UseDB(datebase.NewGorm(driver.Mysql, MysqlDSN))   //使用gorm mysql
	generator.BindModel(generator.GenModelAs("users", "Users")) //绑定模型
	if err := generator.Execute(); err != nil {
		fmt.Println(err)
	}
}

```
2. xxxx
3. xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
