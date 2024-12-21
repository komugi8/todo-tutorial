package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/komugi8/todo-tutorial/cmd"
)


func main() {
    cmd.RunRouter()
}
