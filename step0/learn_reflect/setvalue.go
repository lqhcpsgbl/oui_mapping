package main

import (
  "fmt"
  "reflect"
)

func main() {
  var num int = 42
  valueOfA := reflect.ValueOf(&num)

  valueOfA = valueOfA.Elem()

  // &num是可寻址的，所以可以set值
  fmt.Println(valueOfA.CanSet())
  fmt.Println(valueOfA.CanAddr())

  valueOfA.SetInt(1)
  fmt.Println(valueOfA.Int())
  a := valueOfA.Int()
  fmt.Println(a)

  // 结构体成员中，如果字段没有被导出，即便不使用反射也可以被访问，但不能通过反射修改
  type dog struct {
    LegCount int
  }

  // LegCount 是可以被导出的 大写字母 开头
  valueOfDog := reflect.ValueOf(&dog{})
  // 取出dog实例地址的元素
  valueOfDog = valueOfDog.Elem()

  vLegCount := valueOfDog.FieldByName("LegCount")

  vLegCount.SetInt(4)
  fmt.Println(vLegCount.Int())
}
