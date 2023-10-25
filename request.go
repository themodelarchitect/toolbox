package toolbox

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type RequestMethod byte

const (
	HEAD RequestMethod = iota
	GET
	POST
	PUT
	DELETE
)

func (method RequestMethod) String() string {
	switch method {
	case HEAD:
		return "HEAD"
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return fmt.Sprintf("Unknown(%d)", method)
	}
}

func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func ParseURL(someURL string) string {
	u, err := url.Parse(someURL)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("\nScheme: %s\n  Host: %s\n  Path: %s\n Query: %s\n",
		u.Scheme, u.Hostname(), u.Path, u.RawQuery)
}

func BaseURL(url string) string {
	idx := strings.LastIndex(url, "/")
	if idx != -1 {
		url = url[:idx]
	}
	return url
}

func SendRequest(method RequestMethod, url, body string, headers map[string]string) (int, []byte, error) {
	var statusCode int
	var b []byte

	// create a request
	req, err := http.NewRequest(method.String(), url, strings.NewReader(body))
	if err != nil {
		return statusCode, b, err
	}

	// NOTE this !! -You need to set Req.Close to true (the defer on resp.Body.Close() syntax used in the examples is not enough)
	req.Close = true

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return statusCode, b, err
	}

	defer func(Body io.ReadCloser) {
		// Ignore error explicitly
		_ = Body.Close()
	}(resp.Body)

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, b, err
	}

	return resp.StatusCode, b, nil
}
