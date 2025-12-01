package main


import(
	"fmt"
	"context"
	"time"
	//"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/google/uuid"
)



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

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	//username := s.cfg.Current_User_Name
	//currentUser, err := s.db.GetUser(context.Background(), username)
	//if err != nil{
	//	return err
	//}
	
	listOfFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, feed := range listOfFeeds{
		fmt.Printf("* %v\n", feed.FeedName)
	}

	return nil
}