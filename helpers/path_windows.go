// +build windows

package helpers

import (
	"fmt"
	"os"
)

// FindCWD returns the current working directory
func FindCWD() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Can't get CWD information: %v", err)
	}
	return cwd, nil
}
