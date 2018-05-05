package main

import (
	"flag"
	"fmt"
	"justtest/lib"
)

/*
ex:
 go run server.go -mode server -p 8080
 go run server.go -mode client -file target  -h 192.168.43.58:8080
*/

func main() {
	var mode *string = flag.String("mode", "server", "Use -mode <filesource>")
	var filePath *string = flag.String("file", "nothing", "Use -file <filesource>")
	var host *string = flag.String("h", "127.0.0.1:8083", "Use -h <filesource>")
	var listenPort *string = flag.String("p", "8083", "Use -p <filesource>")
	flag.Parse()

	if *mode == "server" {
		lib.Receive(*listenPort)
	} else if *mode == "client" {
		lib.Send(*filePath, *host)
	} else {
		fmt.Println("No such mode.")
	}
}
