package handlers

import(
	"fmt"
	"time"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
	"github.com/google/uuid"
)

func HandlerRegister(s *app.State, cmd app.Command) error{
	
	if len(cmd.Args) != 1 {
    	return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	
	var newUser database.CreateUserParams
	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now().UTC()
	newUser.UpdatedAt = time.Now().UTC()
	newUser.Name = name

	user, err := s.DB.CreateUser(s.CTX, newUser); 
	if err != nil{
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.CFG.SetUser(user.Name); err != nil {
    	return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully!")
	fmt.Printf("* ID: %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
	return nil
}
