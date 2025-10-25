package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AAlejandro8/RSS/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	currentUser, err := s.db.GetUser(context.Background(),s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
		Url: url,
		UserID: currentUser.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Feed created successfully:")
	return nil
}

func handlerFeed(s *state, cmd command) error {
	feed, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feed) == 0 {
		return fmt.Errorf("no feeds found")
	}
	for _, data  := range feed {
		fmt.Println(data.FeedName)
		fmt.Println(data.UserName)
	}
	return nil
}