package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		url := args[0]
		fmt.Println("URL", url)
		parsePage(url)
	} else {
		fmt.Println("no args")
	}
}

func parsePage(url string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Header.Get("Content-Type"))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
