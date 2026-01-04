package database_errors

import (
	"regexp"
	"tech_challenge/internal/product/domain/exceptions"
)

func HandleDatabaseErrors(err error) error {
	if err == nil {
		return nil
	}

	code := extractDatabaseState(err.Error())
	switch code {
	case "":
		return err
	case "23001", "23503":
		return &exceptions.CategoryHasProductsException{}
	}
	return err
}

func extractDatabaseState(msg string) string {
	re := regexp.MustCompile(`SQLSTATE (\d{5})`)
	matches := re.FindStringSubmatch(msg)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
