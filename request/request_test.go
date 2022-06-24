package request

import (
	"crawler/tests"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func rawHtml() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <base href="http://localhost/" target="_blank">
    <title>Title</title>
</head>
<body>
    <div>
        <img src="/moo.png"  alt="cow image"/>
        <a href="/link_test.html"></a>
        <a href="https://google.com"></a>
        <a href="file://moo/moo.html"></a>
    </div>
</body>
</html>
`
}
func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestRequestGetContent(t *testing.T) {
	request := NewRequest(tests.FileSystemClient(), tests.RelativeFile("test_data/index.html"))
	content, err := request.GetContent()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, rawHtml(), string(content))
}

func TestRequestFailing(t *testing.T) {
	request := NewRequest(http.Client{}, tests.RelativeFile("test_data/moo.html"))
	_, err := request.GetContent()
	assert.True(t, err != nil)
}
