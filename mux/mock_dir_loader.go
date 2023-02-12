package mux

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func loadFromDirectory(dir string) ([]*MockHandler, error) {
	handlers := []*MockHandler{}
	pathErr := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
		}

		mh, err := loadMockHandler(path, info)
		if err != nil {
			return err
		}
		handlers = append(handlers, mh)
		return nil
	})
	if pathErr != nil {
		return nil, fmt.Errorf("file path walk err: %w", pathErr)
	}
	return handlers, nil
}

func loadMockHandler(path string, info fs.FileInfo) (*MockHandler, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("load mock handler file open [%s] %w", path, err)
	}
	defer f.Close()

	h := &MockHandler{}
	if err := json.NewDecoder(f).Decode(h); err != nil {
		return nil, fmt.Errorf("load mock handler file [%s] decode %w", path, err)
	}
	return h, nil
}
