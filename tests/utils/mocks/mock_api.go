// tests/utils/mocks/mock_api.go
package mocks

import (
	"net/http"
	"net/http/httptest"
)

type MockAPI struct {
	server *httptest.Server
}

func NewMockAPI() *MockAPI {
	mock := &MockAPI{}
	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/images/search":
			w.Write([]byte(`[{"url": "http://example.com/cat.jpg"}]`))
		case "/v1/breeds":
			w.Write([]byte(`[{"id": "beng", "name": "Bengal"}]`))
		case "/v1/favourites":
			w.Write([]byte(`{"id": 123, "message": "Success"}`))
		case "/v1/votes":
			w.Write([]byte(`{"message": "SUCCESS"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	return mock
}

func (m *MockAPI) URL() string {
	return m.server.URL
}

func (m *MockAPI) Close() {
	m.server.Close()
}
