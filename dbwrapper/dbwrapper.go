package dbwrapper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // here
)

// Migrate :
func Migrate() {
	db := connectDB()
	defer db.Close()
	sqlStatement := `
	DROP TABLE peers;
	CREATE TABLE peers (
	id SERIAL PRIMARY KEY,
	peer_id TEXT UNIQUE,
	port TEXT NOT NULL,
	ip TEXT NOT NULL
	);`

	sqlStatement += `
	DROP TABLE peers;
	CREATE TABLE peers (
	id SERIAL PRIMARY KEY,
	peer_id TEXT UNIQUE,
	port TEXT NOT NULL,
	ip TEXT NOT NULL
	);`

	sqlStatement += `
	DROP TABLE peers;
	CREATE TABLE peers (
	id SERIAL PRIMARY KEY,
	peer_id TEXT UNIQUE,
	port TEXT NOT NULL,
	ip TEXT NOT NULL
	);`

	sqlStatement += `
	DROP TABLE peers;
	CREATE TABLE peers (
	id SERIAL PRIMARY KEY,
	peer_id TEXT UNIQUE,
	port TEXT NOT NULL,
	ip TEXT NOT NULL
	);`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

}

func connectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[DB]Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[DB] Successfully connected!")
	return db
}

/*
 */
func ExecuteQuery(sqlStatement string) {
	db := connectDB()
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("[DB] Successfully Executed!")

}

/*
 */
func CreatePeer(peerID, port, ip string) {
	db := connectDB()
	defer db.Close()
	sqlStatement := `
	INSERT INTO peers (title,link)
	VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, peerID, port, ip)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[DB] Created web link successfully")
}
