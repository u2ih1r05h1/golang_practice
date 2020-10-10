package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := make(map[string]int)
	var mutex sync.Mutex
	var crawl_sub func(string, int)
	crawl_sub = func(url string, depth int) {
		if depth <= 0 {
			fmt.Printf("too deep: %s\n", url)
			return
		}
		// Don't fetch the same URL twice.
		mutex.Lock()
		cache[url]++
		c := cache[url]
		mutex.Unlock()
		if c > 1 {
			fmt.Printf("fetched: %s\n", url)
			return
		}
		// Fetch URLs in parallel.
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go crawl_sub(u, depth-1)
		}
	}
	crawl_sub(url, depth)
	time.Sleep(time.Second) // goルーチンの完了待ち. chanを使って待ち合わせたほうが良い.
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
