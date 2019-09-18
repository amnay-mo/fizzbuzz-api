package stats

import "testing"

const testingRedisAddr = "localhost:6379"

// the returned function flushes the redis key
// and closes the redis client
func newTestingRedisStore(t *testing.T) (*RedisStore, func()) {
	t.Helper()
	store := NewRedisStatsStore(testingRedisAddr)
	cmd := store.client.Ping()
	if cmd.Err() != nil {
		t.Skipf("can't reach testing redis server")
	}
	return store, func() {
		store.Flush()
		store.client.Close()
	}
}

func TestRedisStore(t *testing.T) {
	t.Run("can store hits and retrieve max", func(t *testing.T) {
		store, cl := newTestingRedisStore(t)
		defer cl()
		store.Increment("car=audi")
		store.Increment("bike=peugeot")
		store.Increment("bike=peugeot")
		hits, err := store.GetMax()
		if err != nil {
			t.Errorf("got unexpected error: %v", err)
			return
		}
		if hits.Parameters != "bike=peugeot" || hits.Count != 2 {
			t.Errorf("got unexpected max hits")
		}
	})
}
