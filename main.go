package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}
	defer conn.Close()

	dao := NewTodoDao(conn)
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Todo CLI")
	log.Println("Enter command: ")
	scanner.Scan()
	cmd := scanner.Text()
	for cmd != "quit" {
		switch cmd {
		case "list":
			log.Println("List of Todos:")
			todos, err := dao.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			for _, todo := range todos {
				log.Printf("%s : %s : %t\n", todo.ID, todo.Title, todo.Completed)
			}
		case "create":
			log.Println("Enter title:")
			scanner.Scan()
			title := scanner.Text()
			if title == "" {
				log.Println("Invalid title")
				break
			}
			id := uuid.New().String()
			id = id[:3]
			todo := &Todo{
				Title: title,
				ID:   id,
			}
			err := dao.Create(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo created")
			log.Printf("Title : %s\n", title)
			log.Printf("ID : %s\n", id)
		case "update":
			log.Println("Enter id:")
			scanner.Scan()
			id := scanner.Text()
			if id == "" {
				log.Println("Invalid id")
				break
			}
			log.Println("Enter title:")
			scanner.Scan()
			title := scanner.Text()
			todo := &Todo{
				ID:        id,
				Title:     title,
			}
			err := dao.Update(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo updated")
			log.Printf("Title : %s\n", title)
			log.Printf("ID : %s\n", id)
		case "done" :
			log.Println("Enter id:")
			scanner.Scan()
			id := scanner.Text()
			if id == "" {
				log.Println("Invalid id")
				break
			}
			err := dao.Done(id)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Todo done! ID : %s\n", id)
		case "delete":
			log.Println("Enter id:")
			scanner.Scan()
			id := scanner.Text()
			if id == "" {
				log.Println("Invalid id")
				break
			}
			err := dao.Delete(id)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo deleted")
		default:
			log.Println("Invalid command")
		}
		log.Println("Enter command: ")
		scanner.Scan()
		cmd = scanner.Text()
	}
	log.Print("Exiting")
}
