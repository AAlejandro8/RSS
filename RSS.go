package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// make the request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error forming request: %w", err)
	}
	// make the client
	client := http.DefaultClient
	
	// set the header to identify the progam to the server
	req.Header.Set("User-Agent", "gator")

	// make the request
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error making the request: %w", err)
	}
	defer res.Body.Close()
	// read the bytes
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading the body: %w", err)
	}
	// make the empty struct and populate
	rssfeed := RSSFeed{}
	if err = xml.Unmarshal(data, &rssfeed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error during unmarshal: %w", err)
	}
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
	for i, item := range rssfeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssfeed.Channel.Item[i] = item
	}
	// return 
	return &rssfeed, nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1{
		return errors.New("not enough arguments")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every: %s\n", timeBetweenRequests)
	
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <- ticker.C {
		scrapeFeeds(s)
	}
	
}