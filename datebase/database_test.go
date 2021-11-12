package datebase

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSnakeToHump(t *testing.T) {
	c.Convey("测试下划线转驼峰", t, func() {
		arr := map[string]string{
			"_snake_hump": "SnakeHump",
			"snake_hump":  "SnakeHump",
			"snakeHump_":  "SnakeHump",
		}

		for key, want := range arr {
			hump := snakeToHump(key)
			c.So(hump, c.ShouldEqual, want)
		}
	})
}
