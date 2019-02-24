package main

// 学习Golang http 模块 之 head 方法 获取 文件最近修改时间

import (
    "fmt"
    "net/http"
    "log"
    "reflect"
)

func main() {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Head(url)

    if err != nil {
        log.Fatalln(err)
    }

    // 通过 反射 获取类型，只是为了了解语法
    typeof_resp := reflect.TypeOf(resp)
    fmt.Println(typeof_resp, typeof_resp.Kind())

    headers := resp.Header

    fmt.Println(reflect.TypeOf(headers))

    for key, value := range headers {
        fmt.Println("HTTP resp header: ", key, " value: ", value[0])
    }

    last_modified := headers.Get("Last-Modified")
    fmt.Println("OUI 文件修改时间是", last_modified)
}

// 运行结果：

// *http.Response ptr
// http.Header
// HTTP resp header:  Content-Length  value:  4083557
// HTTP resp header:  Last-Modified  value:  Thu, 14 Feb 2019 23:04:08 GMT
// HTTP resp header:  Connection  value:  keep-alive
// HTTP resp header:  Etag  value:  "5c65f3e8-3e4f65"
// HTTP resp header:  Accept-Ranges  value:  bytes
// HTTP resp header:  Server  value:  nginx/1.12.0
// HTTP resp header:  Date  value:  Fri, 15 Feb 2019 00:58:51 GMT
// HTTP resp header:  Content-Type  value:  text/plain
// OUI 文件修改时间是 Thu, 14 Feb 2019 23:04:08 GMT
