package handlers

import(
	"fmt"
	"time"
	//"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
)

func HandlerAgg(s *app.State, cmd app.Command) error{
	
	if len(cmd.Args) != 1 {
    	return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}
	
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
    	return fmt.Errorf("invalid duration: %w", err)
	}
	
		fmt.Printf("Collecting feeds every %+v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s, cmd)
	}
	
}
