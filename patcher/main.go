package main

import (
	"fmt"

	"github.com/ketan-10/xo-patcher/xo/cmd"
)

func main() {
	fmt.Println("Hello World")

	cmd.Execute("mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@127.0.0.1:3306/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=true")
}
