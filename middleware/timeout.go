package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (t *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.Next == nil {
		t.Next = http.DefaultServeMux
	}
	ctx := r.Context()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	r.WithContext(ctx)
	ch := make(chan struct{})
	go func() {
		t.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
