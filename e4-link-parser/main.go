package main

import (
	"flag"
	"fmt"
	"os"

	"./linkParser"
)

func main() {
	file := flag.String("file", "ex1.html", "HTML file to parse the links from.")
	flag.Parse()
	fmt.Printf("Parsing %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	links, err := linkParser.Parse(f)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("%s -> %s", link.Href, link.Text)
	}
}
