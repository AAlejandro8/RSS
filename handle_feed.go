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
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
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
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}