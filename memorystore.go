package bucketservice

import "errors"

// MemoryStore how
type MemoryStore struct {
	Store map[string]*[]Message
}

// Create a
func (b *MemoryStore) Create() (id string, err error) {
	id = randomString(9)
	err = nil
	b.Store[id] = new([]Message)
	return id, err
}

// Delete a
func (b *MemoryStore) Delete(bucketid string) error {
	if _, ok := b.Store[bucketid]; !ok {
		err := errors.New("bucket not found")
		return err
	}
	delete(b.Store, bucketid)

	return nil
}

// AddMessage a
func (b *MemoryStore) AddMessage(bucketid string, msg Message) error {
	if _, ok := b.Store[bucketid]; !ok {
		err := errors.New("bucket not found")
		return err
	}

	*b.Store[bucketid] = append(*b.Store[bucketid], msg)

	return nil
}

// GetAllMessages a
func (b *MemoryStore) GetAllMessages(bucketid string) (messages []Message, err error) {
	if _, ok := b.Store[bucketid]; !ok {
		err = errors.New("bucket not found")
		messages = []Message{}
	} else {
		messages = *b.Store[bucketid]
	}

	return messages, err
}

// DoesBucketExist a
func (b *MemoryStore) DoesBucketExist(bucketid string) bool {
	if _, ok := b.Store[bucketid]; !ok {
		return false
	}

	return true
}
