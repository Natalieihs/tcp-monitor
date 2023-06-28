package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

func main() {
	for {
		conns, err := net.Connections("tcp")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var total, timeWait int
		for _, conn := range conns {
			if conn.Status == "TIME_WAIT" {
				timeWait++
			}
			total++
		}

		fmt.Printf("Total TCP connections: %d\n", total)
		fmt.Printf("TCP connections in TIME_WAIT: %d\n", timeWait)

		time.Sleep(5 * time.Second) // 等待5秒
	}
}
