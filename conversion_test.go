// +build integration

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

func TestC2F_OK(t *testing.T) {
	client, route := testSetup(100)
	route = fmt.Sprintf("%s/c2f/100", route)

	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	bodyReader := resp.Body
	bodyBytes, err := ioutil.ReadAll(bodyReader)
	assert.NilError(t, err)

	var pairStruct presenter.DegreePair
	err = json.Unmarshal(bodyBytes, &pairStruct)
	assert.NilError(t, err)

	assert.Assert(t, 200 == resp.StatusCode)
	assert.Assert(t, pairStruct.F == 212)
}

func TestF2C_OK(t *testing.T) {
	client, route := testSetup(100)
	route = fmt.Sprintf("%s/f2c/32", route)

	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	bodyReader := resp.Body
	bodyBytes, err := ioutil.ReadAll(bodyReader)
	assert.NilError(t, err)

	var pairStruct presenter.DegreePair
	err = json.Unmarshal(bodyBytes, &pairStruct)
	assert.NilError(t, err)

	assert.Assert(t, 200 == resp.StatusCode)
	assert.Assert(t, pairStruct.C == 0)
}
