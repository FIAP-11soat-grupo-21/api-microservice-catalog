package mock_interfaces

// MockFileProvider é um mock simples para IFileProvider

type MockFileProvider struct{}

func (m *MockFileProvider) UploadFile(fileName string, fileContent []byte) error {
	return nil
}
func (m *MockFileProvider) DeleteFile(fileName string) error {
	return nil
}
func (m *MockFileProvider) GetPresignedURL(fileName string) (string, error) {
	return "http://localhost/" + fileName, nil
}
func (m *MockFileProvider) DeleteFiles(fileNames []string) error {
	return nil
}
