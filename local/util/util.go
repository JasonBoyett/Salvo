package util

import (
	"fmt"
	"os"
)

func FindDir(path string) (string, error) {
	fmt.Println(path)
	base, err := os.Open(path)
	if err != nil {
		return "", err
	}
	baseInfom, err := base.Stat()
	if err != nil {
		return "", err
	}
	if !baseInfom.IsDir() {
		for i := len(path) - 1; i > 0; i-- {
			if path[i] == os.PathSeparator {
				path = path[:i]
				result, err := FindDir(path)
				if err != nil {
					return "", err
				}
				return result, nil
			}
		}
	}
	return path, nil
}
