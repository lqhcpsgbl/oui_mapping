package main

import (
  "fmt"
  "log"
  "time"

  "github.com/google/gopacket"
  "github.com/google/gopacket/layers"
  "github.com/google/gopacket/pcap"
)

var (
  device string = "en0"
  snapshot_len int32  = 1024
  promiscuous  bool   = true
  err          error
  timeout      time.Duration = 30 * time.Second
  handle       *pcap.Handle
)

// 解码包并获取MAC地址相关信息
func decodePacketInfo(packet gopacket.Packet) {
  // 是否是以太网数据包
  ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
  if ethernetLayer != nil {
    fmt.Println("Ethernet layer detected.")
    // 类型断言
    ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
    fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
    fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
    // 一般是 IPv4 或者ARP
    fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
    fmt.Println()
  } 

  ipLayer := packet.Layer(layers.LayerTypeIPv4)
  if ipLayer != nil {
      fmt.Println("IPv4 layer detected.")
      ip, _ := ipLayer.(*layers.IPv4)
      fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
      fmt.Println("Protocol: ", ip.Protocol)
      fmt.Println()
  }
}

func main() {
  handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
  if err != nil {
    log.Fatal(err)
  }

  defer handle.Close()
  // 获取数据包并输出
  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    decodePacketInfo(packet)
  }
}