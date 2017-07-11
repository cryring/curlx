package curlx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Client is http request client
type Client struct {
	timeout  time.Duration
	formater *Formater
	verbose  bool
}

// NewClient is used to create http client
func NewClient(timeout time.Duration, colorize, verbose bool) *Client {
	c := &Client{
		timeout:  timeout,
		formater: NewFormater(colorize),
		verbose:  verbose,
	}
	return c
}

func (c *Client) printResponseHeader(resp *http.Response) {
	fmt.Printf("%v %v\n", resp.Proto, resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%v: ", k)
		for _, n := range v {
			fmt.Printf("%v", n)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Get is used to send http get request
func (c *Client) Get(url, body string) error {
	req := gorequest.New().Timeout(c.timeout)
	resp, body, errs := req.Get(url).Send(body).End()
	if errs != nil {
		fmt.Println(errs)
		return fmt.Errorf("http get request error")
	}
	defer resp.Body.Close()

	if c.verbose {
		c.printResponseHeader(resp)
	}

	result, err := c.formater.ExportJSON(body)
	if err != nil {
		fmt.Println(body)
		return nil
	}

	fmt.Println(result)
	return nil
}

// Post is used to send http post request
func (c *Client) Post(url, body string) error {
	req := gorequest.New().Timeout(c.timeout)
	resp, body, errs := req.Post(url).Send(body).End()
	if errs != nil {
		fmt.Println(errs)
		return fmt.Errorf("http post request error")
	}
	defer resp.Body.Close()

	if c.verbose {
		c.printResponseHeader(resp)
	}

	result, err := c.formater.ExportJSON(body)
	if err != nil {
		fmt.Println(body)
		return nil
	}

	fmt.Println(result)
	return nil
}

// Delete is used to send http delete request
func (c *Client) Delete(url, body string) error {
	req := gorequest.New().Timeout(c.timeout)
	resp, body, errs := req.Delete(url).Send(body).End()
	if errs != nil {
		fmt.Println(errs)
		return fmt.Errorf("http delete request error")
	}
	defer resp.Body.Close()

	if c.verbose {
		c.printResponseHeader(resp)
	}

	result, err := c.formater.ExportJSON(body)
	if err != nil {
		fmt.Println(body)
		return nil
	}

	fmt.Println(result)
	return nil
}
