package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseType, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	cfg := NewConfig()
	cfg.DatabaseURL = databaseURL
	cfg.DatabaseType = databaseType

	store := New(cfg)

	if err := store.Open(); err != nil {
		t.Fatal(err)
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			_, err := store.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))

			if err != nil {
				t.Fatal(err)
			}
		}

		store.Close()
	}
}
