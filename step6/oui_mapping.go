package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/cavaliercoder/grab"

    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

var (
    device       string = "en0" // Linux一般为 eth0, Mac是 en0
    snapshot_len int32  = 1024
    promiscuous  bool   = true
    err          error
    timeout      time.Duration = 30 * time.Second
    handle       *pcap.Handle
)

// 得到最新修改时间
func getLastModified() string {
    url := "http://standards-oui.ieee.org/oui.txt"
    resp, err := http.Head(url)

    if err != nil {
        log.Fatalln(err)
    }

    last_modified := resp.Header.Get("Last-Modified")
    return last_modified
}

// 下载OUI文件
func downloadOUI(filename string) string {
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
func exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

// 删除本地OUI data文件
func deleteDataFile() {
    files, _ := filepath.Glob("./*.data")

    for _, f := range files {
        os.Remove(f)
    }
}

// 将文件内容读取出来
func readFileContent(filename string) string {

    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln(err)
    }

    return string(bytes)
}

// 解析文件得到 MAC地址前缀 和 网卡制造商
func parseOUI(page_content string) map[string]string {

    prefix_pro := map[string]string{}
    lines := strings.Split(page_content, "\n")
    for _, line := range lines {
        if strings.Contains(line, "(hex)") {
            prefix_pro_info := strings.Replace(line, "   (hex)\t\t", ";", 1)
            prefix_pro_info_list := strings.Split(prefix_pro_info, ";")
            productor := prefix_pro_info_list[1]
            mac_prefix := strings.Replace(prefix_pro_info_list[0], "-", ":", -1)
            prefix_pro[mac_prefix] = productor
        }
    }
    return prefix_pro
}

// 判断是否公有地址
func IsPublicIP(IP net.IP) bool {
    if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
        return false
    }
    if ip4 := IP.To4(); ip4 != nil {
        switch true {
        case ip4[0] == 10:
            return false
        case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
            return false
        case ip4[0] == 192 && ip4[1] == 168:
            return false
        default:
            return true
        }
    }
    return false
}

// 解析包结构，返回源MAC地址，源IP地址
func DecodePacketInfo(packet gopacket.Packet) (source_mac string, source_ip string) {
    source_mac = ""
    source_ip = ""
    // 是否是以太网数据包
    ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
    if ethernetLayer != nil {
        // 类型断言
        ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
        source_mac = strings.ToUpper(ethernetPacket.SrcMAC.String())
    }

    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        ip, _ := ipLayer.(*layers.IPv4)
        source_ip = ip.SrcIP.String()
    }
    return
}

func getPacketAndParse(prefix_pro_map map[string]string) {
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }

    defer handle.Close()

    local_mac_table := make(map[string]string)

    // 获取数据包并输出
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        source_mac, source_ip := DecodePacketInfo(packet)
        // 对私有IP输出源MAC地址，源IP地址，网卡供应商
        if !IsPublicIP(net.ParseIP(source_ip)) {
            mac_prefix := source_mac[0:8]
            if _, ok := local_mac_table[source_mac]; !ok && source_ip != "" {
                local_mac_table[source_mac] = mac_prefix
                fmt.Println(source_mac, source_ip, prefix_pro_map[mac_prefix])
            }
        }
    }
}

func main() {

    last_modified := getLastModified()
    filename := last_modified + ".data"
    fmt.Println(filename)
    file_exists := exists(filename)

    oui_content := ""

    if file_exists {
        fmt.Println("最新OUI文件已存在！")
        oui_content = readFileContent(filename)
    } else {
        fmt.Println("开始下载最新OUI文件！")
        deleteDataFile()
        downloadOUI(filename)
        oui_content = readFileContent(filename)
    }

    prefix_pro_map := parseOUI(oui_content)

    getPacketAndParse(prefix_pro_map)
}
