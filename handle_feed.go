package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/AAlejandro8/RSS/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	// Now follow the feed you create
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID:  newFeed.ID, 
	})
	if err != nil {
		return err
	}
	fmt.Println("Feed created successfully and followed")
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

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("not enough args: %s", cmd.name)
	}
	feed, err := s.db.GetQueryByURL(context.Background(),cmd.arguments[0])
	if err != nil {
		return err
	}
	followRecord, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(followRecord.FeedName)
	fmt.Println(followRecord.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	FeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, data := range FeedFollows{
		fmt.Println(data.FeedName)
		fmt.Println(data.UserName)
	}
	return nil
}