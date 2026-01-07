package env

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv        string
	APIPort      string
	APIHost      string
	APIUrl       string
	APIUploadUrl string
	Database     struct {
		RunMigrations bool
		Host          string
		Name          string
		Port          string
		Username      string
		Password      string
	}
	AWS struct {
		Region string
		S3     struct {
			BucketName        string
			Endpoint          string
			PresignExpiration string
		}
	}
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.Load()
		fmt.Printf("[DEBUG] Configuração final: Bucket=%s, Endpoint=%s, Region=%s\n", instance.AWS.S3.BucketName, instance.AWS.S3.Endpoint, instance.AWS.Region)
	})
	return instance
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func getEnvOptional(key string) string {
	return os.Getenv(key)
}

func (c *Config) Load() {
	dotEnvPath := ".env.aws"
	_, err := os.Stat(dotEnvPath)
	if err == nil {
		err := godotenv.Load(dotEnvPath)
		if err != nil {
			log.Fatalf("Erro ao carregar .env: %v", err)
		}
	}

	c.GoEnv = getEnv("GO_ENV")

	c.APIPort = getEnv("API_PORT")
	c.APIHost = getEnv("API_HOST")
	c.APIUploadUrl = getEnv("API_UPLOAD_URL")
	c.APIUrl = c.APIHost + ":" + c.APIPort

	c.Database.RunMigrations = getEnv("DB_RUN_MIGRATIONS") == "true"
	c.Database.Host = getEnv("DB_HOST")
	c.Database.Name = getEnv("DB_NAME")
	c.Database.Port = getEnv("DB_PORT")
	c.Database.Username = getEnv("DB_USERNAME")
	c.Database.Password = getEnv("DB_PASSWORD")

	c.AWS.Region = getEnv("AWS_REGION")

	c.AWS.S3.BucketName = getEnv("AWS_S3_BUCKET_NAME")
	fmt.Printf("[Config] Bucket lido do env: %s\n", c.AWS.S3.BucketName)
	c.AWS.S3.Endpoint = getEnvOptional("AWS_S3_ENDPOINT")
	c.AWS.S3.PresignExpiration = getEnv("AWS_S3_PRESIGN_EXPIRATION")
}

func (c *Config) IsProduction() bool {
	return c.GoEnv == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.GoEnv == "development"
}
func ResetConfig() {
	instance = nil
	once = sync.Once{}
}
