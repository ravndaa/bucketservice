package bucketservice

import (
	"errors"
	"net/http"
	"time"
)

// BucketStore hmm
type BucketStore interface {
	Create() (string, error)
	Delete(bucketid string) error
	AddMessage(bucketid string, msg Message) error
	GetAllMessages(bucketid string) ([]Message, error)
	DoesBucketExist(bucketid string) bool
}

// Bucket ok
type Bucket struct {
	ID string
}

// Message what we receive
type Message struct {
	AgentName string         `json:"AgentName"`
	Headers   http.Header    `json:"headers"`
	Body      string         `json:"body"`
	Cookies   []*http.Cookie `json:"cookies"`
	Host      string         `json:"host"`
	Method    string         `json:"method"`
	URL       string         `json:"url"`
	Fromip    string         `json:"fromip"`
	DateTime  time.Time      `json:"datetime"`
}

// BucketService Service
type BucketService struct {
	DB map[string]*[]Message
}

// Init creates the service
func Init() (b *BucketService) {
	b = &BucketService{
		DB: make(map[string]*[]Message),
	}

	return b
}

// Create => saves the bucket to memory
func (b BucketService) Create() (id string, err error) {
	id = randomString(9)
	err = nil
	b.DB[id] = new([]Message)
	return id, err
}

// Delete => removes the bucket.
func (b BucketService) Delete(bucketid string) error {
	if _, ok := b.DB[bucketid]; !ok {
		err := errors.New("bucket not found")
		return err
	}
	delete(b.DB, bucketid)

	return nil
}

// AddMessage ok
func (b BucketService) AddMessage(bucketid string, msg Message) error {

	if _, ok := b.DB[bucketid]; !ok {
		err := errors.New("bucket not found")
		return err
	}

	*b.DB[bucketid] = append(*b.DB[bucketid], msg)

	return nil
}

// GetAllMessages List all messages from bucket
func (b BucketService) GetAllMessages(bucketid string) (messages []Message, err error) {
	if _, ok := b.DB[bucketid]; !ok {
		err = errors.New("bucket not found")
		messages = []Message{}
	} else {
		messages = *b.DB[bucketid]
	}

	return messages, err
}

// DoesBucketExist check if bucket exist.
func (b BucketService) DoesBucketExist(bucketid string) bool {
	if _, ok := b.DB[bucketid]; !ok {
		return false
	}

	return true
}
