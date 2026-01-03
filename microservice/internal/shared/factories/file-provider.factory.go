package factories

import (
	file_provider "tech_challenge/internal/shared/infra/file_provider"
)

func NewFileProvider() *file_provider.S3FileProvider {
	return file_provider.NewS3FileProvider()
}
