package service

// import (
// 	. "github.com/yumemi-kkawashima/mydic/message"
// 	"testing"
// )

// // AddWord Test

// type TRepo struct {
// 	GetWord_word	*Word
// 	GetWord_err		error
// 	AddWord_word	*Word
// 	AddWord_err		error
// }
// func (o *TRepo) GetWords(order string, limit int) ([]Word, error) {
// 	return nil, nil
// }
// func (o *TRepo) AddWord(word *Word) error {
// 	return o.AddWord_err
// }
// func (o *TRepo) GetWord(word *Word) (*Word, error) {
// 	return o.GetWord_word, o.GetWord_err
// }
// func (o *TRepo) Close() {
// }

// func TestAddWordNew(t *testing.T) {
// 	sv := DicService { Repo: &TRepo {
// 		GetWord_word: nil,
// 		GetWord_err: nil,
// 		AddWord_err: nil,
// 	}}

// 	word := &Word {}

// 	err := sv.AddWordIfNotExist(word)

// 	if err != nil {
// 		t.Errorf("error was occured. %v", err)
// 	}
// }

// // GetHistories Test

// type EmptyRepo struct {}
// func (o *EmptyRepo) GetWords(order string, limit int) ([]Word, error) {
// 	return []Word{}, nil
// }
// func (o *EmptyRepo) AddWord(word *Word) error {
// 	return nil
// }
// func (o *EmptyRepo) GetWord(word *Word) (*Word, error) {
// 	return nil, nil
// }
// func (o *EmptyRepo) Close() {
// }

// func TestHistoriesEmpty(t *testing.T) {
// 	c := DicService { Repo: &EmptyRepo {} }

// 	histories, err := c.GetHistories()

// 	if err != nil {
// 		t.Errorf("error was occured. %v", err)
// 	}

// 	if len(histories) > 0 {
// 		t.Errorf("histories is not empty.%v", histories)
// 	}
// }

// type OneRepo struct {
// 	orderby string
// 	limit	int
// }
// func (o *OneRepo) GetWords(order string, limit int) ([]Word, error) {
// 	o.orderby = order
// 	o.limit = limit
// 	return []Word{ Word {} }, nil
// }
// func (o *OneRepo) AddWord(word *Word) error {
// 	return nil
// }
// func (o *OneRepo) GetWord(word *Word) (*Word, error) {
// 	return nil, nil
// }
// func (o *OneRepo) Close() {
// }

// func TestHistoriesOne(t *testing.T) {
// 	repo := &OneRepo{}
// 	c := DicService { Repo: repo }

// 	histories, err := c.GetHistories()

// 	if err != nil {
// 		t.Errorf("error was occured. %v", err)
// 	}

// 	if len(histories) != 1 {
// 		t.Errorf("histories are not one.%v", histories)
// 	}

// 	if repo.orderby != "reftimestamp desc" {
// 		t.Errorf("orderby is invalid.%v", repo)
// 	}

// 	if repo.limit != 10 {
// 		t.Errorf("list is invalid.%v", repo)
// 	}
// }

// type LimitRepo struct {
// 	orderby	string
// 	limit	int
// }
// func (o *LimitRepo) GetWords(order string, limit int) ([]Word, error) {
// 	o.orderby = order
// 	o.limit = limit
// 	words := []Word{}
// 	for i := 0; i < 10; i++ {
// 		words = append(words, Word {})
// 	}
// 	return words, nil
// }
// func (o *LimitRepo) AddWord(word *Word) error {
// 	return nil
// }
// func (o *LimitRepo) GetWord(word *Word) (*Word, error) {
// 	return nil, nil
// }
// func (o *LimitRepo) Close() {
// }

// func TestHistoryLimit(t *testing.T) {
// 	repo := &LimitRepo {}
// 	c := DicService { Repo: repo }

// 	histories, err := c.GetHistories()

// 	if err != nil {
// 		t.Errorf("error was occured. %v", err)
// 	}

// 	if len(histories) != 10 {
// 		t.Errorf("histories are not limit")
// 	}
// }
