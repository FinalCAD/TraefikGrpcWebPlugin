package TraefikGrpcWebPlugin_test

import (
	"context"
	"github.com/FinalCAD/TraefikGrpcWebPlugin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	cfg := TraefikGrpcWebPlugin.NewConfig()
	cfg.AllowOrigins = []string{"*"}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := TraefikGrpcWebPlugin.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

}
