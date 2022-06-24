package tests

import (
	"crawler/crawl"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func fileCrawler() crawl.Crawler {
	return crawl.NewCrawler(
		FileSystemClient(),
		RelativeFile("crawler_test_data"),
	)
}

func TestCrawl(t *testing.T) {
	crawler := fileCrawler()
	links, err := crawler.Crawl()
	if err != nil {
		log.Fatal(err)
	}

	assert.Len(t, links, 1)
	assert.Equal(t, RelativeFile("crawler_test_data"), links[0].GetString())

	children := links[0].GetChildren()
	assert.Len(t, children, 1)
	assert.Equal(t, RelativeFile("crawler_test_data/link_test.html"), children[0].GetString())

	subChildren := children[0].GetChildren()
	assert.Len(t, subChildren, 2)
	assert.Equal(t, RelativeFile("crawler_test_data/test.html"), subChildren[0].GetString())
	assert.Len(t, subChildren[0].GetChildren(), 0)
	assert.Equal(t, RelativeFile("crawler_test_data/blink_test.html"), subChildren[1].GetString())
	assert.Len(t, subChildren[1].GetChildren(), 0)
}

func TestCrawlOutput(t *testing.T) {
	crawler := fileCrawler()
	links, err := crawler.Crawl()
	if err != nil {
		log.Fatal(err)
	}

	output := crawl.Output(links, 0)
	assert.Equal(t,
		`-> file:///var/www/crawler/tests/crawler_test_data
--> file:///var/www/crawler/tests/crawler_test_data/link_test.html
---> file:///var/www/crawler/tests/crawler_test_data/test.html
---> file:///var/www/crawler/tests/crawler_test_data/blink_test.html
`,
		output,
	)
}
