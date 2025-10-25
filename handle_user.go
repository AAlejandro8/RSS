package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/AAlejandro8/RSS/internal/database"
	"github.com/google/uuid"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expected a username")
	}
	username := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(),username)
	if err != nil {
		fmt.Println("user doesn't exist!")
		os.Exit(1)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("set user: %w", err)
	}
	fmt.Printf("user set to %s\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expected a name")
	}
	name := cmd.arguments[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create the user: %w", err)

	}
	_, err = s.db.GetUser(context.Background(),user.Name)
	if err != nil {
		fmt.Println("user already exists")
		os.Exit(1)
	
	}

	if err = s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set the user: %w", err)
	}
	fmt.Printf("registed new user %v", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("coulnd't delete the users: %w", err)
	}
	fmt.Println("all users successsfully deleted!")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all users: %w", err)
	}
	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name{
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s \n", user.Name)
	}
	return nil 
}