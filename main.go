package main

import (
	"fmt"
	"time"
)

func crawlData(url string, channel chan string) {
	time.Sleep(time.Second * 2)
	channel <- "Data get from " + url
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

	channel := make(chan string, 5)
	index := 0
	for ; index < 5; index++ {
		go crawlData(listUrl[index], channel)
	}

	lengthUrl := len(listUrl)
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
	fmt.Println("OK")
}
