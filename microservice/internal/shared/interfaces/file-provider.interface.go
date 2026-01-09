package interfaces

type IFileProvider interface {
	UploadFile(fileName string, fileContent []byte) error
	DeleteFile(fileName string) error
	DeleteFiles(filenames []string) error
	GetPresignedURL(fileName string) (string, error)
}
