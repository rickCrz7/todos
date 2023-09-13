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
	dao2 := NewOwnerDao(conn)
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Todo CLI")
	log.Println("Enter command (task) type owner to access owners: ")
	scanner.Scan()
	cmd := scanner.Text()
	for cmd != "quit" {
		switch cmd {
		case "owner":
			log.Println("Enter command (owner): ")
			scanner.Scan()
			cmd = scanner.Text()
			for cmd != "quit" {
				switch cmd {
					case "list":
						log.Println("List of Owners:")
						owners, err := dao2.GetAll()
						if err != nil {
							log.Fatal(err)
						}
						for _, owner := range owners {
							log.Printf("%s : %s\n", owner.ID, owner.Name)
						}
					case "create":
						log.Println("Enter name:")
						scanner.Scan()
						name := scanner.Text()
						if name == "" {
							log.Println("Invalid name")
							break
						}
						id := uuid.New().String()
						id = id[:3]
						owner := &Owner{
							Name: name,
							ID:   id,
						}
						err := dao2.Create(owner)
						if err != nil {
							log.Fatal(err)
						}
						log.Println("Owner created")
						log.Printf("Name : %s\n", name)
						log.Printf("ID : %s\n", id)
					case "update":
						log.Println("Enter id:")
						scanner.Scan()
						id := scanner.Text()
						if id == "" {
							log.Println("Invalid id")
							break
						}
						log.Println("Enter name:")
						scanner.Scan()
						name := scanner.Text()
						owner := &Owner{
							ID:        id,
							Name:     name,
						}
						err := dao2.Update(owner)
						if err != nil {
							log.Fatal(err)
						}
						log.Println("Owner updated")
						log.Printf("Name : %s\n", name)
						log.Printf("ID : %s\n", id)
					case "delete":
						log.Println("Enter id:")
						scanner.Scan()
						id := scanner.Text()
						if id == "" {
							log.Println("Invalid id")
							break
						}
						err := dao2.Delete(id)
						if err != nil {
							log.Fatal(err)
						}
						log.Println("Owner deleted")
					default:
						log.Println("Invalid command")
				}
				log.Println("Enter command (owner): ")
				scanner.Scan()
				cmd = scanner.Text()
			}

		case "list":
			log.Println("List of Todos:")
			todos, err := dao.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			for _, todo := range todos {
				log.Printf("%s : %s : %s(%s): %t \n", todo.ID, todo.Title, todo.OwnerName, todo.Owner_ID, todo.Completed)
			}
		case "create":
			log.Println("Enter title:")
			scanner.Scan()
			title := scanner.Text()
			if title == "" {
				log.Println("Invalid title")
				break
			}
			log.Println("Enter owner_id:")
			scanner.Scan()
			owner_id := scanner.Text()
			id := uuid.New().String()
			id = id[:3]
			todo := &Todo{
				Title: title,
				ID:   id,
				Owner_ID: owner_id,
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
		log.Println("Enter command (task): ")
		scanner.Scan()
		cmd = scanner.Text()
	}
	log.Print("Exiting")
}
