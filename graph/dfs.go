package graph

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// 指定したディレクトリを走査し、ディレクト内にあるすべてのファイルを返す。
// 探索は再帰関数を適応することで深さ優先探索(DFS)を実現。
func walkDir(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("ディレクトリのパスを指定ください")
	}

	var files []string
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range dirs {
		if !f.IsDir() {
			files = append(files, f.Name())
			continue
		}
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		fullPath := filepath.Join(path, f.Name())
		fs, err := walkDir(fullPath)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}
	return files, nil
}
