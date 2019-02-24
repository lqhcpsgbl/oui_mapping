package main

import (
  "fmt"
  "reflect"
)

func main() {
  type cat struct {
    Name string
    Age int
    Type int `json:"type" id:"100"` // 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
  }

  ins := cat{Name: "mimi", Type: 1, Age: 3}

  typeOfCat := reflect.TypeOf(ins)

  for i := 0; i < typeOfCat.NumField(); i++ {
    fieldType := typeOfCat.Field(i)
    fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
  }

  if catType, ok := typeOfCat.FieldByName("Type"); ok {
    // 从tag中取出需要的tag
    fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
  }
}
