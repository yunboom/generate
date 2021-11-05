package check

import (
	c "github.com/smartystreets/goconvey/convey"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/datebase/driver"
	"testing"
)

const MysqlDSN = "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func TestGenBaseStructs(t *testing.T) {
	c.Convey("获取db", t, func() {
		gorm, err := datebase.NewGorm(driver.Mysql, MysqlDSN)
		c.So(err, c.ShouldBeNil)
		c.Convey("获取绑定结构体", func() {
			s, err := GenBaseStructs(gorm, "users", "Users")
			c.So(err, c.ShouldBeNil)
			c.So(s, c.ShouldNotBeNil)
		})

	})
}
