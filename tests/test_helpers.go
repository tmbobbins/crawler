package tests

import (
	"fmt"
	"net/http"
	"os"
)

func FileSystemClient() http.Client {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	return http.Client{Transport: transport}
}

func CurrentWorkingDirectory() string {
	cwd, _ := os.Getwd()
	return cwd
}

func RelativeFile(relativeLocation string) string {
	return fmt.Sprintf("file://%s/%s", CurrentWorkingDirectory(), relativeLocation)
}
