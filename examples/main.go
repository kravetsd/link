package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kravetsd/link"
)

func main() {

	fl, err := os.Open("ex2.html")
	if err != nil {
		log.Println("Erro openning file: ", err)
	}

	// ln, err := link.Parse(fl)
	// if err != nil {
	// 	log.Println("lnk parser error: ", err)
	// }

	// fmt.Printf("%+v", ln)

	links, err := link.Parse(fl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", links)

}
