package main

import (
	"./filters"
	"fmt"
	"log"
)

func main() {
	ffFilters := filters.NewFFMpegFilters(100, 1)

	testFilter := filters.Filter{
		Name: "asetrate",
		Values: map[string]string{
			"r": "50000",
		},
	}

	success, msg := ffFilters.SetFilter(testFilter)
	if !success {
		log.Println(msg)
	}

	output := ffFilters.ToString()
	fmt.Println(output)
}
