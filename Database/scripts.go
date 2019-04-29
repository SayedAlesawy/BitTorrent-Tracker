package dbwrapper

// CreatePeersTable SQL to create the Peers table
const CreatePeersTable string = `
	CREATE TABLE peers (
	id SERIAL PRIMARY KEY,
	peer_id varchar(60) UNIQUE,
	port varchar(60) NOT NULL,
	ip varchar(60) NOT NULL
	);
`

// CreateDownloadsTable SQL to create the Downloads table
const CreateDownloadsTable string = `
	CREATE TABLE downloads (
	id SERIAL PRIMARY KEY,
	info_hash varchar(60) UNIQUE
	);
`

// CreatePeerDownloadTable SQL to create the Peers_Downloads table
const CreatePeerDownloadTable string = `
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

// DropPeersTable SQL to drop the Peers table
const DropPeersTable string = `DROP TABLE IF EXISTS peers;`

// DropDownloadsTable SQL to drop the Downloads table
const DropDownloadsTable string = `DROP TABLE IF EXISTS downloads;`

// DropPeerDownloadTable SQL to drop the Peers_Downloads table
const DropPeerDownloadTable string = `DROP TABLE IF EXISTS peerdownloads;`

// InsertPeer SQL to insert a peer in the peers table
const InsertPeer string = `
	INSERT INTO peers (peer_id, port, ip)
	VALUES ($1, $2, $3)
`

// InsertDownload SQL to insert a download in the downloads table
const InsertDownload string = `
	INSERT INTO downloads (info_hash)
	VALUES ($1)
`

// InsertPeerDownload SQL to insert a peer_download relationship in the peer_download table
const InsertPeerDownload string = `
	INSERT INTO peerdownloads (uploaded, downloaded, amt_left, event,peer_id,info_hash)
	VALUES ($1, $2, $3, $4, $5, $6)
`

// SelectPeer SQL to select a peer identified by its id
const SelectPeer string = `SELECT * FROM peers WHERE peer_id=$1;`

// SelectPeerList SQL to select peer list based on a common file being downloaded
const SelectPeerList string = `SELECT peer_id FROM peerdownloads where  info_hash = $1`

// SelectDownloads SQL to select all downloads
const SelectDownloads string = `SELECT info_hash FROM downloads`

// SelectPeerDownload SQL to select a certain peerdownload
const SelectPeerDownload string = `SELECT uploaded, downloaded, amt_left, event FROM peerdownloads WHERE info_hash = $1 and peer_id = $2`

// UpdatePeerDownload SQL to update the peer-download relationship
const UpdatePeerStatus string = `
		UPDATE peerdownloads
		SET uploaded = $1, downloaded = $2, amt_left = $3, event = $4
		WHERE peer_id = $5 AND info_hash = $6; 
`
