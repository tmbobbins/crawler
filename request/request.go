package request

import (
	"io"
	"net/http"
)

type Request struct {
	client http.Client
	url    string
}

func NewRequest(client http.Client, url string) Request {
	return Request{client, url}
}

func (r Request) GetContent() ([]byte, error) {
	response, err := r.client.Get(r.url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}
