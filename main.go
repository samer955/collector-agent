package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samer955/collector-agent/config"
	"github.com/samer955/collector-agent/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	collector := service.NewCollector()
	collector.Start()

	defer config.GetDbConfig().Connection.Close()

	//Run the program till its stopped (forced)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
