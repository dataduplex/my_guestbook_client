package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var pattern = regexp.MustCompile(`Visits: ([0-9]+).*`)

type GuestbookClient struct {
	port    int
	host    string
	timeout time.Duration
}

func NewGuesbookClient(port int, host string, timeout time.Duration) *GuestbookClient {
	return &GuestbookClient{
		port:    port,
		host:    host,
		timeout: timeout,
	}
}

func (client *GuestbookClient) DoGet(replyCh chan int) {
	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://%s:%d/", client.host, client.port), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error doing get request, possible DEADLOCK: %s\n", err.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	respStr := string(body)
	matches := pattern.FindStringSubmatch(respStr)

	if len(matches) > 0 {
		visits, _ := strconv.Atoi(matches[1])
		replyCh <- visits
	}
}

type Guest struct {
	Name    string
	Special string
}

func (client *GuestbookClient) DoPost(guest Guest) {

	data := url.Values{}
	data.Set("name", guest.Name)
	data.Set("special", guest.Special)

	resp, err := http.PostForm(fmt.Sprintf("http://%s:%d/sign", client.host, client.port), data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("got error status code %d\n", resp.StatusCode)
		return
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
}
