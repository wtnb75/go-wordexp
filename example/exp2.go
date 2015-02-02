package main

import (
	"github.com/wtnb75/go-wordexp"
	"log"
)

func main() {
	// prints "YYYY/MM/DD hh:mm:ss fnmatch: String does not match"
	if err := wordexp.FnMatch("pattern", "target", 0); err != nil {
		log.Println("fnmatch:", err)
	}
}
