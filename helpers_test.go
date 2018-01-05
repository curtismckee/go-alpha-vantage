package av

import (
	"net/http"
	"net/url"
	"bytes"
)

const (
	testApiKey = "test"
)

type responseConnection struct{ response *http.Response }

func NewResponseConnection(res *http.Response) Connection             { return &responseConnection{res} }
func (c responseConnection) Request(*url.URL) (*http.Response, error) { return c.response, nil }

type errorConnection struct{ err error }

func NewErrorConnection(err error) Connection                      { return &errorConnection{err} }
func (c errorConnection) Request(*url.URL) (*http.Response, error) { return nil, c.err }

type ResetBuffer struct {
	contents string
	buf      *bytes.Buffer
}

func NewBuffCloser(contents string) *ResetBuffer {
	buff := &ResetBuffer{
		contents: contents,
	}
	buff.Restart()
	return buff
}

func (b *ResetBuffer) Restart() {
	b.buf = bytes.NewBufferString(b.contents)
}

func (b *ResetBuffer) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

func (b *ResetBuffer) Close() error {
	b.Restart()
	return nil
}
