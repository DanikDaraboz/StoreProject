// cache_test.go
package test

import (
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
)

func TestNewTemplateCache(t *testing.T) {
	cache, err := handlers.NewTemplateCache()
	if err != nil {
		t.Fatalf("failed to create template cache: %v", err)
	}

	for name := range cache {
		t.Logf("Loaded template: %s", name)
	}
}
