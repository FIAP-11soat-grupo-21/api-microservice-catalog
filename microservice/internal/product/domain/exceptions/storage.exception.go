package exceptions

type DeleteImagesStorageException struct {
	Message string
}
type BucketNotFoundException struct {
	Message string
}

func (e *DeleteImagesStorageException) Error() string {
	if e.Message == "" {
		return "Failed to delete file(s) from storage"
	}

	return e.Message
}

func (e *BucketNotFoundException) Error() string {
	if e.Message == "" {
		return "Bucket S3 não existe ou é inválido"
	}
	return e.Message
}
