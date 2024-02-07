package node

import (
	"sync"

	"golang.org/x/net/html"
)

type Client struct {
	domain string
}

func New(domain string) *Client {
	return &Client{domain: domain}
}

func (c *Client) Process(node *html.Node, urlChannel chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	// should only ever activate on first loop all other child go routines should not be nil
	// safety feature more than anything else
	if node == nil {
		return
	}
	go func() {
		defer wg.Done()
		// only looks for nodes that may contain links
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				// attribute key that contains the URL
				if attr.Key == "href" {
					urlChannel <- attr.Val
				}
			}
		}

		// recursively process the children
		child := node.FirstChild

		// will go through all the siblings launching a go routine for each one
		for child != nil {
			c.Process(child, urlChannel, wg)
			// move to the next sibling
			child = child.NextSibling
		}
	}()
}
