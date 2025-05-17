package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var visited = map[string]bool{}

const maxDepth = 5

// crawl funciton fetches the url, parses html content, and visits any detected urls, recursively crawling
func crawl(rawURL string, depth int) {
	if depth > maxDepth {
		return
	}

	if visited[rawURL] {
		return
	}

	visited[rawURL] = true

	//fetch the page
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Println("error: ", err)
	}

	defer resp.Body.Close()

	//if there is no html text to parse, return
	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return
	}

	pageURL, _ := url.Parse(rawURL)
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()

		//if the token is a <a> tag
		if token.Type == html.StartTagToken && token.Data == "a" {
			//loop through tag attributes to find href
			for _, attr := range token.Attr {
				if attr.Key == "href" {

					//skip empty href
					if attr.Val == "" {
						continue
					}

					//skip fragment-only urls
					if strings.HasPrefix(attr.Val, "#") {
						continue
					}

					//we have found a link, parse and visit it
					link, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}

					//fmt.Println("found link: ", link)
					//crawl the link we found
					absURL := pageURL.ResolveReference(link)
					crawl(absURL.String(), depth+1)
				}
			}
		}
	}

	fmt.Println("Visited ", rawURL)
}

func main() {
	start := "https://www.scrapingcourse.com/ecommerce/"
	crawl(start, 0)
}
