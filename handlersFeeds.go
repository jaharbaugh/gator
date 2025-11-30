package main

import(
	"fmt"
	"context"
	"time"
	//"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error{
	
	if len(cmd.Args) != 2 {
        return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
    }


	feedName := cmd.Args[0]
    feedURL := cmd.Args[1]
	
	username := s.cfg.Current_User_Name
	currentUser, err := s.db.GetUser(context.Background(), username)
	if err != nil{
		return err
	}
	

	var newFeed database.CreateFeedParams
	newFeed.ID = uuid.New()
	newFeed.CreatedAt = time.Now().UTC()
	newFeed.UpdatedAt = time.Now().UTC()
	newFeed.Name = feedName
	newFeed.Url = feedURL
	newFeed.UserID = currentUser.ID

	feeds, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil{
		return err
	}

	fmt.Println("Feeds created successfully!")
	fmt.Printf("* ID: %v\n", feeds.ID)
	fmt.Printf("* Name: %v\n", feeds.Name)
	fmt.Printf("* URL: %v\n", feeds.Url)
	fmt.Printf("* UserID: %v\n", feeds.UserID)

	return nil
}

func handlerFeeds(s *state, cmd command) error{
	listOfFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}


	for _, feed := range listOfFeeds{
		fmt.Printf("* %v\n", feed.Name)
		fmt.Printf("* %v\n", feed.Url)
		fmt.Printf("* %v\n", feed.Name_2)
		
	}
	
	return nil
}
