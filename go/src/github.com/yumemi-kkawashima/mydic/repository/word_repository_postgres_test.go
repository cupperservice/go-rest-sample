package repository

import (
	"testing"
	"time"
	. "github.com/yumemi-kkawashima/mydic/message"
)

func TestAddWord(t *testing.T) {
	rep, err := NewRepository(config("dic", "dbsvr", 5432, "root", "root00"))
	if err != nil {
		t.Errorf("connect is failed.%v", err)
	}
	defer rep.Close()

	now := time.Now()
	word := Word {
		Content: "word",
		CreateDate: now,
		ReferenceDate: now,
	}

	err = rep.AddWord(&word)
	if err != nil {
		t.Errorf("adding word is failed.%v", err)
	}
	
	if word.Id == 0 {
		t.Error("Id is not setted.")
	}
}

func TestGetWord(t *testing.T) {
	rep, err := NewRepository(config("dic", "dbsvr", 5432, "root", "root00"))
	if err != nil {
		t.Errorf("connect is failed.%v", err)
	}
	defer rep.Close()

	word := Word {
		Content: "word",
	}
	w, err := rep.GetWord(&word)
	if err != nil {
		t.Errorf("getting word is failed. %v", err)
	}
	if w == nil {
		t.Errorf("word is not found")
	}
}

func TestGetWords(t *testing.T) {
	rep, err := NewRepository(config("dic", "dbsvr", 5432, "root", "root00"))
	if err != nil {
		t.Errorf("connect is failed.%v", err)
	}
	defer rep.Close()

	words, err := rep.GetWords("reference_date desc", 10)
	if err != nil {
		t.Errorf("error was occured. %v", err)
	}

	if words == nil {
		t.Errorf("word is not found")
	}

	if len(words) != 0 {
		t.Errorf("%d", len(words))
	}
}

func config(dbname string, host string, port int, user string, password string) *DbConfig {
	config := DbConfig {
		Dbname:		dbname,
		Host:		host,
		Port:		port,
		User:		user,
		Password:	password,
	}

	return &config
}