package main

import(
	"fmt"
	"github.com/jaharbaugh/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	Name	string
	Args []string
}

type commands struct{
	cliCommands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("No username provided")
	} 
	username := cmd.Args[0]
	if err := s.cfg.SetUser(username); err != nil{
		return err
	}

	fmt.Printf("The user has been set to %s\n", username)
	
	return nil
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