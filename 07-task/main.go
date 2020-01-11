package main

import (
	"path/filepath"
	"./cmd"
	"./db"
	"os"
	homedir "github.com/mitchellh/go-homedir"
)

func main(){

	home , _ := homedir.Dir()
	dbPath := filepath.Join(home,"my.db")

	err := db.Init(dbPath)
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}
	err = cmd.RootCmd.Execute()
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}
}

