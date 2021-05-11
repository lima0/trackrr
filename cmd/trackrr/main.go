package main

import (
	"github.com/lima0/trackr/pkg/dhl"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) >= 11{
		log.Fatal("usage: trackrr 0000000000")
	}
	var IDList string
	for _, ID := range os.Args[1:] {
		IDList += ID + ","
	}
	res, err := dhl.GetResponse(IDList)
	if err != nil {
		log.Fatal(err)
	}
	dhl.PrintResponse(res)
}

