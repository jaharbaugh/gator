package handlers

import(
	"fmt"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)

func HandlerFollowing(s *app.State, cmd app.Command, user database.User) error {
	
	listOfFeeds, err := s.DB.GetFeedFollowsForUser(s.CTX, user.ID)
	if err != nil {
		return err
	}

	for _, feed := range listOfFeeds{
		fmt.Printf("* %v\n", feed.FeedName)
	}

	return nil
}