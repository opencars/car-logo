package carlogos

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/antchfx/htmlquery"
)

const (
	DefaultBaseURL = "https://www.carlogos.org"
	Xpath          = "/html/body/div/div/dl/dd/a/img"
)

var (
	BaseURL = DefaultBaseURL
	Pages   = 8
)

// Client represents wrapper on http.Client{} for scrapping www.carlogos.org website.
type Client struct {
	http.Client
}

// NewClient creates new instance of Client based on http.Client{}.
func NewClient() *Client {
	return &Client{
		Client: *http.DefaultClient,
	}
}

// ScrapeEmblems scrapes emblems from 8 HTML pages on the www.carlogos.org website.
func (c *Client) ScrapeEmblems(base string) error {
	for i := 1; i < Pages; i++ {
		url := fmt.Sprintf("%s/Car-Logos/list_1_%d.html", BaseURL, i)

		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			return err
		}

		nodes, err := htmlquery.QueryAll(doc, Xpath)
		if err != nil {
			return err
		}

		for _, node := range nodes {
			emblem := htmlquery.SelectAttr(node, "src")

			parts := strings.Split(emblem, "/")
			name := parts[len(parts)-1]
			name = strings.ReplaceAll(name, "-logo", "")

			imgPath := path.Join(base, strings.ToLower(name))
			f, err := os.Create(imgPath)
			if err != nil {
				return err
			}

			log.Printf("Loading %s into %s...\n", emblem, imgPath)
			res, err := http.Get(emblem)
			if err != nil {
				return err
			}

			_, err = io.Copy(f, res.Body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
