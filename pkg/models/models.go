package models

import (
	"errors"
	"time"
)

// ErrNoRecord err
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet struct
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
