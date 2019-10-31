package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you'd like
		}
		fmt.Fprint(w, data)
		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	fmt.Println("in goroutine")
		// 	data <- store.Fetch()
		// }()

		// // use select to race between fetching data and receiving cancel request
		// select {
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }

	}
}
