package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/cryring/curlx"
)

var (
	method   = flag.String("m", "GET", "http method")
	url      = flag.String("u", "", "http url")
	body     = flag.String("d", "", "http body data")
	colorize = flag.Bool("c", true, "display http response body with color")
	verbose  = flag.Bool("v", false, "display verbose of http response")
)

func main() {
	flag.Parse()

	if *method == "" {
		fmt.Println("empty http method")
		return
	}

	if *url == "" {
		fmt.Println("empty http url")
		return
	}

	c := curlx.NewClient(3*time.Second, *colorize, *verbose)
	switch *method {
	case "GET", "get":
		if err := c.Get(*url, *body); err != nil {
			fmt.Println(err)
		}
	case "POST", "post":
		if err := c.Post(*url, *body); err != nil {
			fmt.Println(err)
		}
	case "DELETE", "delete":
		if err := c.Delete(*url, *body); err != nil {
			fmt.Println(err)
		}
	}
}
