package main

import (
	"fmt"
	"os"

	"github.com/AAlejandro8/RSS/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	readConfig, err := config.Read() 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	myState := state {
		cfg: &readConfig,
	}
	cmds := commands {
		callBackCommands: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", handlerLogin)



	if len(os.Args) < 2 {
		fmt.Println("Error: must be more than two arguments")
		os.Exit(1)
	}
	enteredCmd := os.Args[1]
	argumentSlice := os.Args[2:]

	err = cmds.run(&myState, command{name: enteredCmd, arguments: argumentSlice})
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}