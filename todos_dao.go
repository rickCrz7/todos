package main

import (
	"database/sql"
	"errors"
)

type TodoDao interface {
	GetAll() ([]*Todo, error)
	Get(id int) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id int) error
}

type TodoDaoImpl struct {
	conn *sql.DB
}

func NewTodoDao(conn *sql.DB) TodoDao {
	return &TodoDaoImpl{conn}
}

func (dao *TodoDaoImpl) GetAll() ([]*Todo, error) {
	return nil, errors.New("Not implemented")
	// rows, err := dao.conn.Query("SELECT * FROM todos")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// todos := make([]*Todo, 0)
	// for rows.Next() {
	// 	todo := new(Todo)
	// 	err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	todos = append(todos, todo)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return todos, nil
}

func (dao *TodoDaoImpl) Get(id int) (*Todo, error) {
	return nil, errors.New("Not implemented")
	// row := dao.conn.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	// todo := new(Todo)
	// err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	// if err != nil {
	// 	return nil, err
	// }
	// return todo, nil
}

func (dao *TodoDaoImpl) Create(todo *Todo) (*Todo, error) {
	return nil, errors.New("Not implemented")
	// stmt, err := dao.conn.Prepare("INSERT INTO todos (title, completed, created_at, updated_at) VALUES (?, ?, ?, ?)")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// res, err := stmt.Exec(todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt)
	// if err != nil {
	// 	return nil, err
	// }
	// todo.ID, err = res.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }
	// return todo, nil
}

func (dao *TodoDaoImpl) Update(todo *Todo) (*Todo, error) {
	return nil, errors.New("Not implemented")
	// stmt, err := dao.conn.Prepare("UPDATE todos SET title = ?, completed = ?, updated_at = ? WHERE id = ?")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(todo.Title, todo.Completed, todo.UpdatedAt, todo.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// return todo, nil
}

func (dao *TodoDaoImpl) Delete(id int) error {
	return errors.New("Not implemented")
	// stmt, err := dao.conn.Prepare("DELETE FROM todos WHERE id = ?")
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(id)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
