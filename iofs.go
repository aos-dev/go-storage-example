//go:build go1.16
// +build go1.16

package example

import (
	"io/fs"
	"io/ioutil"
	"log"

	"go.beyondstorage.io/v5/pkg/fswrap"
	"go.beyondstorage.io/v5/types"
)

func FSOpen(store types.Storager, path string) {
	fsys := fswrap.Fs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ReadAll %v: %v", path, err)
	}

	log.Printf("read content: %s", data)
}

func FSReadFile(store types.Storager, path string) {
	fsys := fswrap.Fs(store)

	rf, ok := fsys.(fs.ReadFileFS)
	if !ok {
		log.Fatalf("fs.ReadFileFS unimplemented")
	}

	data, err := rf.ReadFile(path)
	if err != nil {
		log.Fatalf("ReadFile %v: %v", path, err)
	}

	log.Printf("read content length: %d", len(data))
}

func FSReadDir(store types.Storager, path string) {
	fsys := fswrap.Fs(store)

	rd, ok := fsys.(fs.ReadDirFS)
	if !ok {
		log.Fatalf("fs.ReadDirFS unimplemented")
	}

	list, err := rd.ReadDir(path)
	if err != nil {
		log.Fatalf("ReadDir %v: %v", path, err)
	}

	for _, entry := range list {
		log.Printf("DirEntry: %s, %v", entry.Name(), entry.Type())
	}
}

func FSGlob(store types.Storager, pattern string) {
	fsys := fswrap.Fs(store)

	g, ok := fsys.(fs.GlobFS)
	if !ok {
		log.Fatalf("fs.GlobFS unimplemented")
	}

	names, err := g.Glob(pattern)
	if err != nil {
		log.Fatalf("Glob %v: %v", pattern, err)
	}

	for _, name := range names {
		log.Printf("file name: %s", name)
	}
}

func FSStat(store types.Storager, path string) {
	fsys := fswrap.Fs(store)

	s, ok := fsys.(fs.StatFS)
	if !ok {
		log.Fatalf("fs.StatFS unimplemented")
	}

	info, err := s.Stat(path)
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	log.Printf("file name: %s", info.Name())
}

func FileRead(store types.Storager, path string) {
	fsys := fswrap.Fs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	size := info.Size()
	data := make([]byte, 0, size+1)

	n, err := f.Read(data)
	if err != nil {
		log.Fatalf("Read %v: %v", path, err)
	}

	log.Printf("read data size: %d", n)
}
