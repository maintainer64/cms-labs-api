// Package http_debug общие компоненты для Fiber транспортного уровня
package http_debug

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"gitlab.com/a10869/api-modules/shared/logs"
)

var (
	log = logs.NewZeroLogger("http_transport")
)

type LoggingTransport struct{}

func (s *LoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	bytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)
	// err is returned after dumping the response

	respBytes, _ := httputil.DumpResponse(resp, true)
	bytes = append(bytes, respBytes...)

	log.Debug().Msg(fmt.Sprintf("%s\n", bytes))
	fmt.Printf("%s\n", bytes)

	return resp, err
}
