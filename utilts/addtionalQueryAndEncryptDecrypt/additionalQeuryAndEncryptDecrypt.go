package addtionalQueryAndEncryptDecrypt

import (
	"database/sql"
	"fmt"
	"log"
)

func CountTables(db *sql.DB) {
	var count int
	query := `SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public'`
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatalf("Failed to query table count: %s", err.Error())
	}
	fmt.Printf("There are %d tables in the database.\n", count)
}
