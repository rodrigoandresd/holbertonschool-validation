package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server(t *testing.T) {
	if testing.Short() {
		t.Skip("Flag `-short` provided: skipping Integration Tests.")
	}

	tests := []struct {
		name         string
		URI          string
		responseCode int
		body         string
	}{
		{
			name:         "Home page",
			URI:          "",
			responseCode: 404,
			body:         "404 page not found\n",
		},
		{
			name:         "Health page",
			URI:          "/hello?name=Holberton",
			responseCode: 200,
			body:         "Hello Holberton!",
		},
		// new
		{
			name:         "Health page2",
			URI:          "/health",
			responseCode: 200,
			body:         "ALIVE",
		},
		{
			name:         "Undefined page",
			URI:          "/test",
			responseCode: 404,
			body:         "404 page not found\n",
		},
		{
			name:         "Undefined name",
			URI:          "/hello?name",
			responseCode: 200,
			body:         "Hello there!",
		},
		{
			name:         "Another queryString halo",
			URI:          "/hello?halo=Haloooo!",
			responseCode: 200,
			body:         "Hello there!",
		},
		{
			name:         "Without any param",
			URI:          "/hello?name=test&name=lol",
			responseCode: 200,
			body:         "Hello lol!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(setupRouter())
			defer ts.Close()

			res, err := http.Get(ts.URL + tt.URI)
			if err != nil {
				t.Fatal(err)
			}

			// Check that the status code is what you expect.
			expectedCode := tt.responseCode
			gotCode := res.StatusCode
			if gotCode != expectedCode {
				t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
			}

			// Check that the response body is what you expect.
			expectedBody := tt.body
			bodyBytes, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			gotBody := string(bodyBytes)
			if gotBody != expectedBody {
				t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
			}
		})
	}
}