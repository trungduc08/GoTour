package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}
type fakeResult struct {
	body string
	urls []string
}
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %v", url)
}
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	fmt.Println("first", depth)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {

		Crawl(u, depth-1, fetcher)

	}
	return

}

// func main() {
// 	fetcher := fakeFetcher{
// 		"google.com.org": &fakeResult{
// 			"trang chu gg",
// 			[]string{"google.com.org/pkg", "google.com.org/cmd"},
// 		},

// 		"google.com.org/pkg": &fakeResult{
// 			"trang chu fb",
// 			[]string{"google.com.org", "fb2", "fb3"},
// 		},
// 	}

// 	// fmt.Println(fetcher.Fetch("google.com.org"))

// 	Crawl("google.com.org", 4, fetcher)

// }
