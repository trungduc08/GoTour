package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	quit := make(chan bool)
	visitedUrls := map[string]bool{url: false}
	go doCrawl(url, depth, fetcher, quit, visitedUrls)
	<-quit
}
func doCrawl(url string, depth int, fetcher Fetcher, quit chan bool, visitedUrls map[string]bool) {
	if depth < 0 {
		quit <- true
		return
	}
	visited := visitedUrls[url]
	if visited {
		quit <- true
		return
	} else {
		visitedUrls[url] = true
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}
	fmt.Printf("found %v %v\n", url, body)
	childrenQuit := make(chan bool)
	for _, u := range urls {
		go doCrawl(u, depth-1, fetcher, childrenQuit, visitedUrls)
		<-childrenQuit
	}
	quit <- true
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// func main() {
// 	Crawl("https://golang.org/", 4, fetcher)
// }
