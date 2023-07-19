//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetChan chan<- *Tweet, quit chan<- struct{}) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			quit <- struct{}{}
			return
		}

		tweetChan <- tweet
	}
}

func consumer(tweetChan <-chan *Tweet, quit <-chan struct{}, done chan<- bool) {
	// for _, t := range tweets {
	// 	if t.IsTalkingAboutGo() {
	// 		fmt.Println(t.Username, "\ttweets about golang")
	// 	} else {
	// 		fmt.Println(t.Username, "\tdoes not tweet about golang")
	// 	}
	// }

	for {
		select {
		case tweet := <-tweetChan:
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets about golang")
			} else {
				fmt.Println(tweet.Username, "\tdoes not tweet about golang")
			}
		case <-quit:
			// all consumed
			done <- true
			return
		}
	}

}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweetChan := make(chan *Tweet, 1)
	quit := make(chan struct{}, 1)
	done := make(chan bool, 1)

	// Producer
	go producer(stream, tweetChan, quit)

	// Consumer
	go consumer(tweetChan, quit, done)
	<-done
	fmt.Printf("Process took %s\n", time.Since(start))
}
