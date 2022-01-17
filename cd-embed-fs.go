package cdfs

import (
	"embed"
	"io/fs"
)

type FS struct {
	inner embed.FS
	sub   string
}

func CD(inner embed.FS, sub string) FS {
	return FS{inner: inner, sub: sub}
}

// Open opens the named file for reading and returns it as an fs.File.
func (f FS) Open(name string) (fs.File, error) {
	return f.inner.Open(f.sub + "/" + name)
}

// ReadDir reads and returns the entire named directory.
func (f FS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.inner.ReadDir(f.sub + "/" + name)
}

// ReadFile reads and returns the content of the named file.
func (f FS) ReadFile(name string) ([]byte, error) {
	return f.inner.ReadFile(f.sub + "/" + name)
}
