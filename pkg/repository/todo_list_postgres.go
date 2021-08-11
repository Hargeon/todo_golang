package repository

import (
	"fmt"
	"github.com/Hargeon/todo"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var listId int
	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListTable)
	err = tx.QueryRow(query, list.Title, list.Description).Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, list_id) values($1, $2) RETURNING id", usersListTable)
	_, err = tx.Exec(query, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return listId, tx.Commit()
}

func (t *TodoListPostgres) GetLists(userId int) ([]*todo.TodoList, error) {
	var lists []*todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListTable, usersListTable)
	err := t.db.Select(&lists, query, userId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (t *TodoListPostgres) GetList(listId, userId int) (todo.TodoList, error) {
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListTable, usersListTable)
	var list todo.TodoList
	err := t.db.Get(&list, query, userId, listId)
	fmt.Println(list)
	return list, err
}