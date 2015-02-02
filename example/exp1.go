package main

import (
	"github.com/wtnb75/go-wordexp"
	"log"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		res, err := wordexp.WordExp(v, 0)
		log.Println(v, err, res)
	}
}
