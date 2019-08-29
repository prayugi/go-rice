package main

import (
	"flag"
	"fmt"
)

var fs = flag.NewFlagSet("service", flag.ExitOnError)
var httpAddr = fs.String("http.addr", ":8080", "http listen address")

func main() {
	run()

}

func run() {
	//run the service
	fmt.Println(*httpAddr)
}
