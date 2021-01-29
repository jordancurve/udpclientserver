package main

import (
  "fmt"
  "log"
  "net"
  "encoding/json"
  "os"
  "time"
)

type Message struct {
  Counter int
  Time time.Time
}

func main() {
  hostport := os.Args[1]
  addr, err := net.ResolveUDPAddr("udp", hostport)
  if err != nil {
    log.Fatal(err)
  }
  var conn *net.UDPConn
  counter := 0
  var connected bool
  for t := range time.Tick(1000 * time.Millisecond) {
    counter++
    if !connected {
      var err error
      fmt.Printf("connecting to %s... ", hostport)
      conn, err = net.DialUDP("udp", nil, addr)
      if err != nil {
        fmt.Printf("error connecting: %v\n", err)
	continue
      } else {
        fmt.Printf("OK\n")
	connected = true
      }
    }
    msg, _ := json.Marshal(Message{counter, t})
    fmt.Printf("sending to %s: %s\n", hostport, string(msg))
    if _, err := conn.Write(msg); err != nil {
      fmt.Printf("error: %v\n", err)
      connected = false
    }
  }
}
