package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Miner     Miner
	Email     Email
	Threshold int
}

type Miner struct {
	Path string
	User string
	Pool string
}

type Email struct {
	Start            Start
	Stop             Stop
	From             string
	SubjectWhenStart string
	SubjectWhenStop  string
	To               string
	Smtp             string
	SmtpPort         int
	PasswordFrom     string
	Tls              string
}

type Stop struct {
	Body    string
	Subject string
	Send    bool
}

type Start struct {
	Body    string
	Subject string
	Send    bool
}

func (c *Configuration) Json() string {
	s, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(s)
}

func Load() *Configuration {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("conf.json is missing")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error decoding json file:", err)
	}
	log.Printf("configuration loaded: %s\n", conf.Json())
	return &conf
}
