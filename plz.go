package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {
	altName := flag.String("o", "default", "Output filename")

	flag.Parse()

	u := flag.Arg(0)

	url, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
	}

	fileName := path.Base(url.Path)

	var outFile *os.File

	if *altName == "default" {
		outFile, err = os.Create(fileName)
	} else {
		outFile, err = os.Create(*altName)
	}

	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}

