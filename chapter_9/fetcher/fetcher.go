package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePagesize struct {
	URL  string
	Size int
}

func main() {

	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePagesize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePagesize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePagesize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page is", biggest.URL, "with", biggest.Size, "size")
}
