package crawl

import "fmt"

type Url struct {
	url       string
	childUrls []*Url
}

func NewUrl(url string) *Url {
	return &Url{url, []*Url{}}
}

func (u *Url) AddChildren(urls []string) {
	for _, url := range urls {
		u.childUrls = append(u.childUrls, NewUrl(url))
	}
}

func (u *Url) AddChild(url string) {
	u.childUrls = append(u.childUrls, NewUrl(url))
}

func (u *Url) GetString() string {
	return u.url
}

func (u *Url) GetChildren() []*Url {
	return u.childUrls
}

func Output(urls []*Url, indent int) string {
	output := ""
	for _, url := range urls {
		for i := 0; i <= indent; i++ {
			output += fmt.Sprintf("-")
		}
		output += fmt.Sprintf("> %s\n%s", url.GetString(), Output(url.GetChildren(), indent+1))
	}

	return output
}
