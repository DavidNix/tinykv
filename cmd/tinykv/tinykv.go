package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DavidNix/tinykv"
)

func main() {
	db := tinykv.NewDatabase()

	for {
		fmt.Print(">> ")
		r := bufio.NewReader(os.Stdin)
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		argv := strings.TrimSpace(line)
		input := strings.Fields(argv)
		var (
			cmd  string
			args []string
		)
		if len(argv) > 0 {
			cmd = strings.ToUpper(input[0])
			args = input[1:]
		}

		switch cmd {
		case "BEGIN":
			db.Begin()
		case "ROLLBACK":
			if ok := db.Rollback(); !ok {
				fmt.Println("TRANSACTION NOT FOUND")
			}
		case "COMMIT":
			db.Commit()
		case "GET":
			if len(args) == 0 {
				fmt.Println("Missing KEY argument.")
				continue
			}
			fmt.Println(db.Get(args[0]))
		case "SET":
			if len(args) < 2 {
				fmt.Println("Missing KEY VALUE argument.")
				continue
			}
			db.Set(args[0], args[1])
		case "DELETE":
			if len(args) == 0 {
				fmt.Println("Missing KEY argument.")
				continue
			}
			db.Delete(args[0])
		case "COUNT":
			if len(args) == 0 {
				fmt.Println("Missing VALUE argument.")
				continue
			}
			fmt.Println(db.Count(args[0]))
		case "END":
			return
		case "":
			continue

		default:
			fmt.Printf("Unknown command %q.\n", cmd)
		}
	}
}
