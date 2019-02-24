package main

import (
  "fmt"
  "reflect"
)

type Enum int

const (
  Zero Enum = 0
)

func empty() {

}

func main() {
  // 获取各种数据类型的 type name kind
  a := 10
  fmt.Println(reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind())

  fmt.Println(reflect.TypeOf(empty), reflect.TypeOf(empty).Kind())

  scene := make(map[string]int)
  fmt.Println(reflect.TypeOf(scene), reflect.TypeOf(scene).Kind())

  type cat struct {
  }

  typeOfCat := reflect.TypeOf(cat{})
  fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

  typeOfA := reflect.TypeOf(Zero)
  fmt.Println(typeOfA.Name(), typeOfA.Kind())

}

// Kind取值

// type Kind uint

// const (
//     Invalid Kind = iota  // 非法类型
//     Bool                 // 布尔型
//     Int                  // 有符号整型
//     Int8                 // 有符号8位整型
//     Int16                // 有符号16位整型
//     Int32                // 有符号32位整型
//     Int64                // 有符号64位整型
//     Uint                 // 无符号整型
//     Uint8                // 无符号8位整型
//     Uint16               // 无符号16位整型
//     Uint32               // 无符号32位整型
//     Uint64               // 无符号64位整型
//     Uintptr              // 指针
//     Float32              // 单精度浮点数
//     Float64              // 双精度浮点数
//     Complex64            // 64位复数类型
//     Complex128           // 128位复数类型
//     Array                // 数组
//     Chan                 // 通道
//     Func                 // 函数
//     Interface            // 接口
//     Map                  // 映射
//     Ptr                  // 指针
//     Slice                // 切片
//     String               // 字符串
//     Struct               // 结构体
//     UnsafePointer        // 底层指针
// )