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

// 1. Defina a interface para o client S3

type S3Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

// 2. Altere o S3FileProvider para usar a interface

type S3FileProvider struct {
	client     S3Client
	bucketName string
}

func NewS3FileProvider() *S3FileProvider {
	cfgEnv := env.GetConfig()
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
		//nolint:staticcheck
		customResolver := aws.EndpointResolverWithOptionsFunc( //nolint:staticcheck
			func(service, region string, options ...interface{}) (aws.Endpoint, error) { //nolint:staticcheck
				if service == s3.ServiceID {
					return aws.Endpoint{ //nolint:staticcheck
						URL:               cfgEnv.AWS.S3.Endpoint,
						SigningRegion:     cfgEnv.AWS.Region,
						HostnameImmutable: true,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{} //nolint:staticcheck
			},
		)
		awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfgEnv.AWS.Region), config.WithEndpointResolverWithOptions(customResolver)) //nolint:staticcheck
		if err != nil {
			panic(err)
		}
		client = s3.NewFromConfig(awsCfg)
	}

	// 3. No NewS3FileProvider, converta o client para S3Client
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
		if strings.Contains(err.Error(), "NoSuchBucket") || strings.Contains(err.Error(), "InvalidBucketName") {
			return &exceptions.BucketNotFoundException{}
		}
		return fmt.Errorf("erro ao fazer upload no S3: %w", err)
	}

	return nil
}

func (s *S3FileProvider) DeleteFile(fileName string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

func (s *S3FileProvider) GetPresignedURL(fileName string) (string, error) {
	// Faz type assertion para *s3.Client
	client, ok := s.client.(*s3.Client)
	if !ok {
		return "", fmt.Errorf("client não é *s3.Client, não é possível gerar presigned URL")
	}
	presignClient := s3.NewPresignClient(client)

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
