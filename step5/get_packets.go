package main

import (
  "fmt"
  "log"
  "time"

  "github.com/google/gopacket"
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

func main() {
  handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
  if err != nil {
    log.Fatal(err)
  }

  defer handle.Close()
  // 获取数据包并输出
  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    fmt.Println(packet)
  }
}