package bucketservice

import "github.com/jmoiron/sqlx"

// SqliteStore how
type SqliteStore struct {
	Store *sqlx.DB
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
func (b *SqliteStore) AddMessage() {

}

// GetAllMessages a
func (b *SqliteStore) GetAllMessages() {

}

// DoesBucketExist a
func (b *SqliteStore) DoesBucketExist() {

}
