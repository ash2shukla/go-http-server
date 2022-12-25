package core

import (
	"fmt"
	"strings"
)

type HttpResponseHeader struct {
	HttpVersion       string
	Status            int
	StatusDescription string
	ExtraHeaders      map[string]string
}

func CreateHttpResponse(responseHeaders HttpResponseHeader, content []byte) []byte {
	firstLine := fmt.Sprintf("%s %d %s", responseHeaders.HttpVersion, responseHeaders.Status, responseHeaders.StatusDescription)
	responseParts := []string{firstLine}

	for k, v := range responseHeaders.ExtraHeaders {
		headerLine := fmt.Sprintf("%s: %s", k, v)
		responseParts = append(responseParts, headerLine)
	}

	responseParts = append(responseParts, "\r\n")
	response := strings.Join(responseParts, "\r\n")
	responseBytes := []byte(response)
	responseBytes = append(responseBytes, content...)

	return responseBytes
}
