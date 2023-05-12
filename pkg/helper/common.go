package helper

import (
	"path/filepath"
	"runtime"
)

type ErrorStruct struct {
	Err  error
	Code int
}

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)
