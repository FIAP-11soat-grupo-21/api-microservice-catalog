package database_errors

import (
	"fmt"
	"regexp"
	"tech_challenge/internal/product/domain/exceptions"
)

func HandleDatabaseErrors(err error) error {
	if err == nil {
		return nil
	}

	fmt.Printf("[HandleDatabaseErrors] Mensagem de erro: %v\n", err.Error())
	code := ExtractDatabaseState(err.Error())
	fmt.Printf("[HandleDatabaseErrors] SQLSTATE extraído: %v\n", code)
	switch code {
	case "":
		fmt.Println("[HandleDatabaseErrors] Nenhum SQLSTATE encontrado, retornando erro original")
		return err
	case "23001", "23503":
		fmt.Println("[HandleDatabaseErrors] Mapeando para CategoryHasProductsException")
		return &exceptions.CategoryHasProductsException{}
	}
	fmt.Println("[HandleDatabaseErrors] Retornando erro original")
	return err
}

func ExtractDatabaseState(msg string) string {
	re := regexp.MustCompile(`SQLSTATE (\d{5})`)
	matches := re.FindStringSubmatch(msg)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
