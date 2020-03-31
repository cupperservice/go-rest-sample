package service

import (
	. "github.com/yumemi-kkawashima/mydic/message"
	. "github.com/yumemi-kkawashima/mydic/repository"
)

type DicService struct {
	Repo	WordRepository
}

func (o *DicService) GetWord(word *Word) (*Word, error) {
	w, err := o.Repo.GetWord(word)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (o *DicService) AddWordIfNotExist(word *Word) error {
	w, err := o.Repo.GetWord(word)
	if err != nil {
		return err
	}

	if w == nil {
		return o.Repo.AddWord(word)
	} else {
		return err
	}
}

func (o *DicService) GetHistories() ([]Word, error) {
	return o.Repo.GetWords("reference_date desc", 10)
}

func (o *DicService) GetExamples(id int) ([]Example, error) {
	word := &Word { Id: id }
	return o.Repo.GetExamples(word)
}

func (o *DicService) AddExample(example *Example) error {
	return o.Repo.AddExample(example)
}
