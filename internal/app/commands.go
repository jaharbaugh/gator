package app

import(
	"fmt"
)

type Command struct {
	Name	string
	Args []string
}

type Commands struct{
	CLICommands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error{
	h, ok := c.CLICommands[cmd.Name]
	if !ok {
		return fmt.Errorf("No command found: %s", cmd.Name)
	}
	return h(s,cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error){
	c.CLICommands[name] = f
}
