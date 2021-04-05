package main

import (
	"fmt"
	"math/rand"
	"time"
)

func crawlData(url string, channel chan string) {
	fmt.Println("Running " + url)
	randomTime := int(rand.Intn(5) + 2)
	time.Sleep(time.Second * time.Duration(randomTime))
	channel <- "Data get from " + url
}

func runProcess(tasks chan string, responses chan string) {
	for value := range tasks {
		crawlData(value, responses)
	}
}

func limitCrawl(limit int, listUrl []string) {
	lengthUrl := len(listUrl)
	if limit > lengthUrl {
		limit = lengthUrl
	}
	tasks := make(chan string, lengthUrl)
	responses := make(chan string, lengthUrl)

	for _, url := range listUrl {
		tasks <- url
	}

	for i := 0; i < limit; i++ {
		go runProcess(tasks, responses)
	}

	for i := 0; i < lengthUrl; i++ {
		fmt.Println(<-responses)
	}
}

func main() {
	listUrl := []string{
		"google.com",
		"facebook.com",
		"reddit.com",
		"slack.com",
		"trello.com",
		"24h.com.vn",
		"github.com",
		"golang.org",
		"data.com",
		"aws.com",
		"vnlp.ai",
	}

	limitCrawl(6, listUrl)
	fmt.Println("OK")
}
