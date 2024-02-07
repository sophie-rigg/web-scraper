package processor

import (
	"sync"

	"golang.org/x/net/html"
)

//go:generate mockgen -destination=../mocks/mock_processor.go -package=mocks --source=processor.go
type Processor interface {
	// Process takes a html.Node pushes any URLs found to the urlChannel
	Process(node *html.Node, urlChannel chan string, wg *sync.WaitGroup)
}
