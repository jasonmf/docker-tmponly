// CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o setup setup.go
package main

import (
	"fmt"
	"log"
	"os"
)

func fatalIfError(err error, msg string) {
	if err != nil {
		log.Fatal("error ", msg, ": ", err)
	}
}

func main() {
	fatalIfError(os.Mkdir("/tmp", 01777), "creating tmp") // umask prevents 01777 from working here
	fatalIfError(os.Chmod("/tmp", 01777), "chmod tmp")
	fatalIfError(os.Remove("/setup"), "removing setup")
	f, err := os.Open("/")
	fatalIfError(err, "opening root")
	fis, err := f.Readdir(0)
	fatalIfError(err, "reading root")
	for _, fi := range fis {
		fmt.Printf("%s: %v\n", fi.Name(), fi.Mode())
	}
}
