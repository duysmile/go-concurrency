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

func limitCrawl(limit int, listUrl []string) {
	lengthUrl := len(listUrl)
	if limit > lengthUrl {
		limit = lengthUrl
	}
	channel := make(chan string, limit)
	index := 0
	for ; index < limit; index++ {
		go crawlData(listUrl[index], channel)
	}

	for i := 0; i < lengthUrl; i++ {
		select {
		case response := <-channel:
			fmt.Println(response)
			if index < lengthUrl {
				go crawlData(listUrl[index], channel)
				index++
			}
		}
	}
}

func main() {
	listUrl := []string{
		"goolge.com",
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
