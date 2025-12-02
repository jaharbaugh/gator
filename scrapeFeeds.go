package main

import(
	"fmt"
	"time"
//	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
//	"os"
//	"database/sql"
//	_ "github.com/lib/pq"
	"context"
	"database/sql"
)


func scrapeFeeds(s *state, cmd command) error{
	ctx := context.Background()
	
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil{
		return err
	}

	var feed database.MarkFetchedAtParams
	feed.LastFetchedAt = sql.NullTime{
    	Time:  time.Now().UTC(),
    	Valid: true,
		}
	feed.UpdatedAt = time.Now().UTC()
	feed.ID = nextFeed.ID

	err = s.db.MarkFetchedAt(ctx, feed)
	if err != nil{
		return err
	}

	fetched, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil{
		return err
	}

	fmt.Println("Feed fetched successfully")
	for i := range fetched.Channel.Item{
		fmt.Printf("* %v\n", fetched.Channel.Item[i].Title)
	}
	return nil
}