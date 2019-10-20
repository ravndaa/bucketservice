package bucketservice

// Store can be memory or DB
type Store interface{}

// BucketStore hmm
type BucketStore interface {
	//Init(store Store) (BucketStore, error)
	Store
	Create() (string, error)
	Delete(bucketid string) error
	AddMessage(bucketid string, msg Message) error
	GetAllMessages(bucketid string) ([]Message, error)
	DoesBucketExist(bucketid string) bool
}
