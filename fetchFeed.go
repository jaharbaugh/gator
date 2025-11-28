package main


import(
	"net/http"
	"encoding/xml"
	"context"
	"io"
	"html"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil{
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client:= &http.Client{}

	res, err := client.Do(req)
	if err != nil{
		return &RSSFeed{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err !=nil{
		return &feed, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	
	for i := range feed.Channel.Item{
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
	
	return &feed, nil
}