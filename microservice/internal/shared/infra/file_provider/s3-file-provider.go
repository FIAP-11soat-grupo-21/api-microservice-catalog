package file_service

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"tech_challenge/internal/product/domain/exceptions"
	"tech_challenge/internal/shared/config/env"
)

type S3FileProvider struct {
	client     *s3.Client
	bucketName string
}

func NewS3FileProvider() *S3FileProvider {
	cfgEnv := env.GetConfig()
	fmt.Printf("[DEBUG] Config S3: Bucket=%s, Endpoint=%s, Region=%s\n", cfgEnv.AWS.S3.BucketName, cfgEnv.AWS.S3.Endpoint, cfgEnv.AWS.Region)
	var client *s3.Client

	if cfgEnv.AWS.S3.Endpoint == "" {
		// Endpoint vazio: usar AWS S3 real
		awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfgEnv.AWS.Region))
		if err != nil {
			panic(err)
		}
		client = s3.NewFromConfig(awsCfg)
	} else {
		// Endpoint preenchido: usar MinIO/local
		customResolver := aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if service == s3.ServiceID {
					return aws.Endpoint{
						URL:               cfgEnv.AWS.S3.Endpoint,
						SigningRegion:     cfgEnv.AWS.Region,
						HostnameImmutable: true,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)
		awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfgEnv.AWS.Region), config.WithEndpointResolverWithOptions(customResolver))
		if err != nil {
			panic(err)
		}
		client = s3.NewFromConfig(awsCfg)
	}

	return &S3FileProvider{
		client:     client,
		bucketName: cfgEnv.AWS.S3.BucketName,
	}
}

func (s *S3FileProvider) UploadFile(fileName string, fileContent []byte) error {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileContent),
	})

	if err != nil {
		fmt.Printf("Erro: %v\n", err.Error())
		fmt.Printf("Bucket: %s\n", s.bucketName)
		fmt.Printf("Key: %s\n", fileName)
		if strings.Contains(err.Error(), "NoSuchBucket") || strings.Contains(err.Error(), "InvalidBucketName") {
			return &exceptions.BucketNotFoundException{}
		}
		return fmt.Errorf("erro ao fazer upload no S3: %w", err)
	}

	return nil
}

func (s *S3FileProvider) DeleteFile(fileName string) error {
	fmt.Printf("[S3FileProvider] Tentando deletar do bucket: %s, arquivo: %s\n", s.bucketName, fileName)
	resp, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		fmt.Printf("[S3FileProvider] Erro ao deletar do bucket: %v\n", err)
		return fmt.Errorf("failed to delete file: %w", err)
	}
	fmt.Printf("[S3FileProvider] DeleteObject response: %+v\n", resp)
	return nil
}

func (s *S3FileProvider) GetPresignedURL(fileName string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	presignedRequest, err := presignClient.PresignGetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(fileName),
		},
		func(o *s3.PresignOptions) {
			o.Expires = 15 * time.Minute
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to get presigned URL: %w", err)
	}

	return presignedRequest.URL, nil
}

func (s *S3FileProvider) DeleteFiles(fileNames []string) error {
	errs := make([]error, 0)
	for _, fileName := range fileNames {
		err := s.DeleteFile(fileName)
		if err != nil {
			errs = append(errs, fmt.Errorf("erro ao remover %s: %w", fileName, err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("erros ao deletar arquivos: %v", errs)
	}
	return nil
}
