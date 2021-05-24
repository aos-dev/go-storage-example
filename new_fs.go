package example

import (
	"os"

	fs "github.com/beyondstorage/go-service-fs/v3"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewFs() (types.Storager, error) {
	return fs.NewStorager(
		// WorkDir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_FS_WORKDIR")),
	)
}
