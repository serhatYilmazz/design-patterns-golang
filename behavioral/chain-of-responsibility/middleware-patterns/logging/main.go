package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type ResponseHeader struct {
	handler http.Handler
	headerName string
	headerValue string
}

func NewResponseHeader(handlerToWrap http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.headerName, rh.headerValue)
	rh.handler.ServeHTTP(w, r)
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(writer, request)
	log.Printf("%s %s %v", request.Method, request.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("Current time is %v", time)))
}

func main() {
	mux := http.NewServeMux()
	wrappedMux := NewLogger(NewResponseHeader(mux, "X-My-Header", "Custom Header"))

	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/time", CurrentTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
