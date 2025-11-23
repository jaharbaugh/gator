package main

import (
	"fmt"
	"github.com/jaharbaugh/gator/internal/config"
	"os"
)

//connectionString := "postgres://postgres:postgres@localhost:5432/gator"

func main() {
	cfg, err := config.Read()
	if err != nil{
		panic(err)
	}

	s := state{
		cfg: &cfg,
	}

	cmds := commands{
		cliCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2{
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]

	cmd := command{Name: name, Args: args}

	if err := cmds.run(&s, cmd); err !=nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}