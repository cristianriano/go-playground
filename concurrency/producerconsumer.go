//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package concurrency

import (
	"fmt"
	"time"
)

// func producer(stream Stream) (tweets []*Tweet) {
// 	for {
// 		tweet, err := stream.Next()
// 		if err == ErrEOF {
// 			return tweets
// 		}

// 		tweets = append(tweets, tweet)
// 	}
// }

// func consumer(tweets []*Tweet) {
// 	for _, t := range tweets {
// 		if t.IsTalkingAboutGo() {
// 			fmt.Println(t.Username, "\ttweets about golang")
// 		} else {
// 			fmt.Println(t.Username, "\tdoes not tweet about golang")
// 		}
// 	}
// }

func producer(stream Stream) chan(*Tweet) {
	ch := make(chan *Tweet)

	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				close(ch)
				return
			}

			ch <- tweet
		}
	}()

	return ch
}

func consumer(ch <-chan(*Tweet)) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

// ProducerExample runs excercise 1 from loong/go-concurrency-excercises
func ProducerExample() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweetsChannel := producer(stream)

	// Consumer
	consumer(tweetsChannel)

	fmt.Printf("Process took %s\n", time.Since(start))
}
