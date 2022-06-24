# Crawler
This is a very simple web crawler written in go to crawl through all the urls on a single domain.
utilises cobra and viper for CLI and go http and html libraries for requests / parsing

## Requirements
- Golang 1.18.*

## Usage
### Compile
`go build .`
### Run
`./crawler crawl URL`<br>
`./crawler crawl http://google.com`<br>
`./crawler crawl http://google.com > output.txt`

## Known issues

- downloads and parses file content, such as PDFs etc. 
  causing much slower performance on urls that contain direct links to large file types (need to fix)
- html parsing is synchronous
