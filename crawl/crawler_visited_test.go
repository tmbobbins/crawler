package crawl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrawlerVisitedHasVisited(t *testing.T) {
	crawlerVisited := NewCrawlerVisited()
	assert.False(t, crawlerVisited.hasVisited("example.com"))
	assert.True(t, crawlerVisited.hasVisited("example.com"))

	assert.False(t, crawlerVisited.hasVisited("example.com/path"))
	assert.True(t, crawlerVisited.hasVisited("example.com/path"))
	assert.True(t, crawlerVisited.hasVisited("example.com"))
	assert.True(t, crawlerVisited.hasVisited("example.com/path"))
}
