package main

import (
	"fmt"
)

type command struct {
	name string
	arguments []string
}

type commands struct {
	callBackCommands map[string]func(s *state, cmd command) error
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
