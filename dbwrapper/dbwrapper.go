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

	sqlStatement +=
		`
		DROP TABLE downloads;
		CREATE TABLE downloads (
		id SERIAL PRIMARY KEY,
		info_hash varchar(60) UNIQUE
		);`

	sqlStatement +=
		`
		DROP TABLE peerdownloads;
		CREATE TABLE peerdownloads (
		id SERIAL PRIMARY KEY,
		info_hash varchar(60) ,
		peer_id varchar(60),
		uploaded INTEGER,
		downloaded INTEGER,
		amt_left INTEGER,
		event varchar(60)
		);
		ALTER TABLE peerdownloads
		ADD CONSTRAINT uq_peerdownloads UNIQUE(peer_id, info_hash);

		`
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
func CreatePeerDownload(uploaded int, downloaded int, left int, event bittorrent.EventType, peerID string, infoHash string) bittorrent.PeerDownload {
	db := connectDB()
	defer db.Close()

	sqlStatement :=
		`
		INSERT INTO peerdownloads (uploaded, downloaded, amt_left, event,peer_id,info_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	_, err := db.Exec(sqlStatement, uploaded, downloaded, left, event, peerID, infoHash)

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
		INSERT INTO downloads (info_hash)
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
		INSERT INTO peers (peer_id, port, ip)
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

//GetPeers :
func GetPeers(infoHash string) []bittorrent.Peer {
	db := connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT peer_id FROM peerdownloads where  info_hash = $1", infoHash)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	var peerIDs []bittorrent.Peer
	for rows.Next() {
		var peer_id string
		err = rows.Scan(&peer_id)
		if err != nil {
			// handle this error
			panic(err)
		}
		peerIDs = append(peerIDs, getPeer(peer_id))

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return peerIDs
}

//GetPeer :
func getPeer(peerID string) bittorrent.Peer {
	db := connectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM peers WHERE peer_id=$1;`

	var peer bittorrent.Peer
	var dummyID int

	row := db.QueryRow(sqlStatement, peerID)
	switch err := row.Scan(&dummyID, &peer.ID, &peer.Port, &peer.IP); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(peer)
	default:
		panic(err)

	}
	return peer
}
