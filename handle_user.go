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

func handleRegister(s *state, cmd command) error {
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
	if err == nil {
		fmt.Println("user already exists")
		os.Exit(1)
	
	}

	if err = s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set the user: %w", err)
	}

	return nil
}
