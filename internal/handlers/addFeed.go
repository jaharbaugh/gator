package handlers

import(
	"fmt"
	"time"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
	"github.com/google/uuid"

)

func HandlerAddFeed(s *app.State, cmd app.Command, user database.User) error{

	if len(cmd.Args) != 2 {
        return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
    }

	feedName := cmd.Args[0]
    feedURL := cmd.Args[1]
		
	var newFeed database.CreateFeedParams
	newFeed.ID = uuid.New()
	newFeed.CreatedAt = time.Now().UTC()
	newFeed.UpdatedAt = time.Now().UTC()
	newFeed.Name = feedName
	newFeed.Url = feedURL
	newFeed.UserID = user.ID

	feeds, err := s.DB.CreateFeed(s.CTX, newFeed)
	if err != nil{
		return err
	}

	_, err = s.DB.CreateFeedFollow(s.CTX, database.CreateFeedFollowParams{
    ID: uuid.New(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	UserID: user.ID,
    FeedID: feeds.ID,
	})
	if err != nil {
	    return err
	}

	fmt.Println("Feeds created successfully!")
	fmt.Printf("* ID: %v\n", feeds.ID)
	fmt.Printf("* Name: %v\n", feeds.Name)
	fmt.Printf("* URL: %v\n", feeds.Url)
	fmt.Printf("* UserID: %v\n", feeds.UserID)

	return nil
}
