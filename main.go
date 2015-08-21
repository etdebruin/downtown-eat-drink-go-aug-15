package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/etdebruin/hexdump/gohexdump"
)

func main() {

	// read file from command line
	f, err := os.Open("main.go")
	if err != nil {
		log.Println(err)
	}

	r := gohexdump.Dump(f)

	ioutil.ReadAll(r)

	//d, err := ioutil.ReadAll(f)
	//if err != nil {
	//log.Println(err)
	//}

	// pass []bytes to gohexdump

	// output

	// offer terminal editing
}
