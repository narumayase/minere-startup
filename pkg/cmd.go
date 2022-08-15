package pkg

import (
	"log"
	"minere-startup/pkg/command"
	"minere-startup/pkg/config"
)

type Cmd struct {
	Config *config.Configuration
	Db     *command.DB
	Miner  *command.Miner
	Email  *command.Email
}

func NewCmd(config *config.Configuration) *Cmd {
	return &Cmd{
		Config: config,
		Miner:  command.NewMiner(config),
		Db:     command.NewDB(),
		Email:  command.NewEmailCommand(config),
	}
}

func (c *Cmd) Exec() {

	if c.Config.Email.Start.Send {
		c.Email.Send(c.Config.Email.Start.Subject, c.Config.Email.Start.Body)
	}

	command.SetEnv("GPU_MAX_ALLOC_PERCENT", "100")
	command.SetEnv("GPU_SINGLE_ALLOC_PERCENT", "100")
	command.SetEnv("GPU_MAX_HEAP_SIZE", "100")
	command.SetEnv("GPU_USE_SYNC_OBJECTS", "100")

	db := command.NewDB()
	count := db.GetCount()

	if db.GetCount() > c.Config.Threshold {
		log.Printf("mining has been restarted %d times", count)
		db.SaveCount(0)

		if c.Config.Email.Stop.Send {
			c.Email.Send(c.Config.Email.Start.Subject, c.Config.Email.Start.Body)
		}
		log.Println("finalize...")
	} else {
		count++
		db.SaveCount(count)

		log.Println("executing minering...")
		c.Miner.Exec()
	}
}
