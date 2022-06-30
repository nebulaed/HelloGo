package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch返回URL的body内容，并且将在这个页面上找到的URL放到一个slice中
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	mux     sync.Mutex
	hashMap map[string]int
}

// Crawl使用 fetcher 从某个URL开始递归地爬取页面，直到达到最大深度
func (m *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	m.mux.Lock()
	if _, found := m.hashMap[url]; found {
		m.mux.Unlock()
		return
	}
	m.hashMap[url]++
	m.mux.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go m.Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	m := SafeMap{hashMap: make(map[string]int)}
	m.Crawl("https://golang.org/", 4, fetcher)
	// 这个时间大概要到毫秒级
	time.Sleep(time.Millisecond)
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
