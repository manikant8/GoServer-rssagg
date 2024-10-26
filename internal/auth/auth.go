package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GeAPIKey extracts an API Key from
// the header of an HTTP request
// Example:
// Authorization: ApiKey {insert api key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authenticatioon info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
