package postgresql

import (
	"auth_git/utilts/decrypt"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// func Postgresql() *sql.DB { // 3_15_2024 friday we going to add more error responses for more additional client infomation
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("cannot load .env file: ", err)
// 	}

// 	db, err := sql.Open("postgres", os.Getenv("postgresql"))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Error on connecting to database: %s\n", err.Error())
// 	}
// 	fmt.Println("Connected to PostgreSQL")
// 	return db
// }

func Postgresql() *sql.DB { // 3_15_2024 friday we going to add more error responses for more additional client infomation
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file: ", err)
	}
	encryptedConnectionString := os.Getenv("postgresql")
	if encryptedConnectionString == "" {
		log.Fatal("No PostgreSQL connection information in .env provided")
	}
	splitURL := strings.Split(encryptedConnectionString, "://")
	splitCredentials := strings.SplitN(splitURL[1], ":", 2)
	splitPasswordAndHost := strings.SplitN(splitCredentials[1], "@", 2)
	fmt.Println("splitURL", splitURL)
	fmt.Println("splitCredentials", splitCredentials)
	fmt.Println("splitPasswordAndHost", splitPasswordAndHost)
	encryptedUsername := splitCredentials[0]
	encryptedPassword := splitPasswordAndHost[0]
	fmt.Println("encryptedUsername", encryptedUsername)
	fmt.Println("encryptedPassword", encryptedPassword)
	cipherUsernameEnv, err := decrypt.Detokenize(encryptedUsername)
	if err != nil {
		log.Fatalf("Cannot decrypt username: %s\n", err)
	}
	usernameEnv, err := base64.StdEncoding.DecodeString(cipherUsernameEnv)
	if err != nil {
		log.Fatalf("Cannot decode username: %s\n", err)
	}
	cipherPasswordEnv, err := decrypt.Detokenize(encryptedPassword)
	if err != nil {
		log.Fatalf("Cannot decrypt password: %s\n", err)
	}
	passwordEnv, err := base64.StdEncoding.DecodeString(cipherPasswordEnv)
	if err != nil {
		log.Fatalf("Cannot decode password: %s\n", err)
	}
	fmt.Println("bytes", usernameEnv)
	stringUsernameEnv := string(usernameEnv)
	fmt.Println("string", stringUsernameEnv)
	// decryptedConnectionString := fmt.Sprintf("postgresql://%s:%s@%s/database?sslmode=disable", usernameEnv, passwordEnv, splitPasswordAndHost[1])
	decryptedConnectionString := fmt.Sprintf("postgresql://%s:%s@%s", usernameEnv, passwordEnv, splitPasswordAndHost[1])
	// db, err := sql.Open("postgres", os.Getenv("postgresql"))
	// if err != nil {
	// 	panic(err.Error())
	// }
	fmt.Println("decryptedConnectionString", decryptedConnectionString)
	db, err := sql.Open("postgres", decryptedConnectionString)
	if err != nil {
		log.Fatalf("Error on opening database: %s\n", err)
	}
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Error on connecting to database: %s\n", err.Error())
	// }
	err = db.Ping()
	if err != nil {
		if strings.Contains(err.Error(), "password authentication failed") {
			log.Fatalf("Password authentication failed for user(invalid username or password).\nError: %s", err.Error())
		} else if strings.Contains(err.Error(), "period of time, or established connection failed because connected host has failed to respond.") {
			log.Fatalf("DNS lookup failed(Timeout). Check if the database server address is correct.\nError: %s", err.Error())
		} else if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it") {
			log.Fatalf("Connection refused: Check if the target machine is running and the port is correct. Ensure environment variables in the .env files are correctly set.\nError: %s", err.Error())
		} else {
			log.Fatalf("Error on connecting to database: %s\n", err.Error())
		}
	} else {
		fmt.Println("Connected to PostgreSQL")
	}
	// fmt.Println("Connected to PostgreSQL")
	return db
}
