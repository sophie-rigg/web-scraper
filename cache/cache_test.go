package cache

import (
	"sync"
	"testing"
)

func TestClient_Set(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{
			name: "real test",
			key:  "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Map: sync.Map{},
			}
			c.Set(tt.key)
			got := c.Exists(tt.key)
			if !got {
				t.Errorf("Expected key to exist in cache")
			}
			allResults := c.GetAll()
			if allResults != tt.key+"\n" {
				t.Errorf("Expected %s got %s", tt.key, allResults)
			}
		})
	}
}
