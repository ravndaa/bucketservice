package bucketservice

import (
	"net/http"
	"time"
)

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
