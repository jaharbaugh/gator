package handlers

import(
	"fmt"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)

func HandlerUnfollow(s *app.State, cmd app.Command, user database.User) error {

	if len(cmd.Args) != 1 {
    	return fmt.Errorf("usage: %s <url>")
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(s.CTX, url)
	if err != nil{
		return err
	}

	var  unfollow database.DeleteFollowByUserAndFeedParams
	unfollow.UserID = user.ID
	unfollow.FeedID = feed.ID

	err = s.DB.DeleteFollowByUserAndFeed(s.CTX, unfollow)
	if err != nil{
		return err
	}

	fmt.Println("Feed unfollow successful!")

	return nil
}