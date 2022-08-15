package main

import (
	"log"
	"minere-startup/pkg"
	"minere-startup/pkg/config"
	"os"
)

func main() {
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	conf := config.Load()

	log.Println("starting minere startup")
	cmd := pkg.NewCmd(conf)
	cmd.Exec()
}

/*
//TODO check if setenv works
	_, err := exec.Command("command", "/C", env).CombinedOutput()
	if err != nil {
		log.Fatalf("failed setting env var with %s\n", err)
	}
	log.Println(env)

		command.SetEnvVar("set GPU_MAX_ALLOC_PERCENT=100")
		command.SetEnvVar("set GPU_SINGLE_ALLOC_PERCENT=100")
		command.SetEnvVar("set GPU_MAX_HEAP_SIZE=100")
		command.SetEnvVar("set GPU_USE_SYNC_OBJECTS=100")
*/
