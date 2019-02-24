package main

// 下载 OUI 文件，并打印结果

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

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

    if err != nil {
        log.Fatalln(err)
    }

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
