package main

import (
    "fmt"
    "log"

    "github.com/google/gopacket/pcap"
)

func main() {

    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }

    // 获取本机网络接口和描述
    fmt.Println("Devices found:")
    for _, device := range devices {
        fmt.Println("\nName: ", device.Name)
        fmt.Println("Description: ", device.Description)
        fmt.Println("Devices addresses: ", device.Description)

        for _, address := range device.Addresses {
            fmt.Println("- IP address: ", address.IP)
            fmt.Println("- Subnet mask: ", address.Netmask)
        }
    }
}