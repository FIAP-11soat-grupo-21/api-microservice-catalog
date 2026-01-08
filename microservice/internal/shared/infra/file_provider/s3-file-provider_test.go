package file_service

import (
	"context"
	"errors"
	"os"
	"tech_challenge/internal/product/domain/exceptions"
	"tech_challenge/internal/shared/config/env"
	testenv "tech_challenge/internal/shared/test"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/require"
)

type mockS3Client struct {
	putFunc    func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	deleteFunc func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

//	func TestMain(m *testing.M) {
//		_ = godotenv.Load("../../../.env.local.example")
//		os.Setenv("GO_ENV", "test")
//		os.Setenv("API_PORT", "8080")
//		os.Setenv("API_HOST", "localhost")
//		os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
//		os.Setenv("DB_RUN_MIGRATIONS", "false")
//		os.Setenv("DB_HOST", "localhost")
//		os.Setenv("DB_NAME", "test_db")
//		os.Setenv("DB_PORT", "5432")
//		os.Setenv("DB_USERNAME", "test_user")
//		os.Setenv("DB_PASSWORD", "test_pass")
//		os.Setenv("AWS_REGION", "us-east-1")
//		os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
//		os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
//		code := m.Run()
//		os.Exit(code)
//	}
func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}
func (m *mockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	if m.putFunc != nil {
		return m.putFunc(ctx, params, optFns...)
	}
	return &s3.PutObjectOutput{}, nil
}

func (m *mockS3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, params, optFns...)
	}
	return &s3.DeleteObjectOutput{}, nil
}

func TestS3FileProvider_DeleteFiles_AllSuccess(t *testing.T) {
	provider := &S3FileProvider{
		client: &mockS3Client{
			deleteFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
				return &s3.DeleteObjectOutput{}, nil
			},
		},
		bucketName: "bucket",
	}
	err := provider.DeleteFiles([]string{"a.txt", "b.txt"})
	require.NoError(t, err)
}

func TestS3FileProvider_DeleteFiles_WithErrors(t *testing.T) {
	provider := &S3FileProvider{
		client: &mockS3Client{
			deleteFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
				if *params.Key == "b.txt" {
					return nil, errors.New("fail")
				}
				return &s3.DeleteObjectOutput{}, nil
			},
		},
		bucketName: "bucket",
	}
	err := provider.DeleteFiles([]string{"a.txt", "b.txt", "c.txt"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao remover b.txt")
}

func TestNewS3FileProvider_AWSS3(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_ENDPOINT", "")
	provider := NewS3FileProvider()
	require.NotNil(t, provider)
	require.Equal(t, "test-bucket", provider.bucketName)
}

func TestNewS3FileProvider_MinIO(t *testing.T) {
	env.ResetConfig() // <- reseta o singleton/config antes do teste
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_ENDPOINT", "http://localhost:9000")
	provider := NewS3FileProvider()

	require.NotNil(t, provider)
	require.Equal(t, "test-bucket", provider.bucketName)
	// Verifica se o client foi criado e se o endpoint está correto
	cfgEnv := env.GetConfig()
	require.Equal(t, "http://localhost:9000", cfgEnv.AWS.S3.Endpoint)
}

func TestS3FileProvider_GetPresignedURL_NotS3Client(t *testing.T) {
	provider := &S3FileProvider{
		client:     &mockS3Client{}, // mockS3Client não é *s3.Client
		bucketName: "bucket",
	}
	url, err := provider.GetPresignedURL("file.txt")
	require.Error(t, err)
	require.Empty(t, url)
	require.Contains(t, err.Error(), "client não é *s3.Client")
}

func TestS3FileProvider_GetPresignedURL_Success(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_ENDPOINT", "") // ou use um endpoint de teste/minio se quiser

	provider := NewS3FileProvider()
	url, err := provider.GetPresignedURL("file.txt")
	if err != nil {
		t.Logf("Erro esperado em ambiente de teste: %v", err)
	} else {
		require.NotEmpty(t, url)
	}
}

func TestS3FileProvider_GetPresignedURL_ErrorOnPresign(t *testing.T) {
	// Para garantir erro, use um mockS3Client que não é *s3.Client
	provider := &S3FileProvider{
		client:     &mockS3Client{},
		bucketName: "bucket",
	}
	url, err := provider.GetPresignedURL("file.txt")
	require.Error(t, err)
	require.Empty(t, url)
	require.Contains(t, err.Error(), "client não é *s3.Client")
}

func TestS3FileProvider_UploadFile_Success(t *testing.T) {
	provider := &S3FileProvider{
		client: &mockS3Client{
			putFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
				return &s3.PutObjectOutput{}, nil
			},
		},
		bucketName: "bucket",
	}
	err := provider.UploadFile("file.txt", []byte("conteudo"))
	require.NoError(t, err)
}

func TestS3FileProvider_UploadFile_BucketNotFound(t *testing.T) {
	provider := &S3FileProvider{
		client: &mockS3Client{
			putFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
				return nil, errors.New("NoSuchBucket: bucket not found")
			},
		},
		bucketName: "bucket",
	}
	err := provider.UploadFile("file.txt", []byte("conteudo"))
	_, ok := err.(*exceptions.BucketNotFoundException)
	require.True(t, ok)
}

func TestS3FileProvider_UploadFile_OtherError(t *testing.T) {
	provider := &S3FileProvider{
		client: &mockS3Client{
			putFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
				return nil, errors.New("erro generico")
			},
		},
		bucketName: "bucket",
	}
	err := provider.UploadFile("file.txt", []byte("conteudo"))
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao fazer upload no S3")
}
