package wiseman

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpServer(addr string, h http.Handler) error {
	uaddr, err := net.ResolveUnixAddr("unix", addr)
	if err != nil {
		return err
	}
	log.Printf("server start at addr %v", uaddr)
	listener, err := net.ListenUnix("unix", uaddr)
	if err != nil {
		return fmt.Errorf("can't listen addr %q: %w", addr, err)
	}
	defer func() { _ = listener.Close() }()
	return http.Serve(listener, h)
}

func dumpHttpRequest(r *http.Request) (data map[string]interface{}, err error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"body":   string(b),
		"header": r.Header,
	}, nil
}

var SampleHttpHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("content-type", "application/json")
	d, err := dumpHttpRequest(r)
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(d)
	_, _ = w.Write(b)
}

func HttpHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", SampleHttpHandler)
	r.HandleFunc("/products", SampleHttpHandler)
	r.HandleFunc("/articles", SampleHttpHandler)
	return r
}
