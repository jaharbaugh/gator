package handlers

import(
	"fmt"
	"time"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
	"github.com/google/uuid"
)

func HandlerFollow(s *app.State, cmd app.Command, user database.User) error {
	if len(cmd.Args) != 1 {
    	return fmt.Errorf("usage: %s <url>")
	}

	url := cmd.Args[0]
	
	currentFeed, err := s.DB.GetFeedByURL(s.CTX, url)
	if err != nil{
		return err
	}

	var newFeedFollow database.CreateFeedFollowParams
	newFeedFollow.ID = uuid.New()
	newFeedFollow.CreatedAt = time.Now().UTC()
	newFeedFollow.UpdatedAt = time.Now().UTC()
	newFeedFollow.FeedID = currentFeed.ID
	newFeedFollow.UserID = user.ID



	follow, err := s.DB.CreateFeedFollow(s.CTX, newFeedFollow)
	if err != nil{
		return err
	}


	fmt.Println("Feed followed successfully!")
	fmt.Printf("* %v\n", follow.FeedName)
	fmt.Printf("* %v\n", follow.UserName)

	return nil
}