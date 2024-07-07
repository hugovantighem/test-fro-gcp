package infra

import "time"

type Delegation struct {
	Id          int
	Amount      int
	SenderAddr  string
	BlockHeight int
	Timestamp   time.Time `gorm:"column:ts"`
	Year        int
}
