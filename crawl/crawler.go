package crawl

import (
	"bytes"
	"crawler/html"
	"crawler/request"
	"net/http"
	"strings"
	"sync"
)

type Crawler struct {
	client         http.Client
	rootUrl        string
	validSchemes   []string
	alreadyCrawled CrawlerVisited
}

func NewCrawler(client http.Client, rootUrl string) Crawler {
	return Crawler{
		client,
		rootUrl,
		[]string{"http", "https"},
		NewCrawlerVisited(),
	}
}

func (c *Crawler) Crawl() ([]*Url, error) {
	rootUrls := []*Url{NewUrl(c.rootUrl)}
	return c.crawlUrls(rootUrls)
}

func (c *Crawler) crawlUrls(urls []*Url) ([]*Url, error) {
	var waitGroup sync.WaitGroup
	for _, url := range urls {
		waitGroup.Add(1)
		url := url
		go func() {
			defer waitGroup.Done()
			_ = c.handleURL(url) //change to returning errors through channels rather than suppressing
		}()
	}

	waitGroup.Wait()

	return urls, nil
}

func (c *Crawler) handleURL(url *Url) error {
	if c.alreadyCrawled.hasVisited(url.GetString()) {
		return nil
	}

	childUrls, err := c.crawlUrl(url.GetString())
	if err != nil {
		return err
	}

	for _, childUrl := range childUrls {
		url.AddChild(childUrl)
	}

	_, err = c.crawlUrls(url.GetChildren())
	if err != nil {
		return err
	}

	return nil
}

func (c *Crawler) crawlUrl(url string) ([]string, error) {
	req := request.NewRequest(c.client, url)
	content, err := req.GetContent()
	if err != nil {
		return nil, err
	}

	return c.parse(content)
}

func isAbsoluteUrl(url string, rootUrl string) bool {
	return strings.HasPrefix(url, rootUrl)
}

func isRelativeUrl(url string) bool {
	return strings.HasPrefix(url, "/")
}

func (c *Crawler) filterUrls(urls []string) []string {
	var filtered []string
	for _, url := range urls {
		if isAbsoluteUrl(url, c.rootUrl) {
			filtered = append(filtered, url)
			continue
		}

		if isRelativeUrl(url) {
			filtered = append(filtered, c.rootUrl+url)
			continue
		}
	}

	return filtered
}

func (c *Crawler) parse(content []byte) ([]string, error) {
	parser := html.NewParser(bytes.NewReader(content))
	urls, err := parser.GetAnchorUrls()
	if err != nil {
		return nil, err
	}

	return c.filterUrls(urls), nil
}
