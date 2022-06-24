package crawl

import (
	"strings"
	"sync"
)

type CrawlerVisited struct {
	alreadyCrawled map[string]bool
	urlCounter     int64
	mutex          sync.Mutex
}

func NewCrawlerVisited() CrawlerVisited {
	return CrawlerVisited{make(map[string]bool), 0, sync.Mutex{}}
}

func (c *CrawlerVisited) hasVisited(url string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	url = strings.TrimSuffix(url, "/")
	_, ok := c.alreadyCrawled[url]
	if !ok {
		c.alreadyCrawled[url] = true
		c.urlCounter++
	}

	return ok
}
