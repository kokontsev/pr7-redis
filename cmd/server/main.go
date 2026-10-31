package main

import (
	"context"
	"errors"
	"example.com/pr7/internal/cache"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:16379"
	}
	c := cache.New(addr)
	defer c.Close()
	mux := http.NewServeMux()
	for _, op := range []string{"set", "get", "ttl"} {
		mux.HandleFunc("GET /"+op, func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()
			k := r.URL.Query().Get("key")
			if k == "" {
				http.Error(w, "key required", 400)
				return
			}
			switch op {
			case "set":
				v := r.URL.Query().Get("value")
				if v == "" {
					http.Error(w, "value required", 400)
					return
				}
				if err := c.Set(ctx, k, v, 10*time.Second); err != nil {
					http.Error(w, "redis unavailable", 503)
					return
				}
				fmt.Fprintf(w, "OK: %s=%s (TTL 10s)\n", k, v)
			case "get":
				v, e := c.Get(ctx, k)
				if errors.Is(e, redis.Nil) {
					http.Error(w, "key not found", 404)
					return
				}
				if e != nil {
					http.Error(w, "redis unavailable", 503)
					return
				}
				fmt.Fprintf(w, "VALUE: %s=%s\n", k, v)
			case "ttl":
				v, e := c.TTL(ctx, k)
				if e != nil {
					http.Error(w, "redis unavailable", 503)
					return
				}
				fmt.Fprintf(w, "TTL for %s: %v\n", k, v)
			}
		})
	}
	a := os.Getenv("HTTP_ADDR")
	if a == "" {
		a = ":8080"
	}
	log.Fatal((&http.Server{Addr: a, Handler: mux, ReadHeaderTimeout: 5 * time.Second}).ListenAndServe())
}
