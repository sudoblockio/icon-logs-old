package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// List test
func TestLogsEndpointList(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List limit and skip test
func TestLogsEndpointListLimitSkip(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?limit=20&skip=20")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List number test
func TestLogsEndpointListNumber(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	// Get latest log
	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?limit=1")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))

	// Get testable number
	logNumber := strconv.FormatUint(uint64(bodyMap[0].(map[string]interface{})["block_number"].(float64)), 10)

	// Test number
	resp, err = http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?block_number=" + logNumber)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err = ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap = make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List transaction hash test
func TestLogsEndpointListTransactionHash(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	// Get latest log
	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?limit=1")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))

	// Get testable transaction hash
	logHash := bodyMap[0].(map[string]interface{})["transaction_hash"].(string)

	// Test transaction hash
	resp, err = http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?transaction_hash=" + logHash)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err = ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap = make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List score address test
func TestLogsEndpointListScoreAddress(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	// Get latest log
	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?limit=1")
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))

	// Get testable score address
	logScoreAddress := bodyMap[0].(map[string]interface{})["address"].(string)

	// Test score address
	resp, err = http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?score_address=" + logScoreAddress)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err = ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap = make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}

// List method test
func TestLogsEndpointListMethod(t *testing.T) {
	assert := assert.New(t)

	logsServiceURL := os.Getenv("LOGS_SERVICE_URL")
	if logsServiceURL == "" {
		logsServiceURL = "http://localhost:8000"
	}
	logsServiceRestPrefx := os.Getenv("LOGS_SERVICE_REST_PREFIX")
	if logsServiceRestPrefx == "" {
		logsServiceRestPrefx = "/api/v1"
	}

	// Get latest log
	logMethod := ""
	skip := 0
	for {
		// Keep interating skip until non-null method
		resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?limit=1&skip=" + strconv.Itoa(skip))
		assert.Equal(nil, err)
		assert.Equal(200, resp.StatusCode)

		defer resp.Body.Close()

		// Test headers
		assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Equal(nil, err)

		bodyMap := make([]interface{}, 0)
		err = json.Unmarshal(bytes, &bodyMap)
		assert.Equal(nil, err)
		assert.NotEqual(0, len(bodyMap))

		// Get testable method
		logMethod = bodyMap[0].(map[string]interface{})["method"].(string)
		if logMethod != "" {
			break
		}

		skip++
	}

	// Test method
	resp, err := http.Get(logsServiceURL + logsServiceRestPrefx + "/logs?method=" + logMethod)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	defer resp.Body.Close()

	// Test headers
	assert.NotEqual("0", resp.Header.Get("X-TOTAL-COUNT"))

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	bodyMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &bodyMap)
	assert.Equal(nil, err)
	assert.NotEqual(0, len(bodyMap))
}
