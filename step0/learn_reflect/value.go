package main

import (
  "fmt"
  "reflect"
)

func main() {
  var a int = 1024
  valueOfA := reflect.ValueOf(a)

  fmt.Println(valueOfA, reflect.TypeOf(valueOfA).Name())

  // 获取interface{}类型的值, 通过类型断言转换
  var getA int = valueOfA.Interface().(int)
  fmt.Println(getA)

  // 获取64位的值, 强制类型转换为int类型
  var getA2 int = int(valueOfA.Int())
  fmt.Println(getA2)


  type dummy struct {
    a int
    b string

    // 嵌入字段
    float32
    bool

    next *dummy
  }

  d := reflect.ValueOf(dummy{
      a: 100,
      b: "hello",
      next: &dummy{},
  })

  fmt.Println("NumField", d.NumField())

  floatField := d.Field(2)
  fmt.Println("Field", floatField.Type())

  fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
  // []int{4,0} 中的 4 表示，在 dummy 结构中索引值为 4 的成员，也就是 next。next 的类型为 dummy，也是一个结构体，因此使用 []int{4,0} 中的 0 继续在 next 值的基础上索引，结构为 dummy 中索引值为 0 的 a 字段，类型为 int。
  fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
}
