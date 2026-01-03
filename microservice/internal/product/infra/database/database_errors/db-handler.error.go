package database_errors

import (
	"regexp"
	"tech_challenge/internal/product/domain/exceptions"
)

// func HandleDatabaseErrors(err error) error {
// 	if err == nil {
// 		return nil
// 	}
// 	fmt.Println("Entrou no HandleDatabaseErrors")
// 	var perr *pgconn.PgError
// 	if ok := errors.As(err, &perr); ok {
// 		fmt.Println("Conversão para *pgconn.PgError bem-sucedida")
// 		fmt.Println("Agora vou printar o peer.Code")
// 		fmt.Println(perr.Code)
// 		fmt.Println("Fim do print do peer.Code")
// 	} else {
// 		fmt.Println("Conversão para *pgconn.PgError falhou")
// 	}
// 	// ...
// 	fmt.Println("Agora vou printar o peer.Code")
// 	fmt.Println(perr.Code)
// 	fmt.Println("Fim do print do peer.Code")
// 	// Type assertion direta

// 	// if pgErr, ok := err.(*pgconn.PgError); ok {
// 	// 	fmt.Println("Entrou no IF")
// 	// 	switch pgErr.Code {
// 	// 	case "23501": // FK violation normal
// 	// 		fmt.Println("Caiu no Switch")
// 	// 		return &exceptions.CategoryHasProductsException{}
// 	// 	}
// 	// }

// 	// // Fallback para RESTRICT, que GORM não mapeia
// 	// if strings.Contains(err.Error(), "SQLSTATE 23001") {
// 	// 	fmt.Println("Caiu no fallback")
// 	// 	return
// 	// }
// 	return err
// }

func HandleDatabaseErrors(err error) error {
	if err == nil {
		return nil
	}

	code := extractDatabaseState(err.Error())
	switch code {
	case "":
		return err
	case "23001":
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
