package internal

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/gohugoio/hugo/common/paths"
)

// ResourcePaths holds path information for a resource.
// All directories in here have Unix-style slashes, with leading slash, but no trailing slash.
// Empty directories are represented with an empty string.
type ResourcePaths struct {
	// This is the directory component for the target file or link.
	Dir string

	// Any base directory for the target file. Will be prepended to Dir.
	BaseDirTarget string

	// This is the directory component for the link will be prepended to Dir.
	BaseDirLink string

	// Set when publishing in a multihost setup.
	TargetBasePaths []string

	// This is the File component, e.g. "data.json".
	File string
}

func (d ResourcePaths) join(p ...string) string {
	var s string
	for i, pp := range p {
		if pp == "" {
			continue
		}
		if i > 0 && !strings.HasPrefix(pp, "/") {
			pp = "/" + pp
		}
		s += pp

	}
	if !strings.HasPrefix(s, "/") {
		s = "/" + s
	}
	return s
}

func (d ResourcePaths) TargetLink() string {
	return d.join(d.BaseDirLink, d.Dir, d.File)
}

func (d ResourcePaths) TargetPath() string {
	return d.join(d.BaseDirTarget, d.Dir, d.File)
}

func (d ResourcePaths) Path() string {
	return d.join(d.Dir, d.File)
}

func (d ResourcePaths) TargetPaths() []string {
	if len(d.TargetBasePaths) == 0 {
		return []string{d.TargetPath()}
	}

	var paths []string
	for _, p := range d.TargetBasePaths {
		paths = append(paths, p+d.TargetPath())
	}
	return paths
}

func (d ResourcePaths) TargetFilenames() []string {
	filenames := d.TargetPaths()
	for i, p := range filenames {
		filenames[i] = filepath.FromSlash(p)
	}
	return filenames
}

func (d ResourcePaths) FromTargetPath(targetPath string) ResourcePaths {
	targetPath = filepath.ToSlash(targetPath)
	dir, file := path.Split(targetPath)
	dir = paths.ToSlashPreserveLeading(dir)
	if dir == "/" {
		dir = ""
	}
	d.Dir = dir
	d.File = file
	d.BaseDirLink = ""
	d.BaseDirTarget = ""

	return d
}
