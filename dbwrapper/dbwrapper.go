package dbwrapper

import (
	"BitTorrentTracker/bittorrent"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // here
)

// Migrate A function to perform DB migration
func Migrate() {
	db := connectDB()
	defer db.Close()

	sqlStatement :=
		`
		DROP TABLE peers;
		CREATE TABLE peers (
		id SERIAL PRIMARY KEY,
		peer_id varchar(60) UNIQUE,
		port varchar(60) NOT NULL,
		ip varchar(60) NOT NULL
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

// ExecuteQuery ;
func ExecuteQuery(sqlStatement string) {
	db := connectDB()
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("[DB] Successfully Executed!")

}

// CreatePeerDownload ;
func CreatePeerDownload(uploaded int, downloaded int, left int, event bittorrent.EventType) bittorrent.PeerDownload {
	db := connectDB()
	defer db.Close()

	sqlStatement :=
		`
		INSERT INTO peerdownloads (uploaded, downloaded, left, event)
		VALUES ($1, $2, $3, $4)
		`

	_, err := db.Exec(sqlStatement, uploaded, downloaded, left, event)

	var peerDownload bittorrent.PeerDownload

	if err != nil {
		fmt.Println(err)
	} else {
		peerDownload.Uploaded = uploaded
		peerDownload.Downloaded = downloaded
		peerDownload.Left = left
		peerDownload.Event = event

		fmt.Println("[DB] Created Peer-Download successfully")
	}

	return peerDownload
}

// CreateDownload A function to insert a download
func CreateDownload(infoHash string) bittorrent.Download {
	db := connectDB()
	defer db.Close()

	sqlStatement :=
		`
		INSERT INTO downloads (infoHash)
		VALUES ($1)
		`

	_, err := db.Exec(sqlStatement, infoHash)

	var download bittorrent.Download

	if err != nil {
		fmt.Println(err)
	} else {
		download.InfoHash = infoHash

		fmt.Println("[DB] Created Download successfully")
	}

	return download
}

// CreatePeer A function to insert a peer into the DB
func CreatePeer(peerID, port, ip string) bittorrent.Peer {
	db := connectDB()
	defer db.Close()

	sqlStatement :=
		`
		INSERT INTO peers (peerID, port, ip)
		VALUES ($1, $2, $3)
		`

	_, err := db.Exec(sqlStatement, peerID, port, ip)
	var peer bittorrent.Peer

	if err != nil {
		fmt.Println(err)
	} else {
		peer.ID = peerID
		peer.IP = ip
		peer.Port = port

		fmt.Println("[DB] Created Peer successfully")
	}

	return peer
}
