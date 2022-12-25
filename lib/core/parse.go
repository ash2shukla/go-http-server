package core

import (
	"errors"
	"strings"
)

var validMethods = StrArray{"GET", "POST", "PUT", "DELETE", "HEAD", "TRACE", "OPTIONS"}

type HttpBasicHeaderInfo struct {
	Method      string
	HttpVersion string
	Path        string
}

func GetHttpBasicInfo(reqInitLine string) (HttpBasicHeaderInfo, error) {
	defaultResponse := HttpBasicHeaderInfo{"", "", ""}
	parts := strings.Split(reqInitLine, " ")

	if len(parts) != 3 {
		return defaultResponse, errors.New("Invalid HTTP Request")
	}

	method, path, httpVersion := parts[0], parts[1], parts[2]

	if !validMethods.Has(method) {
		return defaultResponse, errors.New("Invalid HTTP Method")
	}

	if !strings.HasPrefix(httpVersion, "HTTP") {
		return defaultResponse, errors.New("Invalid HTTP Version")
	}

	return HttpBasicHeaderInfo{method, httpVersion, path}, nil
}
