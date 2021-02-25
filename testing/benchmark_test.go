/*
@Time : 2020/11/4 下午11:13
@Author : xiukang
@File : benchmark_test.go
@Software: GoLand
*/
package testings

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		a := []string{"hel1lo", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil", t, func() {
		So(StringSliceEqual(nil, nil), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when a != nil  && b != nil", t, func() {
		a := []string{"hello", "world"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}
