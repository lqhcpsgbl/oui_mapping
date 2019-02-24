package main

import (
  "fmt"
  "reflect"
)

func main() {
  // 空的int指针
  var a *int
  // 似于C语言的 a == NULL操作
  fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

  fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

  fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

  s := struct{}{}
  fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
  fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("test").IsValid())

  m := map[int]int{}
  fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}
