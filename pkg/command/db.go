package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) GetCount() int {
	data, err := ioutil.ReadFile("db.data")
	if err != nil {
		log.Fatalf("error reading db.data %s", err.Error())
	}
	log.Printf("db.data text %s", string(data))

	i, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatalf("file has not a number %s", err.Error())
	}
	return i
}

func (db *DB) SaveCount(count int) {
	err := ioutil.WriteFile("db.data", []byte(fmt.Sprintf("%d", count)), 0777)
	if err != nil {
		log.Fatalf("error writing db.data %s", err.Error())
	}
}
