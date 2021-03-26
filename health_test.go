package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/TomasCruz/grpc-web-fahrenheit/presenter"
	"gotest.tools/v3/assert"
)

func TestHealthOK(t *testing.T) {
	client, route := testSetup(100)
	route = fmt.Sprintf("%s/health", route)

	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	bodyReader := resp.Body
	bodyBytes, err := ioutil.ReadAll(bodyReader)
	assert.NilError(t, err)

	var health presenter.Health
	err = json.Unmarshal(bodyBytes, &health)
	assert.NilError(t, err)

	assert.Assert(t, 200 == resp.StatusCode)
	assert.Assert(t, health.Status == "Up")
}
