package reader

import (
	"golang.org/x/net/html"
)

//go:generate mockgen -destination=../mocks/mock_reader.go -package=mocks --source=reader.go
type Client interface {
	Read(url string) (*html.Node, error)
}
