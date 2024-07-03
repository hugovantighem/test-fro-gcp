package app

import "time"

// Note: Fields are kept public for convenience, but should be private with getters.
type Delegation struct {
	Id          int
	Amount      int
	SenderAddr  string
	BlockHeight int
	Timestamp   time.Time
	Year        int
}
