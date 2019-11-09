package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "Returns Pepper's Score",
			player:             "Pepper",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "Returns Floyd's Score",
			player:             "Floyd",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "Returns 404 on missing player",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := newGetScoreRequest(test.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, test.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), test.expectedScore)
		})
	}
	// t.Run("returns Pepper's score", func(t *testing.T) {
	// 	request := newGetScoreRequest("Pepper")
	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	assertStatus(t, response.Code, http.StatusOK)
	// 	assertResponseBody(t, response.Body.String(), "20")
	// })

	// t.Run("returns Floyd's score", func(t *testing.T) {
	// 	request := newGetScoreRequest("Floyd")
	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	assertStatus(t, response.Code, http.StatusOK)
	// 	assertResponseBody(t, response.Body.String(), "10")
	// })

	// t.Run("returns 404 on missing players", func(t *testing.T) {
	// 	request := newGetScoreRequest("Apollo")
	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	assertStatus(t, response.Code, http.StatusNotFound)
	// })
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner. got %q, want %q", store.winCalls[0], player)
		}
	})
}

// func TestRecordingWindsAndRetrievingThem(t *testing.T) {
// 	store := NewInMemoryPlayerStore()
// 	server := PlayerServer{store}
// 	player := "Pepper"

// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

// 	response := httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetScoreRequest(player))
// 	assertStatus(t, response.Code, http.StatusOK)

// 	assertResponseBody(t, response.Body.String(), "3")
// }

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}
