package main
import(
//	"fmt"
//	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
//	"os"
//	"database/sql"
//	_ "github.com/lib/pq"
	"context"
)
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error{
	ctx := context.Background()
	return func (s *state, cmd command) error {
	username := s.cfg.Current_User_Name
	currentUser, err := s.db.GetUser(ctx, username)
	if err != nil{
		return err
	}
	return handler(s, cmd, currentUser)
	}
}