package use_cases

import (
	"fmt"
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
)

type UploadProductImageUseCase struct {
	gateway gateways.ProductGateway
}

func NewUploadProductImageUseCase(gateway gateways.ProductGateway) *UploadProductImageUseCase {
	return &UploadProductImageUseCase{
		gateway: gateway,
	}
}

func (uc *UploadProductImageUseCase) Execute(productDTO dtos.UploadProductImageDTO) error {
	product, err := uc.gateway.FindByID(productDTO.ProductID)
	if err != nil {
		fmt.Printf("[UploadProductImageUseCase] Produto não encontrado: %v\n", err)
		return &exceptions.ProductNotFoundException{}
	}

	newFileName, err := product.AddImage(productDTO.FileName)
	if err != nil {
		fmt.Printf("[UploadProductImageUseCase] Erro ao adicionar imagem: %v\n", err)
		return &exceptions.InvalidProductImageException{}
	}

	fmt.Println("[UploadProductImageUseCase] Nome do arquivo gerado:", *newFileName)
	url, err := uc.gateway.UploadImage(*newFileName, productDTO.FileContent)
	if err != nil {
		fmt.Printf("[UploadProductImageUseCase] Erro ao fazer upload: %v\n", err)
		return err // Propaga o erro real para o handler
	}

	if err := uc.gateway.AddAndSetDefaultImage(product, url); err != nil {
		fmt.Printf("[UploadProductImageUseCase] Erro ao inserir/atualizar imagem no banco: %v\n", err)
		return &exceptions.InvalidProductDataException{}
	}

	fmt.Println("[UploadProductImageUseCase] Upload finalizado com sucesso!")
	return nil
}
