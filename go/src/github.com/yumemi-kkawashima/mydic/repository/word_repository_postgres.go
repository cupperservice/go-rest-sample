package repository

import (
	"database/sql"
	"fmt"
	"time"
	. "github.com/yumemi-kkawashima/mydic/message"
	_ "github.com/lib/pq"
	// "golang.org/x/xerrors"
)

type DbConfig struct {
	Dbname		string
	Host		string
	Port		int
	User		string
	Password	string
}

func NewRepository(config *DbConfig) (*WordRepositoryPostgres, error) {
	rep := &WordRepositoryPostgres {}

	err := rep.connect(config)
	if err != nil {
		return nil, err
	}

	return rep, nil;
}

type SQL struct {
	insertWord		*sql.Stmt
	getWord			*sql.Stmt
	insertExample	*sql.Stmt
	getExamples		*sql.Stmt
}

type WordRepositoryPostgres struct {
	SQL
	db	*sql.DB
}

func (o *WordRepositoryPostgres) AddWord(word *Word) error {
	now := time.Now()

	err := o.insertWord.QueryRow(word.Content, now, now).Scan(&word.Id)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func (o *WordRepositoryPostgres) GetWord(word *Word) (*Word, error) {
	row := o.getWord.QueryRow(word.Content)
	err := row.Scan(&word.Id, &word.Content, &word.CreateDate, &word.ReferenceDate)
	if err == nil {
		return word, nil
	} else if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return nil, err
	}
}

func (o *WordRepositoryPostgres) GetWords(orderby string, limit int) ([]Word, error) {
	sql := fmt.Sprintf("SELECT id, content, create_date, reference_date from word order by %s limit $1", orderby)
	rows, err := o.db.Query(sql, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := []Word{}
	for rows.Next() {
        word := Word{}
        err = rows.Scan(&word.Id, &word.Content, &word.CreateDate, &word.ReferenceDate)
        if err != nil {
            return nil, err
        }
        words = append(words, word)
    }

	return words, nil
}

func (o *WordRepositoryPostgres) AddExample(example *Example) error {
	now := time.Now()

	err := o.insertExample.QueryRow(example.WordId, example.Content, now, now).Scan(&example.Id)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func (o *WordRepositoryPostgres) GetExamples(word *Word) ([]Example, error) {
	rows, err := o.getExamples.Query(word.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	examples := []Example{}
	for rows.Next() {
        example := Example{}
        err = rows.Scan(&example.Id, &example.WordId, &example.Content, &example.CreateDate, &example.UpdateDate)
        if err != nil {
            return nil, err
        }
        examples = append(examples, example)
    }

	return examples, nil
}

func (o *WordRepositoryPostgres) connect(config *DbConfig) error {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Dbname, config.Password))

	if err == nil {
		o.db = db
		o.prepareStmt(db)
		return nil
	} else {
		return err
	}
}

func (o *SQL) prepareStmt(db *sql.DB) {
	var err error

	o.insertWord, err = db.Prepare("insert into word (content, create_date, reference_date) values ($1, $2, $3) returning id")
	if err != nil {
		panic(err)
	}

	o.getWord, err = db.Prepare("select id, content, create_date, reference_date from word where content = $1")
	if err != nil {
		panic(err)
	}

	o.insertExample, err = db.Prepare("insert into example (word_id, content, create_date, update_date) values ($1, $2, $3, $4) returning id")
	if err != nil {
		panic(err)
	}

	o.getExamples, err = db.Prepare("select id, word_id, content, create_date, update_date from example where word_id = $1 order by create_date desc")
	if err != nil {
		panic(err)
	}
}

func (o *SQL) closeStmt() {
	o.insertWord.Close()
	o.getWord.Close()
}

func (o *WordRepositoryPostgres) Close() {
	o.closeStmt()
	o.db.Close()
}
