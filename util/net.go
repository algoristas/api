package util

import (
	"net"
	"net/http"
	"time"
)

// BuildRequest builds URL with the provided query params.
//
// Generates the format as follows:
//    http://URL_REQUESTED?param1=value1&param2=value2&...&paramn=valuen
// It returns a String that is the well format URL to make the request.
func BuildRequest(method, urlStr string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return req, err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

// NewHTTPClient creates a new HTTP client to make request.
//
// It returns a pointer of type http.Client to make request to servers.
// We do the request in this way so we can control the timeout, otherwise we may have troubles
// if the server we are requesting for fails.
func NewHTTPClient() *http.Client {
	const timeoutRequest = 10 * time.Second // Requests timeouts
	const timeoutDial = 180 * time.Second   // limits the time spent establishing a TCP connection,
	//   often around 3 minutes.
	const timeoutTSL = 30 * time.Second // limits the time spent performing the TLS handshake.

	// cap the TCP connect and TLS handshake timeouts
	// as well as establishing an end-to-end request timeout.
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeoutDial,
		}).Dial,
		TLSHandshakeTimeout: timeoutTSL,
	}
	httpClient := &http.Client{
		Timeout:   timeoutRequest,
		Transport: netTransport,
	}

	return httpClient
}
