package main

import (
	"fmt"
	"github.com/IIAkSISII/tasktracker/internal/config"
	"github.com/IIAkSISII/tasktracker/internal/database"
)

func main() {
	cfg, err := config.NewConfig(config.GetConfigPath())
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()
}
