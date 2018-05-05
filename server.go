package main

import (
	"flag"
	"justtest/lib"
)

var ENV *string = flag.String("d", "development", "Enviorment development staging production")

func main() {
	var filePath *string = flag.String("file", "musicfile", "Use -file <filesource>")
	flag.Parse()
	lib.Send(*filePath)
}
