package main

import (
	"fmt"
	_ "key-value-system/env"
	"key-value-system/router"
	_ "key-value-system/scheduler"
	"os"
)

func main() {
	router.Run(fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT")))
}
