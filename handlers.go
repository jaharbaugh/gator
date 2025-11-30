package main

import(
	"fmt"
	"context"
	//"time"
	//"github.com/jaharbaugh/gator/internal/config"
	//"github.com/jaharbaugh/gator/internal/database"
	//"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No username provided")
	} 
	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil{
		return fmt.Errorf("Could not find user:%w", err)
	}

	if err := s.cfg.SetUser(username); err != nil{
		return err
	}

	fmt.Printf("The user has been set to %s\n", username)
	
	return nil
}



func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Println("Database successfully reset.")
	return nil
}



func handlerAgg(s *state, cmd command) error{
	feedURL := "https://www.wagslane.dev/index.xml"
	
	feed, err := fetchFeed(context.Background(),feedURL)
	if err != nil{
		return err
	}

	fmt.Printf("%+v\n", feed)
	
	return nil
}

