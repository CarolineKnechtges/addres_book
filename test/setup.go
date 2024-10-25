// Package tests/setup.go
package tests

import (
	"os"
	"path/filepath"
	"testing"
)

// InitTestEnvironment initializes the test environment by setting the working directory.
func InitTestEnvironment(t *testing.T) {
	rootDir, err := filepath.Abs(filepath.Join("..")) // 获取项目根目录的绝对路径
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	err = os.Chdir(rootDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
}
