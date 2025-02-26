package main

import (
	"fmt"
	"sync"
	"time"
)

const numGetThreads = 10
const numGetMessages = 10000
const numPostThreads = 10
const numPostMessages = 10000
const httpTimeOutSeconds = 5

func main() {

	client := NewGuesbookClient(8089, "localhost", time.Second*httpTimeOutSeconds)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go runGets(client, &wg)

	wg.Add(1)
	go runPosts(client, &wg)
	wg.Wait()
}

// spawn worker threads and send get requests
func runGets(client *GuestbookClient, wg *sync.WaitGroup) {
	defer wg.Done()

	getWg := sync.WaitGroup{}
	inputCh := make(chan struct{})
	replyCh := make(chan int, 1000)
	for i := 0; i < numGetThreads; i++ {
		getWg.Add(1)
		go runGet(client, inputCh, replyCh, &getWg)
	}

	getWg.Add(1)
	go processGetResponses(replyCh, &getWg)

	for i := 1; i <= numGetMessages; i++ {
		inputCh <- struct{}{}
	}

	time.Sleep(2 * time.Second)

	close(inputCh)
	close(replyCh)
	getWg.Wait()
}

// GET worker, calls guestbook client
func runGet(client *GuestbookClient, inputCh chan struct{}, replyCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for range inputCh {
		client.DoGet(replyCh)
		count++
	}
}

// process responses from all get workers
func processGetResponses(replyCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	max := 0
	for i := range replyCh {
		if i > max {
			max = i
		}
		fmt.Printf("Total expected visits: %d, Visits got: %d\n", numGetMessages, max)
	}
}

// spawn worker threads and send post requests
func runPosts(client *GuestbookClient, wg *sync.WaitGroup) {
	defer wg.Done()

	postWg := sync.WaitGroup{}
	inputCh := make(chan int)

	for i := 0; i < numPostThreads; i++ {
		postWg.Add(1)
		go runPost(client, inputCh, &postWg)
	}

	for i := 1; i <= numPostMessages; i++ {
		inputCh <- i
	}

	time.Sleep(2 * time.Second)

	close(inputCh)
	postWg.Wait()
}

// post worker
func runPost(client *GuestbookClient, inputCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inputCh {
		client.DoPost(Guest{
			Name:    fmt.Sprintf("abc-%d", i),
			Special: "true",
		})
	}
}
