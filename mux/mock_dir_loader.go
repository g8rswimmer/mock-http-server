package mux

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func loadFromDirectory(dir string) ([]*MockEndpoint, error) {
	endpoints := []*MockEndpoint{}
	pathErr := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
		}

		mh, err := loadMockEndpoint(path, info)
		if err != nil {
			return err
		}
		endpoints = append(endpoints, mh)
		return nil
	})
	if pathErr != nil {
		return nil, fmt.Errorf("file path walk err: %w", pathErr)
	}
	return endpoints, nil
}

func loadMockEndpoint(path string, info fs.FileInfo) (*MockEndpoint, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("load mock handler file open [%s] %w", path, err)
	}
	defer f.Close()

	ep := &MockEndpoint{}
	if err := json.NewDecoder(f).Decode(ep); err != nil {
		return nil, fmt.Errorf("load mock handler file [%s] decode %w", path, err)
	}
	return ep, nil
}
