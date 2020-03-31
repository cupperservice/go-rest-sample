package message

import (
	"time"
)

type AddWord struct {
	Word	Word	`json:"word"`
}

type Histories struct {
	Histories	[]Word	`json:"histories"`
}

type AddExample struct {
	Example	Example	`json:"example"`
}

type Examples struct {
	Examples	[]Example	`json:"examples"`
}

type Word struct {
	Id				int			`json:"id"`
	Content			string		`json:"content"`
	CreateDate		time.Time	`json:"create_date"`
	UpdateDate		time.Time	`json:"update_date"`
	ReferenceDate	time.Time	`json:"reference_date"`
}

type Example struct {
	Id				int			`json:"id"`
	WordId			int			`json:"word_id"`
	Content			string		`json:"content"`
	CreateDate		time.Time	`json:"create_date"`
	UpdateDate		time.Time	`json:"update_date"`
}
