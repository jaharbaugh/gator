package main
import(
//	"fmt"
//	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
//	"os"
//	"database/sql"
//	_ "github.com/lib/pq"
//	"context"
)
func middlewareLoggedIn(handler func(s *app.State, cmd app.Command, user database.User) error) func(*app.State, app.Command) error{
	
	return func (s *app.State, cmd app.Command) error {
	username := s.CFG.Current_User_Name
	currentUser, err := s.DB.GetUser(s.CTX, username)
	if err != nil{
		return err
	}
	return handler(s, cmd, currentUser)
	}
}