package util

import (
	"errors"
	"log"
	"os"
	"path"
	"strings"
)

func SrcDirOf(pDir string) (string, error) {
	goPath, ok := os.LookupEnv("GOPATH")
	if !ok {
		log.Fatal("No GOPATH environment variable")
	}

	for _, p := range strings.Split(goPath, ":") {
		p1 := path.Join(p, "src", pDir)
		if _, err := os.Stat(p1); err == nil {
			return path.Join(p, "src"), nil
		}
	}

	return "", errors.New("No source directory for " + pDir)
}

func SrcDir() string {
	goPath, ok := os.LookupEnv("GOPATH")
	if !ok {
		log.Fatal("No GOPATH environment variable")
	}

	// TODO: Find a portable way to split GOPATH
	return path.Join(strings.Split(goPath, ":")[0], "src")
}
