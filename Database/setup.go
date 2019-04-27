package dbwrapper

import (
	logger "BitTorrentTracker/Utils/Log"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //Imports the postgres driver
)

// loadEnvironmentVars A function to load DB environment vairables
func loadEnvironmentVars() string {
	host := os.Getenv(Host)
	port := os.Getenv(Port)
	user := os.Getenv(UserName)
	password := os.Getenv(Password)
	dbName := os.Getenv(DBName)

	vars := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	return vars
}

// ConnectDB A function to obtain a database connection
func ConnectDB() *sql.DB {
	err := godotenv.Load()
	logger.LogErr(err, LogSign, "connetDB(): Error loading environment variables", true)

	enVars := loadEnvironmentVars()

	db, err := sql.Open(DriverName, enVars)
	logger.LogErr(err, LogSign, "connetDB(): Error opening Database", true)

	err = db.Ping()
	logger.LogErr(err, LogSign, "connetDB(): Database ping test failed", true)

	logger.LogSuccess(err, LogSign, "Successfully Connected")

	return db
}

// ExecuteQuery A function to execute queries that don't return any rows
func ExecuteQuery(sqlStatement string, logMsgs logger.LogInfo, abort bool, params ...interface{}) bool {
	db := ConnectDB()
	defer db.Close()

	_, err := db.Exec(sqlStatement, params...)
	logger.LogErr(err, LogSign, "executeQuery():"+logMsgs.Error, abort)
	logger.LogSuccess(err, LogSign, logMsgs.Success)

	return (err == nil)
}

// ExecuteRowQuery A function to execute queries that are expected to return at most 1 row
func ExecuteRowQuery(sqlStatement string, params ...interface{}) *sql.Row {
	db := ConnectDB()
	defer db.Close()

	row := db.QueryRow(sqlStatement, params...)

	return row
}

// ExecuteRowsQuery A function to execute queries that are expected to return many rows
func ExecuteRowsQuery(sqlStatement string, logMsgs logger.LogInfo, abort bool, params ...interface{}) (*sql.Rows, bool) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query(sqlStatement, params...)
	logger.LogErr(err, LogSign, logMsgs.Error, false)
	logger.LogSuccess(err, LogSign, logMsgs.Success)

	return rows, (err == nil)
}

// Migrate A function to perform the DB migration
func Migrate() {
	sqlStatement := CreatePeersTable + CreateDownloadsTable + CreatePeerDownloadTable

	logMsgs := logger.LogInfo{
		Success: "Successfully migrated the Database",
		Error:   "Database Migration failed",
	}

	ExecuteQuery(sqlStatement, logMsgs, true)
}

// CleanUP A function to perform DB clean up
func CleanUP() {
	sqlStatement := DropPeersTable + DropDownloadsTable + DropPeerDownloadTable

	logMsgs := logger.LogInfo{
		Success: "Successfully cleaned up the Database",
		Error:   "Clean Up failed",
	}

	ExecuteQuery(sqlStatement, logMsgs, true)
}
