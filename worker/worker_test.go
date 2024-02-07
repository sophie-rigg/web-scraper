package worker

import (
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sophie-rigg/web-scraper/cache"
	"github.com/sophie-rigg/web-scraper/mocks"
	"golang.org/x/net/html"
)

func TestClient_Do(t *testing.T) {
	type fields struct {
		processor func(p *mocks.MockProcessor)
		reader    func(r *mocks.MockClient)
		domain    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test Do",
			fields: fields{
				processor: func(p *mocks.MockProcessor) {
					p.EXPECT().Process(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(node *html.Node, urlChannel chan string, wg *sync.WaitGroup) {
						defer wg.Done()
						wg.Add(1)
						urlChannel <- "http://example.com/test"
					}).Times(2)
				},
				reader: func(r *mocks.MockClient) {
					r.EXPECT().Read("http://example.com").Return(&html.Node{}, nil)
					r.EXPECT().Read("http://example.com/test").Return(&html.Node{}, nil)
				},
				domain: "http://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			p := mocks.NewMockProcessor(ctrl)
			r := mocks.NewMockClient(ctrl)
			tt.fields.processor(p)
			tt.fields.reader(r)

			c := &Client{
				processor: p,
				reader:    r,
				domain:    tt.fields.domain,
				cache:     cache.Client{},
			}
			c.Do()
		})
	}
}
