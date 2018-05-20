package main

import (
	"flag"
	"fmt"
)

var addr = flag.String("addr", ":6969", "http service address")

// go run main.go --addr="whatsyourna.me:8888"

func main() {

	// loading
	configsrv()

}

func configsrv() {
	flag.Parse()
	fmt.Println(*addr)
}
