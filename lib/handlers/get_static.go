package handlers

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ash2shukla/go-http-server/config"
	"github.com/ash2shukla/go-http-server/lib/core"
)

func GetStaticHandler(basicInfo core.HttpBasicHeaderInfo) []byte {
	pathParts := strings.Split(basicInfo.Path, "?")
	config := config.GetConfig()
	path := filepath.Join(config.Serve.StaticDir, pathParts[0])

	isDir, _ := core.IsDirectory(path)
	if isDir {
		path = filepath.Join(path, "index.html")
	}

	dat, err := ioutil.ReadFile(path)

	if err == nil {
		responseHeader := core.HttpResponseHeader{
			HttpVersion:       basicInfo.HttpVersion,
			Status:            200,
			StatusDescription: "OK",
			ExtraHeaders:      make(map[string]string),
		}
		response := core.CreateHttpResponse(responseHeader, dat)
		return response
	} else if errors.Is(err, os.ErrNotExist) {
		responseHeader := core.HttpResponseHeader{
			HttpVersion:       basicInfo.HttpVersion,
			Status:            404,
			StatusDescription: "Not found",
			ExtraHeaders:      make(map[string]string),
		}
		response := core.CreateHttpResponse(responseHeader, dat)
		return response
	} else {
		responseHeader := core.HttpResponseHeader{
			HttpVersion:       basicInfo.HttpVersion,
			Status:            500,
			StatusDescription: "Internal server error",
			ExtraHeaders:      make(map[string]string),
		}
		response := core.CreateHttpResponse(responseHeader, dat)
		return response
	}
}
