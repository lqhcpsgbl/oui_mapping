package main

// 下载 OUI 文件，并保存本地文件

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/cavaliercoder/grab"
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
func DownloadOUI(filename string) string {
    url := "http://standards-oui.ieee.org/oui.txt"

    client := grab.NewClient()
    req, _ := grab.NewRequest(".", url)

    // start download
    fmt.Printf("Downloading %v...\n", req.URL())
    resp := client.Do(req)
    fmt.Printf("  %v\n", resp.HTTPResponse.Status)

    // start UI loop
    t := time.NewTicker(500 * time.Millisecond)
    defer t.Stop()

Loop:
    for {
        select {
        case <-t.C:
            fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
                resp.BytesComplete(),
                resp.Size,
                100*resp.Progress())

        case <-resp.Done:
            // download is complete
            break Loop
        }
    }

    // check for errors
    if err := resp.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Download saved to ./%v \n", resp.Filename)

    // 重命名OUI文件
    err := os.Rename("./"+resp.Filename, "./"+filename)
    if err != nil {
        log.Fatalln(err)
    }

    return filename
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

// 将文件内容读取出来
func ReadFileContent(filename string) string {

    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln(err)
    }

    return string(bytes)
}

// 解析文件得到 MAC地址前缀 和 网卡制造商
func ParseOUI(page_content string) map[string]string {

    prefix_pro := map[string]string{}
    lines := strings.Split(page_content, "\n")
    for _, line := range lines {
        if strings.Contains(line, "(hex)") {
            prefix_pro_info := strings.Replace(line, "   (hex)\t\t", ";", 1)
            prefix_pro_info_list := strings.Split(prefix_pro_info, ";")

            prefix_pro[prefix_pro_info_list[0]] = prefix_pro_info_list[1]
        }
    }
    return prefix_pro
}

func main() {

    last_modified := GetLastModified()
    filename := last_modified + ".data"
    fmt.Println(filename)
    file_exists := Exists(filename)

    oui_content := ""

    if file_exists {
        fmt.Println("最新OUI文件已存在！")
        oui_content = ReadFileContent(filename)
    } else {
        fmt.Println("开始下载OUI文件！")
        DeleteDataFile()
        DownloadOUI(filename)
        oui_content = ReadFileContent(filename)
    }

    prefix_pro_map := ParseOUI(oui_content)

    // fmt.Println(prefix_pro_map["00-50-3B"])
    fmt.Println(len(prefix_pro_map))
}
