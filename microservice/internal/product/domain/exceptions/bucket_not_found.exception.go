package exceptions

type BucketNotFoundException struct{}

func (e *BucketNotFoundException) Error() string {
	return "Bucket de imagens não existe. Verifique a configuração do ambiente."
}
