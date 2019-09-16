package command

import (
	"../logger"
	"os/exec"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

//read conf.ini and execute command from there

//Command to execute
type Command struct {
	file string
	args []string
}

var (
	queue         []Command
	existCommands map[string]Command
)

func init() {
	existCommands = make(map[string]Command)
	cfg, _ := ini.Load("conf.ini")
	cmdSection, _ := cfg.GetSection("commands")
	for _, key := range cmdSection.Keys() {
		cmdKey := strings.ReplaceAll(key.Name(), "_", " ")
		confCommand := strings.Split(key.Value(), ",")
		file := confCommand[0]
		args := confCommand[1:]
		existCommands[cmdKey] = Command{file, args}

	}
}

//Add for executing later
func Add(key string) {
	log := logger.GetInstance()
	if cmd, ok := existCommands[key]; ok == true {
		log.Printf("adding \"%v\" command %v\n",key, cmd)
		queue = append(queue, cmd)
	}
}

func pop() Command {
	c := queue[0]
	queue = queue[1:]
	return c
}

//Start command, waiting for command to execute
func Start() {
	log := logger.GetInstance()
	for {
		if len(queue) != 0 {
			toExec := pop()
			out, _ := exec.Command(toExec.file, toExec.args...).Output()
			log.Printf("out: %s", out)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
