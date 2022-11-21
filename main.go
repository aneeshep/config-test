package main

import (
	"fmt"

	"github.com/aneeshep/config-test/config/v2"
)

func main() {
	config.ServiceName = "test-svc1"
	config.LoadConfig()

	port := config.Get("PORT")
	fmt.Println(port)
}
