package node

import (
	"os"
	"sync"
	"testing"

	"golang.org/x/net/html"
)

func TestClient_Process(t *testing.T) {
	file, err := os.Open("test-data/monzo.html")
	if err != nil {
		t.Fatal(err)
	}
	node, err := html.Parse(file)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		domain string
	}
	type args struct {
		node       *html.Node
		urlChannel chan string
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "real test monzo",
			fields: fields{
				domain: "monzo.com",
			},
			args: args{
				node:       node,
				urlChannel: make(chan string, 100000),
				wg:         &sync.WaitGroup{},
			},
			want: 67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				domain: tt.fields.domain,
			}
			c.Process(tt.args.node, tt.args.urlChannel, tt.args.wg)
			go func() {
				tt.args.wg.Wait()
				close(tt.args.urlChannel)
			}()
			var result []string
			for url := range tt.args.urlChannel {
				result = append(result, url)
			}

			if len(result) != tt.want {
				t.Errorf("got \n%v\n want \n%v", len(result), tt.want)
			}
		})
	}
}
