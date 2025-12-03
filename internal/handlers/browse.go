package handlers

import(
	"fmt"
	"github.com/jaharbaugh/gator/internal/database"
 	"github.com/jaharbaugh/gator/internal/app"
	"strconv"
	"github.com/microcosm-cc/bluemonday"
)


func HandlerBrowse(s *app.State, cmd app.Command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 1{
		i, err := strconv.Atoi(cmd.Args[1])
		if err != nil{
			return err
		}
		limit = i
	}
	
	var userPosts database.GetPostsForUsersParams
	userPosts.UserID = user.ID
	userPosts.Limit = int32(limit) 

	posts, err := s.DB.GetPostsForUsers(s.CTX, userPosts)
	if err != nil{
		return err
	}

	for i := range posts{
		fmt.Printf("* %v\n", posts[i].Title)
		fmt.Printf("* %v\n", posts[i].Url)
		p := bluemonday.StrictPolicy()
		cleanText := p.Sanitize(posts[i].Description.String)
		fmt.Printf("* %v\n", cleanText)
		fmt.Println()
	}

	return nil
}