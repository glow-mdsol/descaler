package main

import (
	"flag"
	"log"
	"net/url"
	"fmt"
	"strings"
	"encoding/json"
	"os"
)

func toString(urlStruct *url.URL){
	fmt.Println("\nBase URL:\t", urlStruct.Host)
	pth, err := url.PathUnescape(urlStruct.Path)
	if err != nil {
		log.Fatal("Unable to unescape the Path")
	}
	fmt.Println("Path:\t\t", pth)
	if urlStruct.Fragment != "" {
		fmt.Println("Fragment:\t", urlStruct.Fragment)
	}
	qs, err := url.ParseQuery(urlStruct.RawQuery)
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

func toJSON(urlStruct *url.URL){
	type Descaled struct {
		BaseUrl string	`json:"base_url"`
		Path string `json:"path"`
		Fragment string `json:"fragment"`
		Queries map[string][]string `json:"queries"`
	}
	dsc := new(Descaled)
	dsc.BaseUrl = urlStruct.Host
	pth, err := url.PathUnescape(urlStruct.Path)
	if err != nil {
		log.Fatal("Unable to unescape the Path")
	}
	dsc.Path = pth
	if urlStruct.Fragment != ""{
		dsc.Fragment = urlStruct.Fragment
	}
	qs, err := url.ParseQuery(urlStruct.RawQuery)
	if err != nil {
		log.Fatal("Unable to unescape the Query")
	}
	for key, val := range qs {
		if strings.Contains(val[0], ",") {
			qs.Del(key)
			for _, itm := range (strings.Split(val[0], ",")){
				qs.Add(key, itm)
			}
		}
	}
	dsc.Queries = qs
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dsc)
}

func main() {

	passed_url := flag.String("url", "", "Platform URL")
	to_str := flag.Bool("str", false, "Output as a Tabbed String")
	flag.Parse()

	if *passed_url == "" {
		log.Fatal("Need a URL")
	}
	urlStruct, err := url.Parse(*passed_url)
	if err != nil {
		log.Fatal("Unable to parse URL")
	}
	if *to_str == true {
		toString(urlStruct)
	} else {
		toJSON(urlStruct)
	}
}
