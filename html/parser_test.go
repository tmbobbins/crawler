package html

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func htmlString() string {
	return `
<!DOCTYPE html>
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
		<a href=""></a>
		<a></a>
    </div>
</body>
</html>
`
}

func TestParserGetAnchorUrls(t *testing.T) {
	parser := NewParser(strings.NewReader(htmlString()))
	urls, err := parser.GetAnchorUrls()
	if err != nil {
		log.Fatal(err)
	}
	assert.ElementsMatch(t, []string{
		"/link_test.html",
		"https://google.com",
		"file://moo/moo.html",
	}, urls)
}
