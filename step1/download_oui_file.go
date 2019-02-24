package main

// 下载 OUI 文件，并打印结果

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// 将 step0 代码写成函数
func GetLastModified() string {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Head(url)

    if err != nil {
        log.Fatalln(err)
    }

    last_modified := resp.Header.Get("Last-Modified")
    return last_modified
}

func DownloadOUI() string {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Get(url)

    // Fatalln is equivalent to Println() followed by a call to os.Exit(1).
    if err != nil {
        log.Fatalln(err)
    }

    // 关于 defer 查看 defer.md 笔记 , defer需要写在 判断错误之后
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    return string(content)
}

func main() {
    fmt.Println(GetLastModified())
    fmt.Println(DownloadOUI())
}
