package main

import (
	"fmt"
	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/handlers"
	"github.com/jaharbaugh/gator/internal/app"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"context"
)

// goose migration goose -dir ./sql/schema postgres "host=localhost port=5432 user=bootdev password=bootdev DBname=bootdev sslmode=disable" up

//connectionString := "postgres://postgres:postgres@localhost:5432/gator"




func main() {
	CFG, err := config.Read()
	if err != nil{
		panic(err)
	}

	DB, err := sql.Open("postgres", CFG.DBURL)
	if err != nil{
		panic(err)
	}
	defer DB.Close()
	DBQueries := database.New(DB)

	s := app.State{
		CFG: &CFG,
		DB: DBQueries,
		CTX: context.Background(),
	}

	cmds := app.Commands{
		CLICommands: make(map[string]func(*app.State, app.Command) error),
	}

	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("Register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerUsers)
	cmds.Register("agg", handlers.HandlerAgg)
	cmds.Register("addfeed", middlewareLoggedIn(handlers.HandlerAddFeed))
	cmds.Register("feeds", handlers.HandlerFeeds)
	cmds.Register("follow", middlewareLoggedIn(handlers.HandlerFollow))
	cmds.Register("following", middlewareLoggedIn(handlers.HandlerFollowing))
	cmds.Register("unfollow", middlewareLoggedIn(handlers.HandlerUnfollow))
	//cmds.Register("browse", handlerBrowse)


	if len(os.Args) < 2{
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]

	cmd := app.Command{Name: name, Args: args}

	if err := cmds.Run(&s, cmd); err !=nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//fmt.Println(CFG)
}