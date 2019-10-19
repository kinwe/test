package utils

import (
	"fmt"
	"testing"

	gcv "github.com/smartystreets/goconvey/convey"
)

func TestInitStack(t *testing.T) {
	s := InitStack()
	gcv.Convey("栈不应该为空", t, func() {
		gcv.So(s, gcv.ShouldNotBeNil)
	})
}

func TestPush(t *testing.T) {
	s := InitStack()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	gcv.Convey("入栈测试", t, func() {
		gcv.Convey("栈大小应该为10000", func() {
			gcv.So(s.length, gcv.ShouldEqual, 100000)
		})
		gcv.Convey("栈中元素应该为0-9999", func() {
			gcv.So(func() bool {
				for i := 0; i < 100000; i++ {
					if s.data[i] != i {
						return false
					}
				}
				return true
			}(), gcv.ShouldEqual, true)
		})
	})
}

func TestPop(t *testing.T) {
	gcv.Convey("出栈测试", t, func() {
		gcv.Convey("栈中元素应该为9999-0\n", func() {
			gcv.So(func() bool {
				s := InitStack()
				for i := 0; i < 100000000; i++ {
					s.Push(i)
				}
				for i := 99999999; i > -1; i-- {
					t := s.Pop().(int)
					if t != i {
						return false
					}
					fmt.Println(t)
				}
				return true
			}(), gcv.ShouldEqual, true)
		})
	})
}

func TestPeek(t *testing.T) {
	gcv.Convey("栈顶操作", t, func() {
		s := InitStack()
		s.Push(1)
		s.Push(2)
		tmp := s.Peek().(int)
		gcv.So(tmp, gcv.ShouldEqual, 2)
	})
}
