package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPredixTree(t *testing.T) {
	preTree := NewPrefixTree()

	preTree.Insert("apple")
	preTree.Insert("坏人")

	Convey("测试直接包含敏感词的情况: ", t, func() {
		So(preTree.Contains("坏人"), ShouldEqual, true)
		So(preTree.Contains("和珅你可真是个坏人"), ShouldEqual, true)
		So(preTree.Contains("你吃apple不吃"), ShouldEqual, true)
	})
	Convey("测试直接空格逗号等分割敏感词的情况: ", t, func() {
		So(preTree.Contains("你是坏, 人啊"), ShouldEqual, true)
		So(preTree.Contains("你是坏  人啊"), ShouldEqual, true)
	})
	Convey("测试不包含敏感词情况: ", t, func() {
		So(preTree.Contains("appl"), ShouldEqual, false)
		So(preTree.Contains("坏"), ShouldEqual, false)
		So(preTree.Contains("你个坏蛋"), ShouldEqual, false)
	})

}
