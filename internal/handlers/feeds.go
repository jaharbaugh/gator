package handlers

import(
	"fmt"
	//"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)


func HandlerFeeds(s *app.State, cmd app.Command) error{
	
	
	listOfFeeds, err := s.DB.GetFeeds(s.CTX)
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
