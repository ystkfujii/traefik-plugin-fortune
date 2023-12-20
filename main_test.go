package traefik_plugin_fortune

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLottery(t *testing.T) {
	cfg := CreateConfig()
	cfg.Rate = 20

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "lottery-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/who", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	key := "X-FORTUNE"
	if req.Header.Get(key) != "Good" && req.Header.Get(key) != "Bad" {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
