package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch返回URL的body内容，并且将在这个页面上找到的URL放到一个slice中
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	mux     sync.Mutex
	hashMap map[string]string
}

func (m *SafeMap) Set(key string, body string) {
	m.mux.Lock()
	m.hashMap[key] = body
	m.mux.Unlock()
}

func (m *SafeMap) Value(key string) (string, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	val, ok := m.hashMap[key]
	return val, ok
}

// Crawl使用 fetcher 从某个URL开始递归地爬取页面，直到达到最大深度
func Crawl(url string, depth int, fetcher Fetcher, urlMap SafeMap) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	urlMap.Set(url, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		if _, ok := urlMap.Value(u); !ok {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, urlMap)
		}
	}
	return
}

var wg sync.WaitGroup

func main() {
	urlMap := SafeMap{hashMap: make(map[string]string)}

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, urlMap)
	wg.Wait()

	for url := range urlMap.hashMap {
		body, _ := urlMap.Value(url)
		fmt.Printf("found: %s %q\n", url, body)
	}
}

// fakeFetcher 是返回若干结果的Fetcher
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

// fetcher是填充完的fakeFetcher
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
			"https://golang.org/pkg/fmt",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt": &fakeResult{
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
