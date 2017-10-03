package main

import (
	"flag"
	"log"
	"net/url"
	"fmt"
	"strings"
)

func main() {

	passed_url := flag.String("url", "", "Platform URL")
	flag.Parse()

	if *passed_url == "" {
		log.Fatal("Need a URL")
	}
	url_struct, err := url.Parse(*passed_url)
	if err != nil {
		log.Fatal("Unable to parse URL")
	}
	fmt.Println("\nBase URL:\t", url_struct.Host)
	pth, err := url.PathUnescape(url_struct.Path)
	if err != nil {
		log.Fatal("Unable to unescape the Path")
	}
	fmt.Println("Path:\t\t", pth)
	if url_struct.Fragment != "" {
		fmt.Println("Fragment:\t", url_struct.Fragment)
	}
	qs, err := url.ParseQuery(url_struct.RawQuery)
	if err != nil {
		log.Fatal("Unable to unescape the Query")
	}
	fmt.Println("Queries")
	for key, val := range qs {
		if strings.Contains(val[0], ","){
			for idx, itm := range (strings.Split(val[0], ",")) {
				if idx == 0 {
					fmt.Println(key, "\t", itm)
				} else {
					fmt.Println("\t\t", itm)
				}
			}
		} else {
			fmt.Println(key, "\t", val[0])
		}
	}
}
