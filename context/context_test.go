package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	fmt.Println("in fetch")
	data := make(chan string, 1)

	go func() {
		var result string
		fmt.Println("inside go routine")
		// this slow build of the string is just to simulate a slow process
		// so we can interupt it midway
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("spy got cancelled")
				s.t.Log("spy store got cancelled")
				return
			default:
				fmt.Printf("making result %v\n", result)
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		fmt.Println("spy got cancelled pre return")
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	fmt.Println("in cancel")
	// s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "Hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		// store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		// response := httptest.NewRecorder()
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		// store.assertWasCancelled()
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("store was told to cancel")
// 	}
// }
