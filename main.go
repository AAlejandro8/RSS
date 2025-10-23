package main

import (
	"fmt"
	"os"

	"github.com/AAlejandro8/RSS/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	callBackCommands map[string]func(s *state, cmd command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expected a username")
	}
	username := cmd.arguments[0]
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("set user: %w", err)
	}
	fmt.Printf("user set to %s\n", username)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	// check if the command given exists
	exists, ok := c.callBackCommands[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command %s", cmd.name)
	} 
	return exists(s, cmd)
}
func (c *commands) register(name string, f func(*state, command) error) {
	// 'register' the command
	c.callBackCommands[name] = f
}

func main() {
	readConfig, err := config.Read() 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	myState := state {
		cfg: &readConfig,
	}
	cmds := commands {
		callBackCommands: map[string]func(s *state, cmd command) error {},
	}
	cmds.register("login", handlerLogin)



	if len(os.Args) < 2 {
		fmt.Println("Error must be more than two arguments")
		os.Exit(1)
	}
	enteredCmd := os.Args[1]
	argumentSlice := os.Args[2:]

	command := command{
		name: enteredCmd,
		arguments: argumentSlice,
	}

	if err = cmds.run(&myState, command); err != nil {
		fmt.Println(err)
	}

}