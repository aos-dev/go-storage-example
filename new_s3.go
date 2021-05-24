package example

import (
	"os"

	s3 "github.com/beyondstorage/go-service-s3/v2"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewS3() (types.Storager, error) {
	return s3.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_S3_WORKDIR")),
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console.
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_S3_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// endpoint is default to amazon s3's endpoint.
		// If using s3 compatible services, please input their endpoint.
		//
		// Example Value: https:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_S3_ENDPOINT")),
		// location: https://beyondstorage.io/docs/go-storage/pairs/location
		//
		// For s3, location is the bucket's zone.
		// For s3 compatible services, location could be ignored or has other value,
		// please refer to their documents.
		//
		// Example Value: ap-east-1
		pairs.WithLocation(os.Getenv("STORAGE_S3_LOCATION")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_S3_NAME")),
	)
}
