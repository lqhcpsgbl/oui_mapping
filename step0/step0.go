package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Head(url)

    if err != nil {
        log.Fatalln(err)
    }

    headers := resp.Header

    for key, value := range headers {
        fmt.Println("HTTP resp header: ", key, " value: ", value[0])
    }

    last_modified := headers.Get("Last-Modified")
    fmt.Println("OUI 文件修改时间是", last_modified)
}
