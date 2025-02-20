// Package source contains the types and functions related to source files.
package source

import (
	"path/filepath"
	"runtime"

	"github.com/gohugoio/hugo/hugofs/glob"

	"github.com/spf13/afero"

	"github.com/gohugoio/hugo/helpers"
)

// SourceSpec abstracts language-specific file creation.
// TODO(bep) rename to Spec
type SourceSpec struct {
	*helpers.PathSpec

	SourceFs afero.Fs

	shouldInclude func(filename string) bool
}

// NewSourceSpec initializes SourceSpec using languages the given filesystem and PathSpec.
func NewSourceSpec(ps *helpers.PathSpec, inclusionFilter *glob.FilenameFilter, fs afero.Fs) *SourceSpec {
	shouldInclude := func(filename string) bool {
		if !inclusionFilter.Match(filename, false) {
			return false
		}
		if ps.Cfg.IgnoreFile(filename) {
			return false
		}

		return true
	}

	return &SourceSpec{shouldInclude: shouldInclude, PathSpec: ps, SourceFs: fs}
}

// IgnoreFile returns whether a given file should be ignored.
func (s *SourceSpec) IgnoreFile(filename string) bool {
	if filename == "" {
		if _, ok := s.SourceFs.(*afero.OsFs); ok {
			return true
		}
		return false
	}

	base := filepath.Base(filename)

	if len(base) > 0 {
		first := base[0]
		last := base[len(base)-1]
		if first == '.' ||
			first == '#' ||
			last == '~' {
			return true
		}
	}

	if !s.shouldInclude(filename) {
		return true
	}

	if runtime.GOOS == "windows" {
		// Also check the forward slash variant if different.
		unixFilename := filepath.ToSlash(filename)
		if unixFilename != filename {
			if !s.shouldInclude(unixFilename) {
				return true
			}
		}
	}

	return false
}
