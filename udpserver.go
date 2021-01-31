package main
import (
  "encoding/json"
  "fmt"
  "log"
  "net"
  "os"
  "time"
)

type Message struct {
  Counter int
  Time time.Time
}
func main() {
  msg := make([]byte, 2048)
  port := 9876
  addr, err := net.ResolveUDPAddr("udp", os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  ser, err := net.ListenUDP("udp", addr)
  fmt.Printf("listening on port %d\n", port)
  if err != nil {
    log.Fatal(err)
  }
  var oldCounter int
  var oldTime time.Time
  for {
    n, remoteaddr, err := ser.ReadFromUDP(msg)
    if err != nil {
      fmt.Printf("Error: %v\n", err) 
    }
    m := Message{}
    if err := json.Unmarshal(msg[:n], &m); err != nil {
      log.Fatal(err)
    }
    fmt.Printf("read from %v: counter=%d time=%v\n", remoteaddr, m.Counter, m.Time)
    if oldCounter != 0 && m.Counter != oldCounter + 1 {
      fmt.Printf("\ncounter gap\n\n")
    }
    if oldTime != (time.Time{}) && m.Time.Sub(oldTime) >= 2 * time.Second {
      fmt.Printf("\ntime gap\n\n")
    }
    oldCounter = m.Counter
  }
}
