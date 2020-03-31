package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yumemi-kkawashima/mydic/message"
	"github.com/yumemi-kkawashima/mydic/service"
)

type WordController struct {
	Service		service.DicService
}

/**
 *
 */
func (o *WordController) AddWord(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var addword message.AddWord
	err = json.Unmarshal(body, &addword)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = o.Service.AddWordIfNotExist(&addword.Word)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

/**
 *
 */
func (o *WordController) GetHistories(w http.ResponseWriter, r *http.Request) {
	words, err := o.Service.GetHistories()

	res, err := json.Marshal(message.Histories { Histories: words })
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, "%s", res)
	}
}

/**
 *
 */
func (o *WordController) GetExamples(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	examples, err := o.Service.GetExamples(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
	res, err := json.Marshal(message.Examples { Examples: examples })
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, "%s", res)
	}
}

/**
 *
 */
func (o *WordController) AddExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var addexample message.AddExample
	err = json.Unmarshal(body, &addexample)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	addexample.Example.WordId = id
	err = o.Service.AddExample(&addexample.Example)
}
