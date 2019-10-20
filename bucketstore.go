package bucketservice

// BucketStore hmm
type BucketStore interface {
	Create() (string, error)
	Delete(bucketid string) error
	AddMessage(bucketid string, msg Message) error
	GetAllMessages(bucketid string) ([]Message, error)
	DoesBucketExist(bucketid string) bool
}
