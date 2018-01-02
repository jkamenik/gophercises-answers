package urlshort

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// If both the default handler and the map has no value then throw an error
func TestNilMap(t *testing.T) {
	_, err := mapHandler(nil, nil)
	if err != ErrEmptyMapNoDefault {
		t.Errorf("Expected '%v' got '%v'", ErrEmptyMapNoDefault, err)
	}
}

type testHandler struct {
	body string
	code int
}

func (t testHandler) ServeHTTP(response http.ResponseWriter, request *(http.Request)) {
	fmt.Println("Test ServeHTTP")
	response.WriteHeader(t.code)
	response.Write([]byte(t.body))
	return
}

func fakeRequests() []*http.Request {
	foo, _ := http.NewRequest("GET", "/foo", nil)
	bar, _ := http.NewRequest("POST", "/bar", nil)
	root, _ := http.NewRequest("GET", "/", nil)

	return []*http.Request{foo, bar, root, nil}
}

func TestFallback(t *testing.T) {
	handler, err := mapHandler(nil, testHandler{"Testing", http.StatusGone})
	if err != nil {
		t.Errorf("Failed to created handler: %v", err)
	}

	for _, request := range fakeRequests() {
		// Record each request
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, request)

		// Should always be the default handler
		// We expect the request to drop into the default mux
		if rr.Code != http.StatusGone {
			t.Errorf("expected code '%v', got '%v'", http.StatusGone, rr.Code)
		}

		if rr.Body.String() != "Testing" {
			t.Errorf("expected body '%v', got '%v'", "Testing", rr.Body.String())
		}
	}
}

func TestNonFallback(t *testing.T) {
	handler, err := mapHandler(map[string]string{
		"/foo": "/bar",
	}, testHandler{"Default", http.StatusNotModified})
	if err != nil {
		t.Errorf("Failed to create handler: %v", err)
	}

	requests := fakeRequests()
	fmt.Println(requests)

	// the first request is for "/foo"
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, requests[0])

	// Should always be the default handler
	// We expect the request to drop into the default mux
	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("expected code '%v', got '%v'", http.StatusMovedPermanently, rr.Code)
	}

	if rr.Body.String() != "/bar" {
		t.Errorf("expected body '%v', got '%v'", "/bar", rr.Body.String())
	}

	// Anything else should return the default
	for i, request := range fakeRequests() {
		if i == 0 {
			continue
		}

		// Record each request
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, request)

		// Should always be the default handler
		// We expect the request to drop into the default mux
		if rr.Code != http.StatusNotModified {
			t.Errorf("expected code '%v', got '%v'", http.StatusNotModified, rr.Code)
		}

		if rr.Body.String() != "Default" {
			t.Errorf("expected body '%v', got '%v'", "Default", rr.Body.String())
		}
	}
}
