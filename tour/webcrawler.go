package tour

import (
	"fmt"
	"sync"
	"time"
)

// RunWebCrawler runs the crawler example
func RunWebCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}

var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	wg.Add(1)
	// crawlRecursive(url, depth, fetcher)
	// crawlRecursiveUnique(url, depth, fetcher, &safeMap{cache: make(map[string] bool)})

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	crawlRecursiveTimed(url, depth, fetcher, ticker)
	wg.Wait()
}

// Fetcher implement an abstraction for web crawling
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type safeMap struct {
	cache map[string]bool
	mu sync.RWMutex
}

func crawlRecursiveTimed(url string, depth int, fetcher Fetcher, ticker *time.Ticker) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	<- ticker.C
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawlRecursiveTimed(u, depth-1, fetcher, ticker)
	}

	return
}

func crawlRecursive(url string, depth int, fetcher Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawlRecursive(u, depth-1, fetcher)
	}

	return
}

func crawlRecursiveUnique(url string, depth int, fetcher Fetcher, sm *safeMap) {
	defer func() {
		wg.Done()
		sm.add(url)
	}()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if sm.alreadyVisited(u) {
			continue
		}
		wg.Add(1)
		go crawlRecursiveUnique(u, depth-1, fetcher, sm)
	}

	return
}

func (sm *safeMap) alreadyVisited(url string) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return sm.cache[url]
}

func (sm *safeMap) add(url string) {
	sm.mu.Lock()
	sm.cache[url] = true
	sm.mu.Unlock()
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
