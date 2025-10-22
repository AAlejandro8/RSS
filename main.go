package main

import (
	"fmt"
	"github.com/AAlejandro8/RSS/internal/config"
)


func main() {
	cfg, err := config.Read() 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
	
	if err := cfg.SetUser("alejandro"); err != nil {
		fmt.Println(err)
	}
	updated, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updated)
}