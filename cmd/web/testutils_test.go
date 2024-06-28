package main

import (
	"bytes"
	"github.com/alexedwards/scs/v2"
	"github.com/dimastephen/snippetbox/internal/mocks"
	"github.com/go-playground/form/v4"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"
)

func newTestApplication(t *testing.T) *application {
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}
	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       log.New(io.Discard, "ERROR", 0),
		infoLog:        log.New(io.Discard, "LOG", 0),
		users:          &mocks.UserModel{},
		snippets:       &mocks.SnippetModel{},
		formDecoder:    formDecoder,
		templateCache:  templateCache,
		sessionManager: sessionManager,
	}
	return app
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, url string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + url)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)

}
