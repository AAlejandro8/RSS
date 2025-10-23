package main

import "fmt"


func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expected a username")
	}
	username := cmd.arguments[0]
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("set user: %w", err)
	}
	fmt.Printf("user set to %s\n", username)
	return nil
}

