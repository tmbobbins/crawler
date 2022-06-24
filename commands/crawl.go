package commands

import (
	"crawler/colours"
	"crawler/crawl"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

var crawlCommand = &cobra.Command{
	Use:       "crawl URL",
	Short:     "crawls a website and lists out the sitemap",
	Long:      "crawls a website and lists out the sitemap, accepts a ",
	ValidArgs: []string{"URL"},
	Args:      validateCrawl,
	Run:       crawlExec,
}

func validateCrawl(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a url argument")
	}

	_, err := url.ParseRequestURI(args[0])
	if err != nil {
		return err
	}

	return nil
}

func crawlExec(cmd *cobra.Command, args []string) {
	crawler := crawl.NewCrawler(http.Client{}, args[0])
	urls, err := crawler.Crawl()
	if err != nil {
		colours.PrintRed(err)
		return
	}

	fmt.Println(crawl.Output(urls, 0))
}
