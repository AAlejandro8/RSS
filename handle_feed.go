package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetQueryByURL(context.Background(),cmd.arguments[0])
	if err != nil {
		return err 
	}
	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to unfollow: %w", err)
	}
	fmt.Println("Sucessfully unfollowed!")
	return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	if err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return err
	}
	fetchedFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	if _, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title: fetchedFeed.Channel.Title,
		Url: fetchedFeed.Channel.Link,
		Description: fetchedFeed.Channel.Description,
		PublishedAt: time.Now().UTC(),
		FeedID: nextFeed.ID,
	}); err != nil {
		return err
	}
	fmt.Println("post successfully logged to db")
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1{
		if specifiedLimit, err := strconv.Atoi(cmd.arguments[0]); err == nil{
			limit = specifiedLimit
		}else{
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	
	return nil
}