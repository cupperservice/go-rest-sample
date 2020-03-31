package repository

import (
	. "github.com/yumemi-kkawashima/mydic/message"
)

type WordRepository interface {
	AddWord(word *Word) error
	GetWord(word *Word) (*Word, error)
	GetWords(order string, limit int) ([]Word, error)
	AddExample(example *Example) error
	GetExamples(word *Word) ([]Example, error)
	Close()
}
