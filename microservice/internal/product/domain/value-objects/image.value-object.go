package value_objects

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"tech_challenge/internal/product/domain/exceptions"
	"tech_challenge/internal/shared/config/env"
	"time"
)

const DEFAULT_IMAGE_FILE_NAME = "default_product_image.webp"

type Image struct {
	FileName string
	Url      string
}

type ImageValue struct {
	FileName string
	Url      string
}

func NewImage(originalFileName string) (Image, error) {
	if originalFileName == "" {
		return Image{}, &exceptions.InvalidProductDataException{
			Message: "Image file name is required",
		}
	}

	sanitizedName := sanitizeFileName(originalFileName)

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := filepath.Ext(sanitizedName)
	baseName := strings.TrimSuffix(sanitizedName, ext)
	fileName := fmt.Sprintf("%s_%s%s", baseName, timestamp, ext)

	if len(fileName) > 255 {
		return Image{}, &exceptions.InvalidProductDataException{
			Message: fmt.Sprintf("Image file name '%s' exceeds the maximum length of 255 characters", fileName),
		}
	}
	if len(fileName) == 0 {
		return Image{}, &exceptions.InvalidProductDataException{
			Message: "Image file name cannot be empty",
		}
	}

	config := env.GetConfig()

	imageHost := config.APIUploadUrl

	imageUrl := fmt.Sprintf("%s/%s", imageHost, fileName)

	return Image{
		FileName: fileName,
		Url:      imageUrl,
	}, nil
}

func NewImageDefault() (Image, error) {
	config := env.GetConfig()

	imageHost := config.APIUploadUrl

	imageUrl := fmt.Sprintf("%s/%s", imageHost, DEFAULT_IMAGE_FILE_NAME)

	return Image{
		FileName: DEFAULT_IMAGE_FILE_NAME,
		Url:      imageUrl,
	}, nil
}

func NewImageWithFileNameAndUrl(fileName, url string) (Image, error) {
	if fileName == "" || url == "" {
		return Image{}, &exceptions.InvalidProductDataException{
			Message: "Image file name and URL are required",
		}
	}

	return Image{
		FileName: fileName,
		Url:      url,
	}, nil
}

func sanitizeFileName(fileName string) string {
	var sanitized strings.Builder
	for _, r := range fileName {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '.' || r == '_' || r == '-' {
			sanitized.WriteRune(r)
		}
	}
	return sanitized.String()
}

func (i *Image) IsDefault() bool {
	return i.FileName == DEFAULT_IMAGE_FILE_NAME
}

func (i *Image) Value() ImageValue {
	return ImageValue{
		FileName: i.FileName,
		Url:      i.Url,
	}
}
