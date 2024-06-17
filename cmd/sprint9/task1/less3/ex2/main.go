package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

type Client struct {
	client      *http.Client
	rateLimiter *rate.Limiter
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.rateLimiter.Allow() {
		return c.client.Do(req)
	}
	return nil, errors.New("request not allowed")
}

func main() {
	client := &Client{
		client:      http.DefaultClient,
		rateLimiter: rate.NewLimiter(rate.Limit(1), 2),
	}

	// здесь, например, API биржевых котировок
	URL := "https://iss.moex.com/iss/statistics/engines/futures/markets/indicativerates/securities.xml"
	req, _ := http.NewRequest("GET", URL, nil)
	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// логика обработки результата запроса
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 20; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
