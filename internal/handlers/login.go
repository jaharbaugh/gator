package handlers

import(
	"fmt"
	//"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)


func HandlerLogin(s *app.State, cmd app.Command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("No username provided")
	} 
	username := cmd.Args[0]

	_, err := s.DB.GetUser(s.CTX, username)
	if err != nil{
		return fmt.Errorf("Could not find user:%w", err)
	}

	if err := s.CFG.SetUser(username); err != nil{
		return err
	}

	fmt.Printf("The user has been set to %s\n", username)
	
	return nil
}
