package html

import (
	"golang.org/x/net/html"
	"io"
)

type Parser struct {
	reader io.Reader
}

func NewParser(reader io.Reader) Parser {
	return Parser{reader}
}

func getTagByType(tagType string, node *html.Node) []*html.Node {
	var tags []*html.Node
	if node.Type == html.ElementNode && node.Data == tagType {
		tags = append(tags, node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childTags := getTagByType(tagType, child)
		tags = append(tags, childTags...)
	}

	return tags
}

func getHrefFromTag(tag *html.Node) string {
	for _, attribute := range tag.Attr {
		if attribute.Key == "href" {
			return attribute.Val
		}
	}

	return ""
}

func getHrefFromTags(tags []*html.Node) []string {
	var urls []string
	for _, anchor := range tags {
		url := getHrefFromTag(anchor)
		if url == "" {
			continue
		}

		urls = append(urls, url)
	}

	return urls
}

func (p Parser) GetAnchorUrls() ([]string, error) {
	htmlDoc, err := html.Parse(p.reader)
	if err != nil {
		return nil, err
	}

	return getHrefFromTags(getTagByType("a", htmlDoc)), nil
}
