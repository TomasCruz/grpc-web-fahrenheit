package main

import (
	"fmt"
	"net/http"
	"time"
)

func testSetup(millis int) (httpClient http.Client, route string) {
	port := readAndCheckIntEnvVar("GRPC_WEB_PORT")
	route = fmt.Sprintf("http://localhost:%s", port)

	timeout := time.Duration(time.Duration(millis) * time.Millisecond)
	httpClient = http.Client{Timeout: timeout}

	return
}
