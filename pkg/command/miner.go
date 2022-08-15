package command

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"minere-startup/pkg/config"
	"os"
	"os/exec"
)

type Miner struct {
	Config *config.Configuration
}

func NewMiner(c *config.Configuration) *Miner {
	return &Miner{Config: c}
}

func (m *Miner) Exec() {
	cmd := exec.Command("command", "/C",
		m.Config.Miner.Path,
		"-a", "ethash",
		"-o", m.Config.Miner.Pool,
		"-u", m.Config.Miner.User,
		"-p", "x",
	)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	log.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
