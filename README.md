# generate

#### 介绍

go代码生成器、目标适配：

1. Mysql、Postgresql与Gorm、Xorm
2. 一键生成model、service、handle代码
3. 生成Java Mybatis-Plus风格API

#### 安装教程

```go get github.com/yunboom/generate```

#### 使用说明

1. 生成结构体

```go
package main

import (
	"fmt"
	"github.com/yunboom/generate"
	"github.com/yunboom/generate/config"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/datebase/driver"
)

const MysqlDSN = "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
const PostgresDSN = "host=localhost port=54321 user=postgres dbname=clubdb1 password=root sslmode=disable"

func main() {
	gen := generate.New(config.New(
		config.WithModelPath("/Users/zonst/Downloads/"), //model 代码输出路径
	))
	gen.UseDB(datebase.OpenGorm(driver.Postgres, PostgresDSN))        //使用gorm Postgres
	gen.BindModel(gen.GenModelAs("club_invite_log", "ClubInviteLog")) //绑定模型
	if err := gen.Execute(); err != nil {
		fmt.Println(err)
	}
}
```

2. CRUD接口
    1. Insert
    ```go
   // 插入一条记录
    db, _ = gorm.Open(postgres.Open(PostgresDSN))
    dao = NewUserDao(db.Debug())
    err := dao.Insert(&User{
    		Username: "123",
    		Password: "123",
    		Nick:     "张三",
    	})
   ```

    2. Update

    ```go
    //UPDATE "users" SET "password"='321' WHERE id = 123
    err := dao.UpdateById(123, &User{Password: "321"})
    ```

    3. Delete
   
    ```go
       //DELETE FROM "users" WHERE id = 3
       err := dao.DeleteById(3)
       
       //DELETE FROM "users" WHERE id in (5,6,7)
       err := dao.DeleteBatchIds(5, 6, 7)
       
       //DELETE FROM "users" WHERE "users"."username" = '123' AND id > 10
       wrapper := dao.QueryWrapper(&User{Username: "123"}).Where("id > ?", 10)
       err := dao.DeleteByWrapper(wrapper)
    ```
   
    4. Select
   
   ```go
   //根据id查询
   user, err := dao.SelectById(1)
   
   //查询所有 SELECT * FROM "users"
   wrapper := dao.QueryWrapper(nil)
   userList, err := dao.SelectList(wrapper)
   
   //SELECT * FROM "users" WHERE "users"."username" = '123' AND id > 0 ORDER BY id desc
   wrapper := dao.QueryWrapper(&User{Username: "123"}).Where("id > ?", 0).OrderBy("id desc")
   userList, err := dao.SelectList(wrapper)
   
   //SELECT * FROM "users" WHERE "users"."username" = '123' AND id > 10 ORDER BY "users"."id" LIMIT 1
   wrapper := dao.QueryWrapper(&User{Username: "123"}).Where("id > ?", 10)
   user, err := dao.SelectOne(wrapper)
   
   //SELECT count(*) FROM "users"
   count, err := dao.SelectCount(dao.QueryWrapper(nil))
   
   //SELECT * FROM "users" LIMIT 2 OFFSET 1
   userList, err := dao.SelectPage(2, 1, dao.QueryWrapper(nil))
   
   //[map[Id:1 Nick:张三 Password:123 Username:123] map[Id:2 Nick:李四 Password:456 Username:456]]
   userMaps, err := dao.SelectMaps(dao.QueryWrapper(nil))
   