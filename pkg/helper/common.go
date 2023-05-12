package helper

import (
	"math"
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

func GetMeta(limit, pages, totalRows *int) (totalPage int) {
	if *totalRows < *limit {
		*limit = *totalRows
	}

	if *pages < 1 {
		*pages = 1
	}

	if *totalRows > 0 && *limit > 0 {
		return int(math.Ceil(float64(*totalRows) / float64(*limit)))
	}
	return 0

}
