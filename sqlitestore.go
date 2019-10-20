package bucketservice

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

/*
Move to httpBucket when ready, but for now, here for testing in minibucket =)
*/

// SqliteStore how
type SqliteStore struct {
	Store *sqlx.DB
}

// InitDB a
func (b *SqliteStore) InitDB() (sql.Result, error) {

	schema := `
		CREATE TABLE buckets (
		id int,
		bucketid text);
		CREATE TABLE messages (
			id int,
			bucketid text
			agentname text);`

	// execute a query on the server
	result, err := b.Store.Exec(schema)

	return result, err
}

// Create a
func (b *SqliteStore) Create() (id string, err error) {
	id = randomString(9)
	err = nil

	newBucket := `INSERT INTO buckets (bucketid) VALUES (?)`
	_, err = b.Store.Exec(newBucket, id)
	if err != nil {
		return "", err
	}
	return id, err
}

// Delete a
func (b *SqliteStore) Delete() {

}

// AddMessage a
func (b *SqliteStore) AddMessage(bucketid string, msg Message) error {

	newBucket := `INSERT INTO messages (bucketid, agentname) VALUES (?, ?)`
	_, err := b.Store.Exec(newBucket, bucketid, msg.AgentName)
	if err != nil {
		return err
	}
	return nil
}

// GetAllMessages a
func (b *SqliteStore) GetAllMessages(bucketid string) (messages []Message, err error) {
	err = b.Store.Select(&messages, "SELECT * FROM messages WHERE bucketid = ?", bucketid)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// DoesBucketExist a
func (b *SqliteStore) DoesBucketExist(bucketid string) bool {
	var bucket Bucket
	err := b.Store.Get(&bucket, "SELECT * FROM buckets WHERE bucketid = ?", bucketid)
	if err != nil {
		return false
	}
	if bucket.ID == "" {
		return false
	}
	return true
}
