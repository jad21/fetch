package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// NoBody is an io.ReadCloser with no bytes. Read always returns EOF
// and Close always returns nil. It can be used in an outgoing client
// request to explicitly signal that a request has zero bytes.
// An alternative, however, is to simply set Request.Body to nil.
var NoBody = noBody{}

type noBody struct{}

func (noBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (noBody) Close() error                     { return nil }
func (noBody) WriteTo(io.Writer) (int64, error) { return 0, nil }

func TestDefaultOptions(t *testing.T) {
	var opt *Options = DefaultOptions()
	if opt.Timeout != DefaultTimeout {
		t.Errorf("Expected timeout [%v], but got [%v]", DefaultTimeout, opt.Timeout)
	}
	if opt.Host != "" {
		t.Errorf("Expected timeout [%v], but got [%v]", DefaultTimeout, opt.Timeout)
	}
	if opt.Header == nil {
		t.Error("Expected header already initialized, but got nil")
	}
}

func TestNew(t *testing.T) {
	defaultFetch := New()
	if defaultFetch.Option == nil {
		t.Error("Expected Option not nil, but got nil")
	}
	if defaultFetch.Transport == nil {
		t.Error("Expected Transport is nil, but got nil")
	}
	if defaultFetch.Client == nil {
		t.Error("Expected Client not nil, but got nil")
	}
}

func TestFetch_IsJSON(t *testing.T) {
	handlerContentTypeTest := func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("Content-Type")
		if !strings.EqualFold(contentType, "application/json") {
			t.Errorf("Expected [Content-Type=application/json], but got [Content-Type=%s]", contentType)
		}
	}

	t.Run("Test-IsJSON", func(t *testing.T) {
		f := New().IsJSON()

		server := httptest.NewServer(http.HandlerFunc(handlerContentTypeTest))
		defer server.Close()

		f.Get(server.URL, nil)
	})
}

func TestMakeResponse(t *testing.T) {

	t.Run("test-MakeResponseOK", func(t *testing.T) {
		body := "Jose Delgado not found"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, body)
		}))
		defer ts.Close()

		res, err := New().Get(ts.URL, nil)
		if err != nil {
			log.Fatal(err)
		}

		if s := res.String(); !strings.EqualFold(body, s) {
			t.Errorf("Expected body [%s], but got [%s]", body, s)
		}

		if http.StatusNotFound != res.StatusCode {
			t.Errorf("Expected body [%d], but got [%d]", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("test-MakeResponseTimeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
		}))
		defer ts.Close()

		res, err := New(
			WithTimeout(time.Duration(10*time.Millisecond)),
		).Get(ts.URL, nil)
		if err == nil {
			t.Error("Expected timeout error, but got none error")
		}

		if http.StatusGatewayTimeout != res.StatusCode {
			t.Errorf("Expected body [%d], but got [%d]", http.StatusGatewayTimeout, res.StatusCode)
		}
	})
}

func serverHandlerMock(handlerFunc http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handlerFunc))
}
