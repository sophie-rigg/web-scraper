package worker

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/sophie-rigg/web-scrapper/cache"
	"github.com/sophie-rigg/web-scrapper/processor"
	"github.com/sophie-rigg/web-scrapper/reader"
	"github.com/sophie-rigg/web-scrapper/utils"
	"golang.org/x/exp/maps"
)

type Client struct {
	processor processor.Processor
	reader    reader.Client
	domain    string
	cache     cache.Client
}

func New(p processor.Processor, r reader.Client, domain string) *Client {
	return &Client{
		processor: p,
		reader:    r,
		domain:    domain,
	}
}

func (c *Client) Do() {
	/*
		This is the wait group for the entire system
		All go routines launched by the worker will be waited on by this wait group
		Once all the go routines are done, the main go routine will be done
	*/
	var wg sync.WaitGroup
	wg.Add(1)
	c.processUrl(c.domain, &wg)
}

func (c *Client) processUrl(baseUrl string, wg *sync.WaitGroup) {
	// all urls contained within the baseUrl will be sent to this channel
	channel := make(chan string, 10)

	// wait group for a singular URL since each url launches its own channel
	// each routine must have a separate wait group for closing its own channel
	var urlWG sync.WaitGroup

	page, err := c.reader.Read(baseUrl)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read")
	}

	// gets all the urls from the page and sends them to the channel then calls wg.Done()
	c.processor.Process(page, channel, &urlWG)

	go func() {
		urlWG.Wait()
		close(channel)
	}()

	results := make(map[string]struct{})
	for url := range channel {
		if _, ok := results[url]; ok {
			continue
		}
		// checks if the url is appropriate for the domain and if it is not in the cache
		if nextUrl, ok := utils.CheckDomainMatch(c.domain, url); ok && !c.cache.Exists(nextUrl) {
			// once a url is found it is added to the cache so another go routine does not process it
			c.cache.Set(nextUrl)
			wg.Add(1)
			go c.processUrl(nextUrl, wg)
		}
		results[url] = struct{}{}
	}
	logrus.WithFields(logrus.Fields{
		"url":      baseUrl,
		"sub_urls": maps.Keys(results),
	}).Info("results")
	wg.Done()
	wg.Wait()
}
