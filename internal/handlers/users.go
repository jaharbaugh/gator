package handlers

import(
	"fmt"
	//"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)

func HandlerUsers(s *app.State, cmd app.Command) error {
	listOfUsers, err := s.DB.GetUsers(s.CTX)
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}

	currentUser := s.CFG.Current_User_Name

	for _, user := range listOfUsers{
		if user == currentUser{
			fmt.Printf("* %v (current)\n", user)
		}else{
		fmt.Printf("* %v\n", user)
		}
	}
	return nil
}