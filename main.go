package main

import (
	"fmt"
	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

//connectionString := "postgres://postgres:postgres@localhost:5432/gator"

type state struct {
	cfg *config.Config
	db *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil{
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db: dbQueries,
	}

	cmds := commands{
		cliCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIN(handlerUnfollow))

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

	//fmt.Println(cfg)
}