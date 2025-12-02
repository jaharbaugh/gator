package handlers

import(
	"fmt"
	//"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)

func HandlerReset(s *app.State, cmd app.Command) error {
	err := s.DB.DeleteUsers(s.CTX)
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Println("Database successfully reset.")
	return nil
}
