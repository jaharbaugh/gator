package handlers

import(
	"fmt"
	"time"
//	"github.com/jaharbaugh/gator/internal/config"
	"github.com/jaharbaugh/gator/internal/database"
	"github.com/jaharbaugh/gator/internal/app"
//	"os"
//	"database/sql"
//	_ "github.com/lib/pq"
//	"context"
	"database/sql"
	"github.com/google/uuid"
	"strings"
)


func scrapeFeeds(s *app.State, cmd app.Command) error{
	
	
	nextFeed, err := s.DB.GetNextFeedToFetch(s.CTX)
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

	err = s.DB.MarkFetchedAt(s.CTX, feed)
	if err != nil{
		return err
	}

	fetched, err := fetchFeed(s.CTX, nextFeed.Url)
	if err != nil{
		return err
	}

	fmt.Println("Feed fetched successfully")
	for i := range fetched.Channel.Item{
	
		pubtime := sql.NullTime{}
		t, err := time.Parse(time.RFC1123Z, fetched.Channel.Item[i].PubDate) 
		if err == nil{
			pubtime = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}

		var post database.CreatePostParams
		post.ID = uuid.New()
		post.CreatedAt = time.Now().UTC()
		post.UpdatedAt = time.Now().UTC()
		post.Title = fetched.Channel.Item[i].Title
		post.Url = fetched.Channel.Item[i].Link
		post.Description = sql.NullString{
    		String: fetched.Channel.Item[i].Description,
    		Valid:  true,
			}
		post.PublishedAt = pubtime
		post.FeedID = nextFeed.ID

		_, err = s.DB.CreatePost(s.CTX, post)
		if err != nil{
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint"){
				continue
			}else{
				fmt.Println("error creating post:", err)
    			continue
			}
		}
	}
	return nil
}

