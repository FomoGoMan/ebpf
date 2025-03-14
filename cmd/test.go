package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	collector "github.com/FomoGoMan/container-netprobe"
)

func main() {
	// usage ./main <container id>
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <container-id>")
		return
	}

	containerID := os.Args[1]
	fmt.Printf("Target Container ID: %s\n", containerID)

	monitor, err := collector.NewGeneralCollector(containerID)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = monitor.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer monitor.Cleanup()

	go func() {
		for {
			in, out := monitor.CollectTotal(monitor.CGroupId())
			fmt.Printf("[%s] In: %d bytes, Out: %d bytes\n", time.Now().Format("15:04:05"), in, out)
			time.Sleep(2 * time.Second)
		}
	}()

	// wait ctrl-c
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
