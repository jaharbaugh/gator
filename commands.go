package main

import(
	"fmt"
)

type command struct {
	Name	string
	Args []string
}

type commands struct{
	cliCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error{
	h, ok := c.cliCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("No command found: %s", cmd.Name)
	}
	return h(s,cmd)
}

func (c *commands) register(name string, f func(*state, command) error){
	c.cliCommands[name] = f
}
