package main

import (
	"fmt"
	"sync"
)

var URLs = make(map[string]string)

type SafeUrls struct {
	mu sync.Mutex 
	v map[string]string
}

func (c *SafeUrls) LockURLs(key string, value string)  {
	c.mu.Lock()
	c.v[key]= value
	c.mu.Unlock()
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *SafeUrls) Crawl(wg *sync.WaitGroup, url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if _, ok := c.v[url]; ok {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		c.LockURLs(url, err.Error())
		return
	}
	c.LockURLs(url, fmt.Sprintf("found: %s %q\n", url, body))
	for _, u := range urls {
		wg.Add(1)
		c.Crawl(wg, u, depth-1, fetcher)
	}
	return
}

func main() {
	var wg sync.WaitGroup
	c := SafeUrls{v: make(map[string]string)}
	wg.Add(1)
	go c.Crawl(&wg, "https://golang.org/", 100, fetcher)
	wg.Add(1)
	go c.Crawl(&wg, "https://golang.org/", 100, fetcher)

	wg.Wait()
	for _,u := range c.v {
		fmt.Println(u)	
	}
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