package crawl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlAddChild(t *testing.T) {
	url := NewUrl("example.com")
	url.AddChild("example.com/path")

	assert.Equal(t, "example.com", url.GetString())
	assert.Equal(t, "example.com/path", url.GetChildren()[0].GetString())
}

func TestUrlAddChildren(t *testing.T) {
	url := NewUrl("example.com")
	url.AddChildren([]string{"example.com/path", "example.com/path2"})

	assert.Equal(t, "example.com", url.GetString())
	assert.Equal(t, "example.com/path", url.GetChildren()[0].GetString())
	assert.Equal(t, "example.com/path2", url.GetChildren()[1].GetString())
}

func TestOutput(t *testing.T) {
	var urls []*Url
	url := NewUrl("example.com")
	url.AddChildren([]string{"example.com/path", "example.com/path2"})
	urls = append(urls, url)
	output := Output(urls, 0)
	assert.Equal(t,
		`-> example.com
--> example.com/path
--> example.com/path2
`,
		output,
	)
}
