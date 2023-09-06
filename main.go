package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"task-manager/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
