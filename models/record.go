package models

import "time"

type RecordItem struct {
	ID         int64     `db:"id" json:"id"`
	CreatedAt  time.Time `db:"createdAt" json:"createdAt"`
	TotalMarks int       `db:"totalMarks" json:"totalMarks"`
}
