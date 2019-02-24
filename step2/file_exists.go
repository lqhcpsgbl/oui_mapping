package main

import (
    "fmt"
    "os"
)

// 判断文件是否存在
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}

func main() {
    fmt.Println(Exists("./defer.md"))
    fmt.Println(Exists("step2.go"))
}