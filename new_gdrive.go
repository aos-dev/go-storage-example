package example

import (
	"fmt"
	"os"

	gdrive "go.beyondstorage.io/services/gdrive"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func NewGdrive() (types.Storager, error) {
	return gdrive.NewStorager(
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Example Value: file:<abs_path_of_credential>
		pairs.WithCredential(os.Getenv("STORAGE_GDRIVE_CREDENTIAL")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_GDRIVE_NAME")),
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_GDRIVE_WORKDIR")),
	)
}

func NewGdriveFromString() (types.Storager, error) {
	str := fmt.Sprintf(
		"gdrive://%s%s?credential=%s",
		os.Getenv("STORAGE_GDRIVE_NAME"),
		os.Getenv("STORAGE_GDRIVE_WORKDIR"),
		os.Getenv("STORAGE_GDRIVE_CREDENTIAL"),
	)
	return services.NewStoragerFromString(str)
}
