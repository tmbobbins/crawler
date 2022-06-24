package crawl

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFilterUrls(t *testing.T) {
	crawler := NewCrawler(http.Client{}, "example.com")
	urls := crawler.filterUrls(
		[]string{
			"example.com",
			"example.com/path",
			"example.example.com",
			"exmaple.example.com/path",
			"google.com",
			"/newPath",
		},
	)

	assert.ElementsMatch(t, []string{
		"example.com",
		"example.com/path",
		"example.com/newPath",
	}, urls)
}
