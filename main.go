package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/AAlejandro8/RSS/internal/config"
	"github.com/AAlejandro8/RSS/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}

func main() {
	// read the file 
	cfg, err := config.Read() 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// load db
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}	
	defer db.Close()
	dbQueries := database.New(db)

	myState := state {
		cfg: &cfg,
		db: dbQueries,
	}
	

	cmds := commands {
		callBackCommands: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

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