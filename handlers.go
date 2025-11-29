package main

import(
	"fmt"
	"context"
	"time"
	//"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/google/uuid"
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

func handlerRegister(s *state, cmd command) error{
	if len(cmd.Args) != 1 {
    	return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	
	var newUser database.CreateUserParams
	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now().UTC()
	newUser.UpdatedAt = time.Now().UTC()
	newUser.Name = name

	user, err := s.db.CreateUser(context.Background(), newUser); 
	if err != nil{
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
    	return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully!")
	fmt.Printf("* ID: %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
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

func handlerUsers(s *state, cmd command) error {
	listOfUsers, err := s.db.GetUsers(context.Background())
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}

	currentUser := s.cfg.Current_User_Name

	for _, user := range listOfUsers{
		if user == currentUser{
			fmt.Printf("* %v (current)\n", user)
		}else{
		fmt.Printf("* %v\n", user)
		}
	}
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

func handlerAddFeed(feedName string, feedURL string) error{
	currentUser := s.db.User.ID

	var newFeed database.CreateFeedParams
	newFeed.ID = uuid.New()
	newFeed.CreatedAt = time.Now().UTC()
	newFeed.UpdatedAt = time.Now().UTC()
	newFeed.Name = feedName
	newFeed.URL = feedURL
	newFeed.UserID = currentUser

	dbQueries.database.CreateFeed(context.Background(), newFeed)

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil{
		return err
	}

	return nil
}