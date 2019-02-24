package main

// 下载 OUI 文件，并保存本地文件

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

// 得到最新修改时间
func GetLastModified() string {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Head(url)

    if err != nil {
        log.Fatalln(err)
    }

    last_modified := resp.Header.Get("Last-Modified")
    return last_modified
}

// 下载OUI文件
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

// 判断文件是否存在
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}

// 删除本地OUI data文件
func DeleteDataFile() {
    files, _ := filepath.Glob("./*.data")

    for _, f := range files {
        os.Remove(f)
    }
}

// 将内容写入文件
func WriteContent(filename string, content string) {

    f, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    content_len, err := f.WriteString(content)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(content_len, "个字符写入文件！")
    defer f.Close()
}

func main() {
    last_modified := GetLastModified()
    file_exists := Exists(last_modified)
    if file_exists {
        fmt.Println("最新OUI文件已存在！")
        os.Exit(1)
    }
    oui_content := DownloadOUI()
    DeleteDataFile()
    filename := last_modified + ".data"
    WriteContent(filename, oui_content)
}
