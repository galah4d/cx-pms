package utils

import (
	"errors"
	"os"
)

// FileExists reports whether the named file exists as a boolean
func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists as a boolean
func DirExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

func DirInit(path string) error {
	if DirExists(path) {
		return errors.New("directory already exists")
	}
	return os.MkdirAll(path, os.ModePerm)
}

func GetGOPATH() (string, error) {
	if gopath, ok := os.LookupEnv("GOPATH"); ok {
		return gopath, nil
	}
	return "", errors.New("GOPATH not ser")
}
